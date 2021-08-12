package taskmaker

import (
	"context"
	"mnimidamonbackend/client/backup"
	"mnimidamonbackend/client/group_computer"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
	"time"
)

var TaskMaker *taskMakerImpl

func init() {
	TaskMaker = &taskMakerImpl{
		tasks: []QueuedTask{},
	}

	// When the current computer is selected.
	events.CurrentComputerUpdated.Register(TaskMaker)
	// When the configuration reset then it means that it should stop.
	events.RestartConfiguration.Register(TaskMaker)
}

type taskMakerImpl struct {
	tasks []QueuedTask

	cancelContext context.CancelFunc
}

func (t *taskMakerImpl) HandleRestartConfigurationHandler() {
	t.Stop()
}

func (t *taskMakerImpl) HandleCurrentComputerUpdated() {
	t.Start()
}

func (t *taskMakerImpl) Stop() {
	if t.cancelContext != nil {
		t.cancelContext()
		t.cancelContext = nil
	}

	t.DumpTasks()
}

// Starts making tasks.
func (t *taskMakerImpl) Start() {
	t.Stop()

	// Make the context, so it can be canceled.
	ctx := context.Background()
	ctx, t.cancelContext = context.WithCancel(ctx)

	// Spin up the task queuer.
	go t.TaskQueuer(ctx)
}

func (t *taskMakerImpl) TaskQueuer(ctx context.Context) {
	global.Log("task queuer started")
	defer t.DumpTasks()

	for {
		select {
		case <-ctx.Done():
			global.Log("task queuer: context canceled, stopping.")
			return
		default:
			// Get all the group computers of the current computer.
			backups, err := t.GetAllBackups(ctx)
			if err != nil {
				global.Log("task queuer: sleeping for 5 seconds because of an error")
				time.Sleep(time.Second * 5)
				continue
			}

			// Get the deletion tasks.
			deletionTasks := GetAllDeletionTasks(backups)
			// Get the upload tasks.
			uploadTasks := GetAllUploadTasks(backups)

			t.tasks = append(t.tasks, deletionTasks...)
			t.tasks = append(t.tasks, uploadTasks...)

			// Execute every task queued.
			t.ExecuteTasks(ctx)

			// Sleep for 10 seconds.
			time.Sleep(time.Second * 10)
		}
	}
}


func (t *taskMakerImpl) ExecuteTasks(ctx context.Context) {
	if len(t.tasks) > 0 {
		global.Log("task executor started...")
		// Log string.
		log := "task executor results: "
		for _, task := range t.tasks {
			select {
			case <-ctx.Done():
				global.Log(log)
				global.Log("task executor: context canceled, stopping execution")
				return
			default:

				// Execute the task at hand.
				progressPercent := new(uint)
				err := task.Execute(ctx, progressPercent)

				switch {
				case err != nil:
					log += "\nfail: " + task.Label() + " err:" + err.Error()
				default:
					log += "\nsuccess: " + task.Label()
				}

			}
		}
		global.Log(log)
		t.DumpTasks()
	}
}

func (t *taskMakerImpl) DumpTasks() {
	t.tasks = []QueuedTask{}
}

// Get all backups of every group computer that the currently logged in computer is member of.
func (t *taskMakerImpl) GetAllBackups(ctx context.Context) ([]*models.Backup, error) {
	resp, err := server.Mnimidamon.GroupComputer.GetGroupComputersOfComputer(&group_computer.GetGroupComputersOfComputerParams{
		ComputerID: viewmodels.CurrentComputer.Model.ComputerID,
		Context:    server.ApiContext,
	}, viewmodels.CurrentUser.Auth)

	if err != nil {
		global.Log("task maker: error when getting group computers, %v", err)
		return nil, err
	}

	var backups []*models.Backup
	gcs := resp.Payload
	for _, gc := range gcs {
		select {
		case <-ctx.Done():
			global.Log("task maker: context canceled, stopping backups fetching.")
			return nil, context.Canceled
		default:
			resp, err := server.Mnimidamon.Backup.GetGroupBackups(&backup.GetGroupBackupsParams{
				GroupID: gc.GroupID,
				Context: server.ApiContext,
			}, viewmodels.CurrentComputer.Auth)

			if err != nil {
				global.Log("task maker: error when getting backups of group %v, %v", gc.GroupID, err)
			}

			backups = append(backups, resp.Payload...)
		}
	}

	return backups, nil
}
