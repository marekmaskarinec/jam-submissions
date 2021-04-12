extends RigidBody2D

func _on_Timer_timeout():
	var ob = $Area2D.get_overlapping_bodies()
	
	for i in range(len(ob)):
		if ob[i].is_in_group("enemies"):
			ob[i].die()
