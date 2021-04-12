extends KinematicBody2D

const SPEED = 600
const MAX_FALL = 1024
const JUMP = 700
const GRAVITY_SCALE = 700

var motion
var gravity = 0
var can_jump = false

func _process(delta):
	motion = Vector2(0, 0)


	if !can_jump and gravity <= MAX_FALL-(GRAVITY_SCALE*delta):
		gravity += GRAVITY_SCALE*delta

	if Input.is_action_pressed("left"):
		motion.x = -SPEED
		$icon.flip_h = true
	if Input.is_action_pressed("right"):
		motion.x = SPEED
		$icon.flip_h = false

	if Input.is_action_just_pressed("up"):
		if can_jump:
			gravity = -JUMP

	motion.y = gravity

	var move = move_and_slide(motion)

	if get_slide_count():
		can_jump = true
	else:
		can_jump = false

