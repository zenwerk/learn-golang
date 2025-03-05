package main

import (
	. "modernc.org/tk9.0"
)

func main() {
	nb := TNotebook()

	// フレーム内のButtonウィジェットの作成
	frame := Frame(Background(Black))
	button1 := frame.Button(Txt("Button1"))
	button2 := frame.Button(Txt("Button2"))
	button3 := frame.Button(Txt("Button3"))
	button4 := frame.Button(Txt("Button4"))
	// ButtonウィジェットをGridレイアウトで配置
	Grid(button1, Column(0), Row(0), Sticky("ew"))
	Grid(button2, Column(1), Row(0))
	Grid(button3, Column(0), Row(1))
	Grid(button4, Column(1), Row(1))
	// ノートブックに追加
	nb.Add(frame, Txt("タブ１"))

	// ２つ目のタブ
	frame2 := Frame(Background(Gray))
	text := frame2.Text()
	Grid(text, Column(0), Row(0), Sticky("nsew"))
	nb.Add(frame2, Txt("タブ２"))

	// アプリ全体のレイアウト
	Grid(nb, Column(0), Row(0), Sticky("nsew"))
	Grid(TExit(), Column(0), Row(1))

	// アプリケーションのメインループを開始
	App.Wait()
}
