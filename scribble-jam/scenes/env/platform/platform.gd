extends Node2D

var platform
var book
var enemies = []

func _ready():
	randomize()
	platform = []

	# creates blank platform
	for i in range(Generator.platform_scale):
		platform.append([])
		for j in range(Generator.platform_scale):
			platform[i].append(false)

	wallgen(plgen())
	draw()

	add_enemies()

	if book != null:
		$TileMap.set_cell(book.y, book.x, 0)

func add_enemies():
	var enemy = load("res://scenes/entity/enemy/enemy.tscn")
	var inst
	for i in range(len(enemies)):
		inst = enemy.instance()
		inst.position = Vector2(enemies[i].y, enemies[i].x) * 64
		add_child(inst)

# draws the platform to tilemap
func draw():
	for i in range(len(platform)):
		for j in range(len(platform[i])):
			if platform[i][j]:
				$TileMap.set_cell(j, i, 1)

	$TileMap.update_bitmask_region()

# adds elevation to platform
func wallgen(plsize):
	var border = (Generator.platform_scale-plsize)/2
	var wall = -1
	var rng
	var last_size

	# runs through the platform and decides starting point
	for i in range(plsize):
		if randf() <= 0.8:
			wall = i
			break

	if wall == -1:
		return
	
	last_size = randi()%(Generator.platform_scale-6)+2
	# if there is a wall, runs through wall to platform size
	while wall < plsize-border-1 and wall < Generator.platform_scale:
		# decides, if it should exit the generator
		rng = randi()%4
		if rng == 1:
			last_size += 4*(randi()%3-1)

		if last_size >= Generator.platform_scale-2:
			break

		rng = randf()
		if rng <= 0.02:
			book = Vector2(len(platform)-2-last_size, border+wall)
		if rng >= 0.9:
			enemies.append(Vector2(len(platform)-2-last_size, border+wall))

		# decides, the elevation change
		#rng = (randi()%6)/2-1
		for i in range(last_size):
			platform[len(platform)-i-2][border+wall] = true
		wall += 1
  

# generates the base platform
func plgen():
	var rng = randi()%Generator.platform_scale+2
	var secondrow = randi()%2

	var i = (Generator.platform_scale-rng)/2
	while i < Generator.platform_scale - int((Generator.platform_scale-rng)/2) -1:
		platform[Generator.platform_scale-2][i] = true
		
		if secondrow:
			platform[Generator.platform_scale-1][i] = true
	
		i += 1

	return rng
