package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ASSET_SCALING  = float32(0.10)
	ASSET_VELOCITY = float32(100)
	SCREEN_FPS     = 60
	DELTA_TIME     = 1.0 / SCREEN_FPS
	SCREEN_WIDTH   = 450
	SCREEN_HEIGHT  = 450
)

var (
	assetPosition = rl.Vector2Zero()
	assetVelocity = rl.Vector2{
		X: ASSET_VELOCITY * DELTA_TIME,
		Y: ASSET_VELOCITY * DELTA_TIME,
	}
)

func main() {
	rl.InitWindow(450, 450, "Go Probe")

	rl.SetTargetFPS(60)

	asset := rl.LoadTexture("./pepega.png")

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		if assetPosition.Y+float32(asset.Height)*ASSET_SCALING > SCREEN_HEIGHT || assetPosition.Y < 0 {
			assetVelocity.Y = -assetVelocity.Y
		}
		if assetPosition.X+float32(asset.Width)*ASSET_SCALING > SCREEN_WIDTH || assetPosition.X < 0 {
			assetVelocity.X = -assetVelocity.X
		}
		assetPosition = rl.Vector2Add(assetPosition, assetVelocity)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTextureEx(asset, assetPosition, 0.5, ASSET_SCALING, rl.White)
		rl.EndDrawing()
	}

}
