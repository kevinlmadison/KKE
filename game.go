package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running      = true
	bkgColor     = rl.NewColor(147, 211, 196, 255)
	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D
	playerSrc    rl.Rectangle
	playerDst    rl.Rectangle
	playerSpeed  float32 = 3
	musicPaused  bool
	music        rl.Music
	cam          rl.Camera2D
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDst, rl.NewVector2(playerDst.Width, playerDst.Height), 0.0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyF) || rl.IsKeyDown(rl.KeyUp) {
		playerDst.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDst.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyR) || rl.IsKeyDown(rl.KeyLeft) {
		playerDst.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyT) || rl.IsKeyDown(rl.KeyRight) {
		playerDst.X += playerSpeed
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func update() {
	running = !rl.WindowShouldClose()
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
	cam.Target = rl.NewVector2(float32(playerDst.X-(playerDst.Width/2)), float32(playerDst.Y-(playerDst.Height/2)))
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)
	drawScene()
	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "The Orchestrator")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	grassSprite = rl.LoadTexture("assets/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("assets/Characters/BasicCharakterSpritesheet.png")
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDst = rl.NewRectangle(200, 200, 100, 100)
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("assets/Sounds/UraniBorg.wav")
	musicPaused = false
	rl.PlayMusicStream(music)
	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDst.X-(playerDst.Width/2)), float32(playerDst.Y-(playerDst.Height/2))), 0.0, 1.0)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	defer rl.CloseWindow()
}

func main() {
	for running {
		input()
		update()
		render()
	}
	quit()
}
