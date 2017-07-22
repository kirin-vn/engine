package engine

// Engine is the workhorse for a read-through of a Novel.
// It keeps track of the current page and any state defined by the scenes.
// It exposes information about the current page and scene to callers
// and additionally about state to the scenes.
type Engine struct {
	novel *Novel
	scene string
	page  string
}

// New constructs a new engine for reading through the Novel.
func New(novel *Novel) *Engine {
	return &Engine{
		novel: novel,
		scene: novel.FirstScene,
		page:  novel.Scenes[novel.FirstScene].FirstPage(),
	}
}

// Name returns the name of the visual novel.
func (e *Engine) Name() string {
	return e.novel.Name
}

func (e *Engine) currentScene() Scene {
	return e.novel.Scenes[e.scene]
}

// CurrentPage returns the page that the reader is currently reading.
func (e *Engine) CurrentPage() Page {
	return e.currentScene().GetPage(e.page)
}

// AtEnding tests whether the read-through has reached an ending to the story.
func (e *Engine) AtEnding() bool {
	return e.CurrentPage().NextPage() == "" && e.currentScene().NextScene() == ""
}

// GoToNextPage moves to the next page. If the page is the end of its scene,
// it will also transition to the next scene. It must not be called when
// AtEnding() is true; in this case it will panic.
func (e *Engine) GoToNextPage() {
	if e.AtEnding() {
		panic("Endings don't have a next page.")
	}
	if nextPage := e.CurrentPage().NextPage(); nextPage != "" {
		e.page = e.CurrentPage().NextPage()
	} else {
		e.startScene(e.currentScene().NextScene())
	}
}

func (e *Engine) startScene(scene string) {
	e.scene = scene
	e.page = e.currentScene().FirstPage()
}
