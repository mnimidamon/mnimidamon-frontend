package fragments

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"time"
)

func NewLoaderProcess(name string, refresher func() string) *BackupLoaderProcess {
	infoBind := binding.NewString()

	return &BackupLoaderProcess{
		Name:         name,
		infoBinding:  infoBind,
		canvasObject: container.NewHBox(widget.NewLabel(name), layout.NewSpacer(), widget.NewLabelWithData(infoBind)),
		Refresher:    refresher,
	}
}

type BackupLoaderProcess struct {
	Name         string
	infoBinding  binding.String
	canvasObject fyne.CanvasObject

	parent      *fyne.Container
	Refresher   func() string
	refreshStop context.CancelFunc
}

func (b *BackupLoaderProcess) UpdateInfo(msg string) {
	b.infoBinding.Set(msg)
}

func (b *BackupLoaderProcess) AddToParentContainer(parent *fyne.Container) {
	b.parent = parent
	parent.Add(b.GetCanvasObject())
	parent.Refresh()
	b.StartRefreshing()
}

func (b *BackupLoaderProcess) GetCanvasObject() fyne.CanvasObject {
	return b.canvasObject
}

func (b *BackupLoaderProcess) RemoveFromParentContainer() {
	b.StopRefreshing()
	if b.parent != nil {
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
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if b.parent != nil {
				b.UpdateInfo(b.Refresher())
			}
			time.Sleep(time.Millisecond * 150)
		}
	}
}
