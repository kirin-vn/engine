package engine

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/kirin-vn/lexer"
)

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

// NewNovel initializes a new novel without any scenes defined.
func NewNovel(name, firstScene string) *Novel {
	return &Novel{
		Name:       name,
		FirstScene: firstScene,
		Scenes:     make(map[string]Scene),
	}
}

// ParseScene uses the lexer to parse a scene from an io.Reader
// and add it to the novel.
func (novel *Novel) ParseScene(name, nextScene string, reader io.Reader) error {
	var pageList []string
	for token := range lexer.Tokenize(reader) {
		switch token.Name {
		case lexer.Line:
			pageList = append(pageList, strings.Join(token.Args, ""))
		default:
			log.Printf("Unrecognized kind of token: %s\n", token.Name)
		}
	}
	if len(pageList) == 0 {
		return errors.New("tried to parse empty scene")
	}
	pages := make(map[string]Page)
	var prevID string
	for i := len(pageList) - 1; i >= 0; i-- {
		id := fmt.Sprintf("page-%d", i)
		pages[id] = SimplePage(id, pageList[i], prevID)
		prevID = id
	}
	novel.Scenes[name] = SimpleScene(name, "page-0", pages, nextScene)
	return nil
}

// ParseSceneString parses and adds a new scene from a string.
func (novel *Novel) ParseSceneString(name, nextScene, source string) error {
	reader := strings.NewReader(source)
	return novel.ParseScene(name, nextScene, reader)
}

// Validate validates that the novel references no undefined IDs,
// has a starting scene, and has a starting page for every scene.
func (novel *Novel) Validate() error {
	if !novel.validateSceneID(novel.FirstScene) {
		return errors.New("undefined first scene")
	}
	for _, scene := range novel.Scenes {
		if next := scene.NextScene(); next != "" && !novel.validateSceneID(next) {
			return errors.New("undefined next scene")
		}
		page := scene.FirstPage()
		if page == "" {
			return errors.New("undefined first page")
		}
		for page != "" {
			if !novel.validatePageID(scene, page) {
				return errors.New("undefined next page")
			}
			page = scene.GetPage(page).NextPage()
		}
	}
	return nil
}

func (novel *Novel) validatePageID(scene Scene, page string) bool {
	return scene.GetPage(page) != nil
}

func (novel *Novel) validateSceneID(scene string) bool {
	_, found := novel.Scenes[scene]
	return found
}

// A Scene represents a series of pages/videos/choices/etc. within a novel.
type Scene interface {
	ID() string
	FirstPage() string
	GetPage(string) Page
	NextScene() string
}

type simpleScene struct {
	id        string
	firstPage string
	pages     map[string]Page
	next      string
}

// SimpleScene allows straightforwardly constructing a scene which needs no custom logic.
func SimpleScene(id string, firstPage string, pages map[string]Page, next string) Scene {
	return &simpleScene{
		id:        id,
		firstPage: firstPage,
		pages:     pages,
		next:      next,
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

func (s *simpleScene) NextScene() string {
	return s.next
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
