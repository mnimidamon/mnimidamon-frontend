package views

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/client/backup"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func NewBackupsAndInvitedContent(processContainer fyne.CanvasObject) *backupsInvitedContent {
	backupsLabel := widget.NewLabelWithStyle("backups", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	invitesLabel := widget.NewLabelWithStyle("invited", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	membersLabel := widget.NewLabelWithStyle("members", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	backupsToolbarLabel := fragments.NewToolbarObject(backupsLabel)
	invitesToolbarLabel := fragments.NewToolbarObject(invitesLabel)
	membersToolbarLabel := fragments.NewToolbarObject(membersLabel)

	// For reference.
	var bc *backupsInvitedContent

	backupsToolbar := widget.NewToolbar(
		backupsToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Backups.GetAllBackups()
		}),
		widget.NewToolbarAction(resources.DiskSaveSvg, func() {
			bc.dialogCreateNewBackup()
		}),
	)

	membersToolbar := widget.NewToolbar(
		membersToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.GroupMembers.GetAllMembers()
			viewmodels.GroupComputers.GetAllGroupComputers()
		}),
	)

	invitesToolbar := widget.NewToolbar(
		invitesToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.GroupInvitees.GetAllInvitees()
		}),
		widget.NewToolbarAction(resources.EmailPlusSvg, func() {
			dialogInviteUser()
		}),
	)

	leftNavigation := container.NewVBox(
		widget.NewButtonWithIcon("backups", resources.DiskStackSvg, func() {
			bc.DisplayBackupsContent()
		}),
		widget.NewButtonWithIcon("invites", resources.EmailMultipleSvg, func() {
			bc.DisplayInvitesContent()
		}),
		widget.NewButtonWithIcon("members", resources.AccountMultipleSvg, func() {
			bc.DisplayMembersContent()
		}),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("groups", resources.GroupsSvg, func() {
			events.RequestGroupsContent.Trigger()
		}),
	)

	backupsListContainer := container.NewVBox(NewItalicLabel("loading..."))
	invitesListContainer := container.NewVBox(NewItalicLabel("loading..."))
	membersListContainer := container.NewVBox(NewItalicLabel("loading..."))

	rightContent := container.NewMax()

	mainContainer := container.NewBorder(nil, nil, leftNavigation, nil, rightContent)

	bc = &backupsInvitedContent{
		Container:      mainContainer,
		LeftNavigation: leftNavigation,
		RightContent:   rightContent,

		BackupsRightContent: container.NewBorder(backupsToolbar, processContainer, nil, nil, container.NewVScroll(container.NewPadded(backupsListContainer))),
		InvitesRightContent: container.NewBorder(invitesToolbar, processContainer, nil, nil, container.NewVScroll(container.NewPadded(invitesListContainer))),
		MembersRightContent: container.NewBorder(membersToolbar, processContainer, nil, nil, container.NewVScroll(container.NewPadded(membersListContainer))),

		BackupsListContainer: backupsListContainer,
		InvitesListContainer: invitesListContainer,
		MembersListContainer: membersListContainer,
	}

	// Default content is backups.
	bc.DisplayBackupsContent()

	// Register listeners.
	events.GroupInviteesUpdated.Register(bc)
	events.GroupMembersUpdated.Register(bc)
	events.GroupComputersUpdated.Register(bc)
	events.BackupsUpdated.Register(bc)

	return bc
}

type backupsInvitedContent struct {
	Container      *fyne.Container // The encapsulating container.
	LeftNavigation *fyne.Container // Left split content.
	RightContent   *fyne.Container // Right split content.

	BackupsRightContent *fyne.Container // Content displayed upon Invites navigation.
	InvitesRightContent *fyne.Container // Content displayed upon Backups navigation.
	MembersRightContent *fyne.Container // Content displayed upon Members navigation.

	BackupsListContainer *fyne.Container // Containing the backups list.
	InvitesListContainer *fyne.Container // Containing the invites group list.
	MembersListContainer *fyne.Container // Containing the group members list.

	mu sync.Mutex // Lock when rendering UI elements.
}

func (c *backupsInvitedContent) HandleBackupsUpdate() {
	c.rerenderBackups()
}

