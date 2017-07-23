package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kirin-vn/engine"
	"github.com/kirin-vn/lexer"
)

func main() {
	novel := testNovel()
	engine := engine.New(novel)
	fmt.Printf("Playing %s\n\n", engine.Name())
	for {
		page := engine.CurrentPage()
		fmt.Println(page.Text())
		if engine.AtEnding() {
			fmt.Println()
			break
		}
		fmt.Scanln() // just to wait for a line
		engine.GoToNextPage()
	}
}

const (
	sceneOne = "This is the first page of the VN.\n" +
		"This is the second and last page of the first scene.\n"
	sceneTwo = "This is the only page of the second scene.\n"
)

func testNovel() *engine.Novel {
	return &engine.Novel{
		Name:       "Test Novel",
		FirstScene: "first-scene",
		Scenes: map[string]engine.Scene{
			"first-scene":  readScene("first-scene", "second-scene", sceneOne),
			"second-scene": readScene("second-scene", "", sceneTwo),
		},
	}
}

func readScene(id string, next string, source string) engine.Scene {
	reader := strings.NewReader(source)
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
		panic("tried to parse empty scene")
	}
	pages := make(map[string]engine.Page)
	var prevID string
	for i := len(pageList) - 1; i >= 0; i-- {
		id := fmt.Sprintf("page-%d", i)
		pages[id] = engine.SimplePage(id, pageList[i], prevID)
		prevID = id
	}
	return engine.SimpleScene(id, "page-0", pages, next)
}
