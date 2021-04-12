package main

import (
	"clengine"
	"time"
)

type Animation struct {
	Frames [][][]string
	FPS    float64
	Link   *clengine.Object
}

func Timer(lenght int, timeout func()) {
	select {
	case <-time.After(time.Duration(lenght) * time.Millisecond):
		timeout()
	}
}

func (a *Animation) Play(o *clengine.Object) {

	if len(a.Frames) == 0 {
		return
	}

	if o.Custom["current"].(int) < len(a.Frames)-1 {
		o.Custom["current"] = o.Custom["current"].(int) + 1
	} else {
		o.Custom["current"] = 0
	}

	if a.Link != nil {
		a.Link.Layer.PixMap = a.Frames[o.Custom["current"].(int)]
		a.Link.Update()
	}

	o.Completed = true
}
