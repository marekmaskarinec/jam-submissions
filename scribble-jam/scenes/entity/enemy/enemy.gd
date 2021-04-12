extends KinematicBody2D

const SPEED = 256

onready var ray = $RayCast2D
var facing = -1

func die():
	print("dying")
	queue_free()

func _process(_delta):
	var motion = Vector2(SPEED * facing, 20)

	move_and_slide(motion)

	if !ray.is_colliding():
		facing *= -1
		if facing == 1:
			$sprite.flip_h = false
			ray.position = $right_pos.position
		else:
			$sprite.flip_h = true
			ray.position = $left_pos.position


