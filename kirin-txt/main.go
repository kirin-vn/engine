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
			fmt.Println()
			break
		}
		fmt.Scanln() // just to wait for a line
		engine.GoToNextPage()
	}
}

func testNovel() *engine.Novel {
	return &engine.Novel{
		Name:       "Test Novel",
		FirstScene: "only-scene",
		Scenes: map[string]engine.Scene{
			"only-scene": engine.SimpleScene(
				"only-scene", "first-page",
				map[string]engine.Page{
					"first-page":  engine.SimplePage("first-page", "This is the first page of the VN.", "second-page"),
					"second-page": engine.SimplePage("second-page", "This is second and last page of the VN.", ""),
				}),
		},
	}
}
