package ui

import (
	"github.com/galaco/lambda-client/engine"
	vguiCore "github.com/galaco/lambda-core/loader/vgui"
	"github.com/galaco/lambda-core/vgui"
	"github.com/galaco/tinygametools"
	"github.com/golang-source-engine/filesystem"
)

type Gui struct {
	engine.Manager
	window      *tinygametools.Window
	masterPanel vgui.MasterPanel
}

func (ui *Gui) Register() {

}

func (ui *Gui) Update(dt float64) {
	ui.Render()
}

func (ui *Gui) Render() {
	ui.masterPanel.Draw()
}

// LoadVGUIResource
func (ui *Gui) LoadVGUIResource(fs *filesystem.FileSystem, filename string) error {
	p, err := vguiCore.LoadVGUI(fs, filename)
	if err != nil {
		return err
	}
	ui.masterPanel.AddChild(p)

	return nil
}

func (ui *Gui) MasterPanel() *vgui.MasterPanel {
	return &ui.masterPanel
}

func NewGUIManager(win *tinygametools.Window) *Gui {
	return &Gui{
		window: win,
	}
}
