package main

import (
	cr "clengine/wasm"
	"clengine"
)

func ClengineInit(s *clengine.Scene) {
	cr.Init(s)
}

func ClengineClose() {
	cr.Close()
}
