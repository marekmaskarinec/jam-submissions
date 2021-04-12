package main

import (
	"clengine"
	"cliw"
	"time"
)

func BulletInit(o *clengine.Object) {
	o.Name = "bullet"
	if !global.WASM {
		o.Layer.PixMap = cliw.LoadPixMap(resourceFolder + "bullet.pix")
	} else {
		o.Layer.PixMap = global.WebLoad["player-walk/enemy-left.pix"]
	}
	o.Layer.Pos = cliw.Ve2{
					global.Player.Layer.Pos.X + 3,
					(global.Player.Layer.Pos.Y + 3)	+ 4 * o.Custom["rotation"].(int),
	}
	
	o.Custom["rotation"] = global.Player.Custom["rotation"].(int)
	o.Collider = clengine.AutoColliderColor(o.Layer.PixMap, []string{"43523d"}, o.Hash)

	go BulletCycle(o)
}


func BulletCycle(o *clengine.Object) {
	for i:=0; i < global.BulletLifetime; i++ {
		select{
		case <-time.After(time.Duration(global.BulletSpeed) * time.Millisecond):
		
			o.PrevPos = o.Layer.Pos
			if o.Layer.Pos.Y + o.Custom["rotation"].(int) * 2 < 1 && o.Custom["rotation"].(int) == -1 {
				o.Layer.Pos.Y = 82
			} else if o.Layer.Pos.Y + o.Custom["rotation"].(int) * 2 > 82 && o.Custom["rotation"].(int) == 1 {
				o.Layer.Pos.Y = 0
			} else {
				o.Layer.Pos.Y += o.Custom["rotation"].(int) * 2
			}

			coll := o.Update()
			if coll != "" {
				if e, exists := o.Root.Objects[coll]; exists && e.Type == "hostile" {
					delete(global.Enemies, coll)
					o.Root.Delete(coll)
					o.Root.Changed = true
					global.Score += 10
					o.Root.Delete(o.Hash)
					return
				}
			}
		}
		
	}
	o.Root.Delete(o.Hash)
	o.Completed = true
}
