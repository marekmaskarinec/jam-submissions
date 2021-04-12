extends Node2D

var plcount = 16
var plsize = 64 * Generator.platform_scale - 1

func _ready():
	var inst
	var res = load("res://scenes/env/platform/platform.tscn")

	for i in range(plcount):
		for j in range(plcount):
			inst = res.instance()
			inst.position = Vector2(i, j+i%2) * plsize
			inst.position += Vector2(randi()%256-128, randi()%256-128)
			inst.position.y /= 1.1
			$platforms.add_child(inst)
  
