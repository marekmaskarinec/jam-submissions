extends TextEdit

var mode = 0
var commands

func _ready():
	var save_game = File.new()
	save_game.open("res://etc/commands.json", File.READ)
	commands = parse_json(save_game.get_as_text())

func _process(_delta):
	if Input.is_action_just_pressed("insert_mode"):
		mode = 1
		visible = true	
		get_tree().paused = true

	if Input.is_action_just_pressed("normal_mode"):
		mode = 0
		visible = false
		get_tree().paused = false

	if Input.is_action_just_pressed("enter"):
			on_enter(text)

func on_enter(input):
	if not input in commands:
		return

	var inst = load(commands[input]).instance()
	get_node("/root/game/commands").add_child(inst)	


func _on_CommandBox_text_changed():
	if len(text) != 0 and text[len(text)-1] == "\n":
		print("entering")
		on_enter(text)
