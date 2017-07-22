package main

import (
	"fmt"

	"github.com/kirin-vn/engine"
)

func main() {
	novel := testNovel()
	engine := engine.New(novel)
	fmt.Printf("Playing %s\n\n", engine.Name())
	for {
		page := engine.CurrentPage()
		fmt.Println(page.Text())
		if engine.AtEnding() {
			break
		}
		fmt.Scanln() // just to wait for a line
	}
}

func testNovel() *engine.Novel {
	return &engine.Novel{
		Name:       "Test Novel",
		FirstScene: "only-scene",
		Scenes: map[string]engine.Scene{
			"only-scene": engine.SimpleScene(
				"only-scene", "only-page",
				map[string]engine.Page{
					"only-page": engine.SimplePage("only-page", "This is the single page of the VN."),
				}),
		},
	}
}
