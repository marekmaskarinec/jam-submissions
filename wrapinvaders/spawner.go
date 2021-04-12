package main

import (
	"clengine"
	"cliw"
)

func SpawnerInit(o *clengine.Object) {
	o.Name = "spawner"
	o.Type = "nodraw"
	o.Layer.Pos = cliw.Ve2{ 0, 0 }
	o.Cycle = SpawnerCycle
}

func SpawnerCycle(o *clengine.Object) {
	if len(global.Enemies) < int((float64(global.Score) / 256) + 3) {
		o.Root.Add(clengine.Object{Init: EnemyInit})
	}

	global.Time += o.Delta

	if global.Invincible {
		if int(global.Time * 10) - global.InvincibilityStart >= global.InvincibilityLenght {
			global.Invincible = false
		}
	}

	o.Completed = true
}
