package main

import (
	"cliw"
	"fmt"
	"clengine"
	"strings"
)

var font map[string]interface{}

func AutoCenter(text string, screenSize, fontSize int) string {
	return strings.Repeat(" ", ((screenSize/fontSize)-len(text))/2) + text
}

func DeathScreen(s *clengine.Scene) {
	var tr [][]string
	tr = append(tr, cliw.TextToPixMap(font, "c7f0d8", "", "you died", cliw.Ve2{3, 3})...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, cliw.TextToPixMap(font, "c7f0d8", "", fmt.Sprintf("score:%d", global.Score), cliw.Ve2{3, 3})...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, cliw.TextToPixMap(font, "c7f0d8", "", "restart the", cliw.Ve2{3, 3})...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, cliw.TextToPixMap(font, "c7f0d8", "", "game to retry", cliw.Ve2{3, 3})...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, cliw.TextToPixMap(font, "c7f0d8", "", "q:quit", cliw.Ve2{3, 3})...)

	if !global.WASM {
		s.Renderer.PixMap = cliw.ReturnWithPixLayers(cliw.LoadPixMap(resourceFolder + "death-screen.pix"), []cliw.PixLayer{cliw.PixLayer{PixMap: tr, Pos: cliw.Ve2{ 12, 10 }}})
	} else {
		s.Renderer.PixMap = cliw.ReturnWithPixLayers(global.WebLoad["death-screen.pix"], []cliw.PixLayer{cliw.PixLayer{PixMap: tr, Pos: cliw.Ve2{ 12, 10 }}})
	}
	s.Renderer.Render(&s.Renderer)
}

func Menu() (tr [][]string) {
	text := cliw.TextToPixMap(font, "c7f0d8", "", AutoCenter("wrap invaders", 84, 4), cliw.Ve2{3, 3})
	text = append(text, [][]string{make([]string, 1)}...)
	text = append(text, [][]string{make([]string, 1)}...)
	text = append(text, [][]string{make([]string, 1)}...)
	text = append(text, [][]string{make([]string, 1)}...)
	text = append(text, cliw.TextToPixMap(font, "c7f0d8", "", AutoCenter("press enter to play", 84, 4), cliw.Ve2{3, 3})...)

	if global.WASM {
		return cliw.ReturnWithPixLayers(global.WebLoad["death-screen.pix"], []cliw.PixLayer{ cliw.PixLayer{ PixMap: text, Pos: cliw.Ve2{ (48-11)/2, 0 } } })
	}
	return cliw.ReturnWithPixLayers(cliw.LoadPixMap(resourceFolder + "death-screen.pix"), []cliw.PixLayer{ cliw.PixLayer{ PixMap: text, Pos: cliw.Ve2{ (48-11)/2, 0 } } })
}

func PrepareUI() (tr [][]string) {

	tr = append(tr, cliw.TextToPixMap(font, "43523d", "", fmt.Sprintf("hp:%d", global.PlayerHP/10), cliw.Ve2{3, 3})...)
	tr = append(tr, [][]string{make([]string, 1)}...)
	tr = append(tr, cliw.TextToPixMap(font, "43523d", "", fmt.Sprintf("score:%d", global.Score), cliw.Ve2{3, 3})...)

	return
}

func RenderFunc(s *clengine.Scene) {
	s.Renderer.PixMap = cliw.ReturnWithPixLayers(cliw.DuplicatePix(s.WorldPix), []cliw.PixLayer{ cliw.PixLayer{ Pos: cliw.Ve2{ 1, 1 }, PixMap: PrepareUI() } })
	s.Renderer.Render(&s.Renderer)
}
