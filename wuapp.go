package wuapp

// Settings is to configure the window's appearance
type WindowSettings struct {
	Title          string //Title of the application window
	UIDir          string //Directory of the UI/Web related files, default: "ui"
	Index          string //Index html file, default: "index.html"
	Url            string //Full url address if you don't use WebDir + Index
	Left           int
	Top            int
	Width          int
	Height         int
	Resizable      bool
	Closable       bool
	Miniaturizable bool
	Borderless     bool
	FullScreen     bool
	Debug          bool
}

type Widget interface {
	Register()
}

func Settings(settings WindowSettings) {

}

func Run(settings WindowSettings) (err error) {
	err = create(settings)
	defer exit()

	return
}

func RegisterWidgets(widgets ...Widget) {
	for _, w := range widgets {
		w.Register()
	}
}