func (c *backupsInvitedContent) HandleGroupComputersUpdated() {
	c.rerenderMembers()
}

func (c *backupsInvitedContent) HandleGroupMembersUpdated() {
	c.rerenderMembers()
}

func (c *backupsInvitedContent) HandleGroupInviteesUpdated() {
	c.rerenderInvitees()
}

func (c *backupsInvitedContent) DisplayInvitesContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.InvitesRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) DisplayBackupsContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.BackupsRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) rerenderInvitees() {
	c.mu.Lock()
	global.Log("updating invitees list")
	c.InvitesListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.GroupInvitees.Models) == 0 {
		c.InvitesListContainer.Add(NewItalicLabel("There are no pending invites"))
		c.InvitesListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	for _, i := range viewmodels.GroupInvitees.Models {
		c.InvitesListContainer.Add(NewInviteeCanvasObject(i))
	}

	c.InvitesListContainer.Refresh()
	c.mu.Unlock()
}

func (c *backupsInvitedContent) rerenderMembers() {
	global.Log("updating members list")

	c.mu.Lock()
	c.MembersListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.GroupMembers.Models) < 1 {
		c.MembersListContainer.Add(NewItalicLabel("loading..."))
		c.MembersListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	if len(viewmodels.GroupMembers.Models) == 1 {
		c.MembersListContainer.Add(NewItalicLabel("You are the only member of the group"))
	}

	for _, m := range viewmodels.GroupMembers.Models {
		c.MembersListContainer.Add(NewMemberCanvasObject(m))
	}

	c.MembersListContainer.Refresh()
	c.mu.Unlock()
}

func (c *backupsInvitedContent) DisplayMembersContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.MembersRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) rerenderBackups() {
	global.Log("updating backups list")

	c.mu.Lock()
	c.BackupsListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.Backups.Models) == 0 {
		c.BackupsListContainer.Add(NewItalicLabel("There are no backups in this group"))
		c.BackupsListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	for _, b := range viewmodels.Backups.Models {
		c.BackupsListContainer.Add(NewBackupCanvasObject(b))
	}

	c.BackupsListContainer.Refresh()
	c.mu.Unlock()
}

func NewBackupCanvasObject(b *models.Backup) fyne.CanvasObject {
	return widget.NewLabel(b.Filename)
}

func NewMemberCanvasObject(m *models.User) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle(m.Username, fyne.TextAlignLeading, fyne.TextStyle{}),
		NewComputersCanvasObject(m.UserID),
	)
}

func NewComputersCanvasObject(userID int64) fyne.CanvasObject {
	c := container.NewVBox()

	for _, gc := range viewmodels.GroupComputers.GetAllOf(userID) {
		c.Add(NewGroupComputerCanvasObject(gc))
	}

	return c
}

func NewGroupComputerCanvasObject(gc *models.GroupComputer) fyne.CanvasObject {
	return widget.NewLabel(fmt.Sprintf("   %v  \t%vMB", gc.Computer.Name, gc.StorageSize/1024))
}

func NewInviteeCanvasObject(i *models.Invite) fyne.CanvasObject {
	return container.NewHBox(
		widget.NewLabel(fmt.Sprintf("%v @ %v", i.User.Username, i.Date)),
	)
}

func (c *backupsInvitedContent) dialogCreateNewBackup() {
	// Create a random key and display it.
	key, err := services.NewRandomEncryptionKey()
	if err != nil {
		infoDialog("could not create random encryption key: " + err.Error())
		return
	}
	keyEntry := widget.NewEntry()
	keyEntry.Disable()
	keyEntry.SetText(hex.EncodeToString(key))

	// File entry.
	fileEntry := widget.NewEntry()
	fileEntry.SetPlaceHolder("path to the file..")
	fileEntry.Validator = func(s string) error {
		if _, err := os.Stat(s); os.IsNotExist(err) {
			return errors.New("file does not exist")
		}
		return nil
	}

	// The dialog to get the file selection.
	selectFileDialog := dialog.NewFileOpen(func(uri fyne.URIReadCloser, err error) {
		if uri != nil {
			fileEntry.SetText(uri.URI().Path())
			uri.Close()
		}
	}, global.MainWindow)

	// Button for file selection dialog show.
	buttonSelectFolder := widget.NewButtonWithIcon("Select file", resources.FolderOpenSvg, func() {
		selectFileDialog.Show()
	})

	// Combined dialog for creating new backup.
	dialog.NewForm("Backup a new file inside the group "+viewmodels.SelectedGroup.Model.Name, "Back it up", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("File", fileEntry),
			widget.NewFormItem("", buttonSelectFolder),
			widget.NewFormItem("Decryption key", keyEntry),
			// TODO copy key
		}, func(b bool) {
			if b {
				path := fileEntry.Text
				global.Log("requested backup creation of %v", fileEntry.Text)

				// Start the procedure.
				c.BackupUploadProcedure(path, key)
			}
		}, global.MainWindow).Show()
}

