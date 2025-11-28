package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

const (
	// 1タイルのサイズ（ピクセル単位）
	tileW, tileH = 16, 16
	// マップのサイズ（タイル数単位）
	mapW, mapH = 30, 20
	// 画面サイズ（ピクセル単位）= マップサイズ × タイルサイズ
	screenW, screenH = mapW * tileW, mapH * tileH
)

const mapPath = "maps/map.tmx" // Path to your Tiled Map.

type game struct {
	// mapImage はレンダリングされたマップ画像を保持する
	mapImage *ebiten.Image
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	if g.mapImage != nil {
		screen.DrawImage(g.mapImage, nil)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	// Parse .tmx file.
	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	if len(gameMap.Tilesets) != 1 {
		fmt.Println("expected have exactly one tileset in this case.")
	}

	// You can also render the map to an in-memory image for direct
	// use with the default Renderer, or by making your own.
	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Render just layer 0 to the Renderer.
	if err := renderer.RenderLayer(0); err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}
	if err := renderer.RenderLayer(1); err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}
	if err := renderer.RenderObjectGroup(0); err != nil {
		fmt.Printf("object group unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Get a reference to the Renderer's output, an image.NRGBA struct.
	img := renderer.Result

	// image.NRGBA を ebiten.Image に変換
	ebitenImg := ebiten.NewImageFromImage(img)

	// ゲームを初期化して実行
	g := &game{
		mapImage: ebitenImg,
	}

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Tiled Map Viewer")

	if err := ebiten.RunGame(g); err != nil {
		fmt.Printf("error running game: %s", err.Error())
		os.Exit(2)
	}
}
