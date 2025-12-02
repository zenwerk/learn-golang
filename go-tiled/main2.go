package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

const (
	// 1タイルのサイズ（ピクセル単位）
	tileW, tileH = 16, 16
	// マップのサイズ（タイル数単位）
	mapW, mapH = 30, 20
	// 画面サイズ（ピクセル単位）= マップサイズ × タイルサイズ
	screenW, screenH = mapW * tileW, mapH * tileH
)

// loadImage は埋め込みファイルシステムから画像を読み込み、Ebitengine用の画像として返す。
// assetPath: 埋め込みファイルシステム内の画像パス
// 戻り値: Ebitengineで描画可能な画像オブジェクト
func loadImage(assetPath string) *ebiten.Image {
	// 埋め込みファイルシステムから画像ファイルを開く
	log.Printf("loading image %s", assetPath)
	f, err := os.Open(assetPath)
	if err != nil {
		log.Panic(err)
	}

	// 画像をデコード（PNGデコーダーはブランクインポートで登録済み）
	img, _, err := image.Decode(f)
	if err != nil {
		log.Panic(err)
	}

	// Go標準の画像からEbitengine用の画像に変換して返す
	return ebiten.NewImageFromImage(img)
}

// loadMap は指定された名前のTiledマップファイル（.tmx）を読み込む。
// name: マップ名（拡張子なし）
// 戻り値: パースされたマップデータとエラー
func loadMap(name string) (*tiled.Map, error) {
	// マップファイルのパスを構築（asset/map/[name].tmx）
	path := filepath.Join("maps", name+".tmx")
	log.Printf("loading map %s", path)

	// go-tiledライブラリでマップファイルをパース
	// WithFileSystem オプションで埋め込みファイルシステムを使用
	gameMap, err := tiled.LoadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load map %s: %s", path, err)
	}
	return gameMap, nil
}

// game はEbitengineのゲームインターフェースを実装するメイン構造体。
// Ebitengineでは、この構造体にUpdate、Draw、Layoutメソッドを実装する必要がある。
type game struct {
	// t はタイルセットの定義情報（タイルサイズ、タイル数など）を保持する
	t *tiled.Tileset

	// m はTiledで作成したマップデータ（レイヤー、タイル配置など）を保持する
	m *tiled.Map

	// tileImages はタイルセット画像から切り出した各タイルの画像を配列で保持する
	// インデックスはタイルIDに対応している
	tileImages []*ebiten.Image

	// initialized はゲームの初期化が完了したかどうかを示すフラグ
	// Update内で一度だけ初期化処理を行うために使用
	initialized bool
}

// loadTileset はマップに関連付けられたタイルセットを読み込み、
// 各タイルの画像を切り出して配列に格納する。
func (g *game) loadTileset() {
	// マップに含まれるタイルセットの数を確認(このサンプルでは1つのタイルセットのみをサポート)
	numTilesets := len(g.m.Tilesets)
	if numTilesets != 1 {
		log.Panicf("failed to load map: expected 1 tileset, got %d", numTilesets)
	}

	// タイルセットへの参照を保存
	g.t = g.m.Tilesets[0]

	// タイルセット画像を読み込む
	// g.t.Image.Source にはタイルセット画像のファイル名が格納されている
	path := filepath.Join("maps", g.t.Image.Source)
	log.Printf("loading tileset image %s", path)
	img := loadImage(filepath.Join(path))
	log.Printf("%+v\n", g.t)

	// タイルセット画像から各タイルを切り出してサブイメージとして保存
	// サブイメージは元画像の一部を参照するだけなので、メモリ効率が良い
	// 各サブイメージは同じタイルセット画像を参照している
	g.tileImages = make([]*ebiten.Image, g.t.TileCount)
	for i := 0; i < g.t.TileCount; i++ {
		// このタイルの矩形領域を計算
		// GetTileRect は GlobalID を期待するので FirstGID + i を渡す
		r := g.t.GetTileRect(g.t.FirstGID + uint32(i))

		// タイルのサブイメージ参照を保存
		g.tileImages[i] = img.SubImage(r).(*ebiten.Image)
	}
}

// initialize はゲームの初期状態をセットアップする。
// マップとタイルセットの読み込み、各種フラグの初期化を行う。
// この関数はUpdate内で一度だけ呼ばれる。
func (g *game) initialize() {
	// マップファイルを読み込む（asset/map/map.tmx）
	var err error
	g.m, err = loadMap("map")
	if err != nil {
		log.Panic(err)
	}

	// タイルセットを読み込み、各タイルの画像を準備
	g.loadTileset()

	// ゲームの初期設定
	g.initialized = true // 初期化完了フラグをセット
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func (g *game) Update() error {
	// 初期化処理（初回のみ実行）
	// Update内で初期化を行うのは、Ebitengineの初期化が完了してから
	// アセットを読み込むため
	//if !g.initialized {
	//	g.initialize()
	//}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// 初期化が完了していない場合は描画しない
	if !g.initialized {
		return
	}

	// 描画オプション（位置、回転、スケールなどを設定できる）
	op := &ebiten.DrawImageOptions{}

	// 全レイヤーを順番に描画
	for _, l := range g.m.Layers {
		// レイヤー内の全タイルをループ
		for i, tile := range l.Tiles {
			// 空のタイルはスキップ
			if tile.Nil {
				continue
			}

			// タイルマップ上の位置を計算（タイル単位）
			// i はタイルの1次元インデックス、これを2次元座標に変換
			tileX, tileY := i%mapW, i/mapW

			// 画面上の描画位置を計算（ピクセル単位）
			screenX, screenY := tileX*tileW, tileY*tileH

			// tileImages のインデックスを計算
			// tile.ID は GlobalID なので FirstGID を引いてローカルインデックスにする
			tileIndex := int(tile.ID - g.t.FirstGID)
			if tileIndex < 0 || tileIndex >= len(g.tileImages) {
				continue
			}

			// タイルを描画
			op.GeoM.Reset()                                       // 変換行列をリセット
			op.GeoM.Translate(float64(screenX), float64(screenY)) // 描画位置を設定
			screen.DrawImage(g.tileImages[tileIndex], op)
		}
	}

	// 画面左上にデバッグ情報を表示
	// FPS: 実際のフレームレート（描画頻度）
	// TPS: 実際のティックレート（Update呼び出し頻度）
	debugText := fmt.Sprintf("FPS %.0f\nTPS %.0f", ebiten.ActualFPS(), ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, debugText)
}

func main() {
	ebiten.SetWindowTitle("Creating and Loading Tilemaps Using Ebitengine - Trevors-Tutorials.com #11")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenW, screenH)
	//ebiten.SetFullscreen(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(100)

	g := &game{}
	g.initialize()

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
