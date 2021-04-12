package main

import (
	"clengine"
	"cliw"
	"fmt"
)

func PlayerInit(o *clengine.Object) {
	o.Cycle = PlayerCycle
	o.Name = "player"
	o.Layer.Pos = cliw.Ve2{16, 36}
	o.Layer.PixMap = cliw.LoadPixMap(resourceFolder + "player-left.pix")
	if !global.WASM {
		playerFrames = map[string][][][]string{
			"-1": [][][]string{ cliw.LoadPixMap(resourceFolder + "player-walk/player-walk-0.pix"), cliw.LoadPixMap(resourceFolder + "player-walk/player-walk-1.pix") },
			"1": [][][]string{ cliw.FlipPixMapV(cliw.LoadPixMap(resourceFolder + "player-walk/player-walk-0.pix")), cliw.FlipPixMapV(cliw.LoadPixMap(resourceFolder + "player-walk/player-walk-1.pix")) },
			"hurt--1": [][][]string{ cliw.LoadPixMap(resourceFolder + "player-hit/player-hit-0.pix"), cliw.LoadPixMap(resourceFolder + "player-hit/player-hit-1.pix") },
			"hurt-1": [][][]string{ cliw.FlipPixMapV(cliw.LoadPixMap(resourceFolder + "player-hit/player-hit-0.pix")), cliw.FlipPixMapV(cliw.LoadPixMap(resourceFolder + "player-hit/player-hit-1.pix")) },
			"idle-left": [][][]string{ cliw.LoadPixMap(resourceFolder + "player-idle/player-idle-0.pix"), cliw.LoadPixMap(resourceFolder + "player-idle/player-idle-1.pix") },
			"idle-right": [][][]string{ cliw.FlipPixMapV(cliw.LoadPixMap(resourceFolder + "player-idle/player-idle-0.pix")), cliw.FlipPixMapV(cliw.LoadPixMap(resourceFolder + "player-idle/player-idle-1.pix")) },
		}
	} else {
		playerFrames = map[string][][][]string{
			"-1": [][][]string{ global.WebLoad["player-walk/player-walk-0.pix"],global.WebLoad["player-walk/player-walk-1.pix"] },
			"1": [][][]string{ cliw.FlipPixMapV(global.WebLoad["player-walk/player-walk-0.pix"]), cliw.FlipPixMapV(global.WebLoad["player-walk/player-walk-1.pix"]) },
			"hurt--1": [][][]string{ global.WebLoad["player-hit/player-hit-0.pix"], global.WebLoad["player-hit/player-hit-0.pix"] },
			"hurt-1": [][][]string{ cliw.FlipPixMapV(global.WebLoad["player-hit/player-hit-0.pix"]), cliw.FlipPixMapV(global.WebLoad["player-hit/player-hit-1.pix"]) },
			"idle-left": [][][]string{ global.WebLoad["player-idle/player-idle-0.pix"], global.WebLoad["player-idle/player-idle-1.pix"] },
			"idle-right": [][][]string{ cliw.FlipPixMapV(global.WebLoad["player-idle/player-idle-0.pix"]), cliw.FlipPixMapV(global.WebLoad["player-idle/player-idle-1.pix"]) },
		}
	}
	o.Collider = clengine.AutoCollider(o.Layer.PixMap, o.Hash)
	o.Custom = map[string]interface{}{}
	o.Custom["rotation"] = -1
	o.Custom["animation"] = "walk-right"
	o.Custom["can-animate"] = true
	global.Player = o
}

var playerFrames map[string][][][]string
var framerates map[string]float64

func PlayerCycle(o *clengine.Object) {


	speed := int(1.0/32 * o.Delta * 1000)

	// input handling
	switch o.Root.KeyPress {
	case "w":
		o.PrevPos = o.Layer.Pos
		if o.Layer.Pos.X > speed {
			o.Layer.Pos.X -= speed
		} else {
			o.Layer.Pos.X = 39 - speed
		}
		//o.Root.KeyPress = ""
		global.Invincible = false
		if o.Custom["can-animate"].(bool) {
			o.Custom["animation"] = fmt.Sprintf("%d", o.Custom["rotation"].(int))
		}
		o.Update()
		
	case "a":
		o.PrevPos = o.Layer.Pos
		if o.Layer.Pos.Y > speed {
			o.Layer.Pos.Y -= speed
		} else {
			o.Layer.Pos.Y = 75 - speed
		}
		//o.Root.KeyPress = ""
		if o.Custom["rotation"].(int) != -1 {
			if o.Custom["can-animate"].(bool) {
				o.Custom["animation"] = "-1"
			}
			o.Custom["rotation"] = -1
		}

		global.Invincible = false

		o.Update()

	case "s":
		o.PrevPos = o.Layer.Pos
		if o.Layer.Pos.X < 38 - speed {
			o.Layer.Pos.X += speed
		} else {
			o.Layer.Pos.X = speed
		}
		//o.Root.KeyPress = ""

		global.Invincible = false
		if o.Custom["can-animate"].(bool) {
			o.Custom["animation"] = fmt.Sprintf("%d", o.Custom["rotation"].(int))
		}

		o.Update()

	case "d":
		o.PrevPos = o.Layer.Pos
		if o.Layer.Pos.Y < 74 - speed {
			o.Layer.Pos.Y += speed
		} else {
			o.Layer.Pos.Y = 0//speed
		}
		//o.Root.KeyPress = ""
		if o.Custom["rotation"].(int) != 1 {
			if o.Custom["can-animate"].(bool) {
				o.Custom["animation"] = "1"
			}
			o.Custom["rotation"] = 1
		}

		global.Invincible = false
		
		o.Update()
	case "e":
		o.Root.KeyPress = ""
		if o.Layer.Pos.Y > 1 && o.Layer.Pos.Y < 82 {
			c := map[string]interface{}{}
			c["rotation"] = o.Custom["rotation"].(int)
			o.Root.Add(
				clengine.Object{
					Init: BulletInit,
					Custom: c,
				},
			)
		}

	case "q":
		o.Root.Active = false
	case "":
		if o.Custom["can-animate"].(bool) {
			if o.Custom["rotation"].(int) == 1 {
				o.Custom["animation"] = "idle-right"
			} else {
				o.Custom["animation"] = "idle-left"
			}
			o.Update()
		}

	}

	o.Layer.PixMap = playerFrames[o.Custom["animation"].(string)][int(global.Time * 5) % (len(playerFrames[o.Custom["animation"].(string)]))]
	
	o.Completed = true
}
