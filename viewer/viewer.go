package viewer

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Init() {
	if err := ui.Init(); err != nil {
		fmt.Println("Failed to start termui")
		fmt.Println(err)
	}
}

func WriteToScreen(name string, author string, album string) {
	p := widgets.NewParagraph()
	p.Text = "Now playing:\n\n" + name + "\nAuthor:" + author + "\nAlbum:" + album
	p.SetRect(0, 0, 50, 8)

	ui.Render(p)
}

func Poll() {
	defer ui.Close()
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
