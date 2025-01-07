package main

import (
	. "modernc.org/tk9.0"
)

func main() {
	// Frameウィジェットの作成
	frame := TFrame(Width("50m"), Height("10m"))

	// フレーム内のButtonウィジェットの作成
	button1 := frame.Button(Txt("Button1"))
	button2 := frame.Button(Txt("Button2"))
	button3 := frame.Button(Txt("Button3"))
	button4 := frame.Button(Txt("Button4"))

	// ButtonウィジェットをGridレイアウトで配置
	Grid(button1, Column(0), Row(0))
	Grid(button2, Column(1), Row(0))
	Grid(button3, Column(0), Row(1))
	Grid(button4, Column(1), Row(1))

	Grid(frame, Column(0), Row(0), Sticky("news"))
	Grid(TExit(), Column(0), Row(1), Sticky("news"))

	// アプリケーションのメインループを開始
	App.Wait()
}