func (c *backupsInvitedContent) BackupUploadProcedure(path string, key services.EncryptionKey) {

	// Open the file.
	file, err := os.Open(path)
	if err != nil {
		infoDialog("could not open file: " + err.Error())
		return
	}

	// File stats to get the size.
	fi, err := file.Stat()
	if err != nil {
		infoDialog("could not get file info: " + err.Error())
		return
	}

	// Get the file name.
	fname := filepath.Base(path)
	payload := &models.InitializeGroupBackupPayload{
		FileName: &fname,
		Hash:     new(string),
		Size:     new(int64),
	}

	// Number of bytes encrypted.
	readBytes := new(int)

	// Shorten the name if its too long.
	loaderName := *payload.FileName
	if len(loaderName) > 15 {
		loaderName = loaderName[:15] + "..."
	}

	// New backup loader process for UI.
	bl := NewBackupLoaderProcess(loaderName, func() string {
		// Calculate the percentage based on the size and the pointer value of bytes already processed.
		percentage := int(float64(*readBytes) / float64(fi.Size()) * 100)
		if percentage == 100 {
			return "Encrypted"
		}
		return "Encrypting " + strconv.Itoa(percentage) + "% ..."
	})

	// Inform about a new process.
	events.ProcessStarted.Trigger(bl)

	go func() {
		// Remove the process from the parent container.
		defer bl.RemoveFromParentContainer()
		defer time.Sleep(time.Second)

		// Presave the groupID in case it gets switched up inbetween processing.
		groupID := viewmodels.SelectedGroup.Model.GroupID

		// Encrypt and get the encrypted file. Payload will also be populated with the size and the hash.
		encryptedFile, err := services.BackupCryptography.Encrypt(payload, key, file, readBytes)
		if err != nil {
			infoDialog("error encrypting file: " + err.Error())
			return
		}

		// Defer clean the temp folder of that encrypted file.
		defer services.BackupStorage.DeleteTempFile(*payload.FileName)
		defer encryptedFile.Close()


		// Encryption is complete.
		global.Log("encryption complete %v %v", encryptedFile.Name(), payload)

		// Start the backup initialization on the server.
		bl.StopRefreshing()
		time.Sleep(time.Second)
		bl.UpdateInfo("Initializing...")
		time.Sleep(time.Second * 2)

		global.Log("backup init payload %v", payload)

		// Request the initialization on the server.
		respInit, err := server.Mnimidamon.Backup.InitializeGroupBackup(&backup.InitializeGroupBackupParams{
			Body:       payload,
			GroupID:    groupID,
			Context:    server.ApiContext,
		}, viewmodels.CurrentComputer.Auth)

		// Check if error occurred.
		if err != nil {
			infoDialog(err.Error())
			bl.UpdateInfo("Initialization error, cancelling.")
			return
		}


		bl.UpdateInfo("Initialized")
		time.Sleep(time.Second * 2)
		bl.UpdateInfo("Uploading...")
		global.Log("backup init response %v", respInit.Payload)

		// Add it to the view models.
		viewmodels.Backups.Add(respInit.Payload)

		// Upload it to the server.
		respUpload, err := server.Mnimidamon.Backup.UploadBackup(&backup.UploadBackupParams{
			BackupData: encryptedFile,
			BackupID:   respInit.Payload.BackupID,
			GroupID:    groupID,
			Context:    server.ApiContext,
		}, viewmodels.CurrentComputer.Auth)


		if err != nil {
			infoDialog(err.Error())
			bl.UpdateInfo("Upload error, cancelling.")
			return
		}

		time.Sleep(time.Second * 1)
		bl.UpdateInfo("Uploaded")
		global.Log("upload response %v", respUpload.Payload)

		// Move the file from the temp folder to the main folder.
		if err := services.BackupStorage.MoveFromTemp(*payload.FileName, respInit.Payload.BackupID); err != nil {
			infoDialog(err.Error())
			bl.UpdateInfo("File moving error, cancelling.")
			return
		}

		time.Sleep(time.Second * 1)
		bl.UpdateInfo("Successful")
	}()
}

