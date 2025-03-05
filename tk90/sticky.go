package main

import (
	. "modernc.org/tk9.0"
)

func main() {
	t := Text(Background(White), Foreground(Black), State("normal"), Insertbackground(Black))

	StyleConfigure("TScrollbar", Foreground(White), Background(White))

	scroll := TScrollbar(Command(func(e *Event) { e.Yview(t) }))
	Grid(t, Column(0), Row(0), Sticky("nesw"))
	Grid(scroll, Column(1), Row(0), Sticky("nes"))
	// 伸長するように指定
	GridColumnConfigure(App, 0, Weight(1))
	GridRowConfigure(App, 0, Weight(1))

	// アプリケーションのメインループを開始
	App.Wait()
}
