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

// CurrentPage returns the page that the reader is currently reading.
func (e *Engine) CurrentPage() Page {
	return e.novel.Scenes[e.scene].GetPage(e.page)
}

// AtEnding tests whether the read-through has reached an ending to the story.
func (e *Engine) AtEnding() bool {
	return true // TODO allow page/scene transitions
}
