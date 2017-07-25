package main

import (
	"fmt"

	"github.com/kirin-vn/engine"
)

func main() {
	novel := testNovel()
	engine, err := engine.New(novel)
	if err != nil {
		panic(err)
	}
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
	novel := engine.NewNovel("Test Novel", "first-scene")
	if err := novel.ParseSceneString("first-scene", "second-scene", sceneOne); err != nil {
		panic(err)
	}
	if err := novel.ParseSceneString("second-scene", "", sceneTwo); err != nil {
		panic(err)
	}
	return novel
}
