package main

import (
	"clengine"
	"cliw"
	"time"
	"fmt"
	"runtime"
	"os"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"math/rand"
)


type Global struct {
	BulletSpeed int
	BulletLifetime int
	Player *clengine.Object
	MaxEnemies int
	PlayerHP int
	Score int
	Enemies map[string]*clengine.Object
	EnemySpeed float64
	Invincible bool
	InvincibilityStart int
	InvincibilityLenght int
	Time float64
	WASM bool
	WebLoad map[string][][]string
}

var global Global
var resourceFolder string

func main() {
	if len(os.Args) > 1 && os.Args[1] == "controls" {
		fmt.Println("Controls\nYou can move using `wasd`. You don't have to hold the key to move, since the player will keep moving. To stop him from moving, press any other key.\n To shoot, press the `e` key. Shooting will also stop all movement.\n You can quit from the game at any time using `q`. Enjoy!")
		return
	}

	fmt.Println("Welcome to wrap invaders. To open controls, run this game with the `controls flag`. Please report bugs on github.com/marekmaskarinec/nokia-jam/issues. Thank you :]")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		resourceFolder = "assets/"
	} else {
		switch runtime.GOOS {
		case "linux":
			resourceFolder = os.Getenv("HOME") + "/.config/wrapinvaders/assets/"
		case "windows":
			resourceFolder = "assets/"
		case "js":
			resourceFolder = "https://marekmaskarinec.github.io/files/nokia-jam/"
			global.WASM = true
		default:
			fmt.Println("This game supports only windows and linux. Did you build it from source, or where did you even get the binary? :]")
			fmt.Println("Since you obtained the binary, I'm gonna let you play it in compatibility mode.")
			resourceFolder = "assets/"
		}
	}

	if !global.WASM {
		font = cliw.LoadFont(resourceFolder + "font.json")
	} else {
		toLoad := []string{
			"splash-screen.pix",
			"bg.pix",
			"bullet.pix",
			"player-walk/player-walk-0.pix",
			"player-walk/player-walk-1.pix",
			"player-idle/player-idle-0.pix",
			"player-idle/player-idle-1.pix",
			"player-hit/player-hit-0.pix",
			"player-hit/player-hit-1.pix",
			"enemy-left.pix",
			"death-screen.pix",
		}
		loaded := map[string][][]string{}

		for _, curr := range toLoad {
			resp, _ := http.Get(resourceFolder + curr)
			defer resp.Body.Close()

			dat, _ := ioutil.ReadAll(resp.Body)
			tmp := [][]string{}
			json.Unmarshal(dat, &tmp)
			loaded[curr] = tmp
		}

		resp, _ := http.Get(resourceFolder + "font.json")
		defer resp.Body.Close()
		
		dat, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(dat, &font)

		global.WebLoad = loaded
	}

	rand.Seed(time.Now().UTC().UnixNano())

	cliw.SetCursor(cliw.Ve2{ 0 ,0 })
	if !global.WASM {
		cliw.DrawWorld(cliw.ParsePixMap(cliw.LoadPixMap(resourceFolder + "splash-screen.pix")))
	} else {
		cliw.DrawWorld(cliw.ParsePixMap(global.WebLoad["bg.pix"]))
	}
	time.Sleep(time.Second)
	cliw.SetCursor(cliw.Ve2{ 0 ,0 })
	if !global.WASM {
		cliw.DrawWorld(cliw.ParsePixMap(Menu()))
		fmt.Scanln()
	}
	
	game()
}

func game() {
	// loading stuff
	var pix [][]string
	if !global.WASM {
		pix = cliw.LoadPixMap(resourceFolder + "bg.pix")
		font = cliw.LoadFont(resourceFolder + "font.json")
	} else {
		pix = global.WebLoad["bg.pix"]
	}

	// global stuff
	global = Global{}
	global.BulletSpeed = 80
	global.EnemySpeed = 160
	global.BulletLifetime = 24
	global.MaxEnemies = 3
	global.Enemies = map[string]*clengine.Object{}
	global.PlayerHP = 100
	global.InvincibilityLenght = 10

	// player
	p := clengine.Object{Init: PlayerInit}

	// clengine stuff
	scene := clengine.Scene{}
	scene.Renderer = clengine.Renderer{}
	ClengineInit(&scene)
	defer ClengineClose()
	scene.MinDelta = time.Duration(44 * time.Millisecond)
	scene.BlankPix = pix
	scene.CollMask = clengine.AutoColliderColor(pix, []string{"43523d"}, "env")
	scene.Init([]clengine.Object{p, clengine.Object{Init: SpawnerInit}}, pix)
	scene.RenderFunc = RenderFunc
	scene.Cycle()	
}