func dialogInviteUser() {
	nameEntry := widget.NewEntry()
	nameEntry.Validator = func(s string) error {
		if len(s) < 3 {
			return errors.New("at least 3 characters long")
		}
		return nil
	}

	dialog.NewForm("Invite user to "+viewmodels.SelectedGroup.Model.Name, "Send", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("Name", nameEntry),
		}, func(b bool) {
			if b {
				inviteUserToGroup(nameEntry.Text)
			}
		}, global.MainWindow).Show()
}

func inviteUserToGroup(name string) {
	go func() {
		resp, err := server.Mnimidamon.Group.InviteUserToGroup(&group.InviteUserToGroupParams{
			Body:    &models.InviteUserPayload{Username: &name},
			GroupID: viewmodels.SelectedGroup.Model.GroupID,
			Context: server.ApiContext,
		}, viewmodels.CurrentComputer.Auth)

		if err != nil {
			if br, ok := err.(*group.InviteUserToGroupBadRequest); ok {
				infoDialog(br.GetPayload().Message)
				return
			}
			infoDialog(err.Error())
			return
		}

		// Add the created invitations to the group invitees.
		viewmodels.GroupInvitees.Add(resp.Payload)
	}()
}

func NewItalicLabel(msg string) *widget.Label {
	return widget.NewLabelWithStyle(msg, fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
}

func NewBackupLoaderProcess(name string, refresher func() string) *BackupLoaderProcess {
	infoBind := binding.NewString()

	return &BackupLoaderProcess{
		Name:         name,
		infoBinding:  infoBind,
		canvasObject: container.NewHBox(widget.NewLabel(name), layout.NewSpacer(), widget.NewLabelWithData(infoBind)),
		refresher:    refresher,
	}
}

type BackupLoaderProcess struct {
	Name         string
	infoBinding  binding.String
	canvasObject fyne.CanvasObject

	parent      *fyne.Container
	refresher   func() string
	refreshStop context.CancelFunc
}

func (b *BackupLoaderProcess) UpdateInfo(msg string) {
	b.infoBinding.Set(msg)
}

func (b *BackupLoaderProcess) AddToParentContainer(parent *fyne.Container) {
	b.parent = parent
	parent.Add(b.GetCanvasObject())
	b.StartRefreshing()
}

func (b *BackupLoaderProcess) GetCanvasObject() fyne.CanvasObject {
	return b.canvasObject
}

func (b *BackupLoaderProcess) RemoveFromParentContainer() {
	if b.parent != nil {
		b.StopRefreshing()
		go func() {
			time.Sleep(time.Millisecond * 1500)
			b.parent.Remove(b.canvasObject)
			b.parent.Refresh()
			b.parent = nil
		}()
	}
}
func (b *BackupLoaderProcess) StopRefreshing() {
	if b.refreshStop != nil {
		b.refreshStop()
		b.refreshStop = nil
	}
}

func (b *BackupLoaderProcess) StartRefreshing() {
	if b.refreshStop != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	b.refreshStop = cancel
	go b.RefreshRoutine(ctx)
}

func (b *BackupLoaderProcess) RefreshRoutine(ctx context.Context) {
	global.Log("task ui refreshing started...")
	for {
		select {
		case <-ctx.Done():
			global.Log("stopped task refreshing, context canceled")
			return
		default:
			if b.parent != nil {
				b.UpdateInfo(b.refresher())
			}
			time.Sleep(time.Millisecond * 150)
		}
	}
}
