package engine

// A Novel is a data structure representing a visual novel.
// It's the script of the game. It defines what scenes there are,
// what pages and those scenes contain,
// what text and images and videos and sounds and choices those pages contain,
// how those scenes transition from one to another,
// what impact choices have, etc.
// It doesn't track any of the state of a particular read-through
// of the novel. Just what the novel itself is.
type Novel struct {
	Name       string
	FirstScene string
	Scenes     map[string]Scene
}

// A Scene represents a series of pages/videos/choices/etc. within a novel.
type Scene interface {
	ID() string
	FirstPage() string
	GetPage(string) Page
}

type simpleScene struct {
	id        string
	firstPage string
	pages     map[string]Page
}

// SimpleScene allows straightforwardly constructing a scene which needs no custom logic.
func SimpleScene(id string, firstPage string, pages map[string]Page) Scene {
	return &simpleScene{
		id:        id,
		firstPage: firstPage,
		pages:     pages,
	}
}

func (s *simpleScene) ID() string {
	return s.id
}

func (s *simpleScene) FirstPage() string {
	return s.firstPage
}

func (s *simpleScene) GetPage(id string) Page {
	return s.pages[id]
}

// A Page represents a single "page" within a scene, which may display
// some subset of text, foreground and background images, animations,
// videos, and audio.
type Page interface {
	ID() string
	Text() string
	NextPage() string
}

type simplePage struct {
	id   string
	text string
	next string
}

// SimplePage allows straightforwardly constructing a page which needs no custom logic.
func SimplePage(id string, text string, next string) Page {
	return &simplePage{
		id:   id,
		text: text,
		next: next,
	}
}

func (pg *simplePage) ID() string {
	return pg.id
}

func (pg *simplePage) Text() string {
	return pg.text
}

func (pg *simplePage) NextPage() string {
	return pg.next
}
