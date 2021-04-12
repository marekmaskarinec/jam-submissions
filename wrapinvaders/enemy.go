package main

import (
	"clengine"
	"cliw"
	"math/rand"
	"time"
	"fmt"
)

func EnemyInit(o *clengine.Object) {
	o.Name = "enemy"
	o.Type = "hostile"
	o.Layer.Pos = cliw.Ve2{rand.Intn(40), rand.Intn(76)}
	o.Collider = clengine.AutoCollider(o.Layer.PixMap, o.Hash)
	o.Cycle = EnemyCycle
	global.Enemies[o.Hash] = o

	if !global.WASM {
		o.Layer.PixMap = cliw.LoadPixMap(resourceFolder + "enemy-left.pix")
		enemyTextures =	  map[string][][]string{
			"left": cliw.LoadPixMap(resourceFolder + "enemy-left.pix"),
			"right": cliw.FlipPixMapV(cliw.LoadPixMap(resourceFolder + "enemy-left.pix")),//cliw.LoadPixMap(resourceFolder + "enemy-right.pix"),
		}	
	} else {
		o.Layer.PixMap = global.WebLoad["enemy-left.pix"]
		enemyTextures =	  map[string][][]string{
			"left": global.WebLoad["enemy-left.pix"],
			"right": cliw.FlipPixMapV(global.WebLoad["enemy-left.pix"]),//cliw.LoadPixMap(resourceFolder + "enemy-right.pix"),
		}	
	}

}

var enemyTextures map[string][][]string

func EnemyCycle(o *clengine.Object) {
	var coll string

	for {
		select {
		case <-time.After(time.Duration(global.EnemySpeed) * time.Millisecond):
			o.PrevPos = o.Layer.Pos
			if o.Layer.Pos.X < global.Player.Layer.Pos.X {
				o.Layer.Pos.X++
			} else {
				o.Layer.Pos.X--
			}

			if o.Layer.Pos.Y < global.Player.Layer.Pos.Y {
				o.Layer.Pos.Y++
				o.Layer.PixMap = enemyTextures["right"]
			} else {
				o.Layer.Pos.Y--
				o.Layer.PixMap = enemyTextures["left"]
			}

			if _, exists := o.Root.Objects[o.Hash]; exists {
				coll = o.Update()
			}
			if coll == global.Player.Hash && !global.Invincible {
				if int(global.PlayerHP / 10) > 0 {
					global.PlayerHP--
					global.Invincible = true
					global.InvincibilityStart = int(global.Time * 10)
					o.Root.Objects[coll].Custom["can-animate"] = false
					o.Root.Objects[coll].Custom["animation"] = fmt.Sprintf("hurt-%d", o.Root.Objects[coll].Custom["rotation"].(int))
					o.Root.Objects[coll].Update()
					go Timer(250, func() {o.Root.Objects[global.Player.Hash].Custom["can-animate"] = true; o.Root.Objects[global.Player.Hash].Custom["animation"] = fmt.Sprintf("%d", o.Root.Objects[global.Player.Hash].Custom["rotation"].(int))})
					coll = ""
					o.Layer.Pos = o.PrevPos
				} else {
					o.Root.RenderFunc = DeathScreen
				}
			}
		}
	}
	

	o.Completed = true
}
