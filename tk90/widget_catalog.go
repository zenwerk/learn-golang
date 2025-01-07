package main

import (
	_ "embed"
	"fmt"
	. "modernc.org/tk9.0"
)

//go:embed gotk.png
var icon []byte

func NotebookExample() *TNotebookWidget {
	nb := TNotebook()
	// NoteBook のタブ
	lbl1 := Label(Txt("Hello, World!"))
	nb.Add(lbl1, Txt("Tab 1"))
	lbl2 := Label(Txt("こんにちは、世界！"))
	nb.Add(lbl2, Txt("Tab 2"))

	return nb
}

func FrameExample() *TFrameWidget {
	frame := TFrame(
		// デザイン
		Borderwidth(2),
		Relief("sunken"),
	)
	fbtn1 := frame.Button(Txt("フレームボタン1"))
	fbtn2 := frame.Button(Txt("フレームボタン2"))
	flbl1 := frame.Label(Txt("フレームラベル1"))
	flbl2 := frame.Label(Txt("フレームラベル2"))
	Grid(flbl1, Column(0), Row(0))
	Grid(flbl2, Column(0), Row(1))
	Grid(fbtn1, Column(1), Row(0))
	Grid(fbtn2, Column(1), Row(1))

	return frame
}

func LabelFrameExample() *TLabelframeWidget {
	lf := TLabelframe(
		Txt("ラベルフレーム"),
		Width("100"),
		Height("100"),
		Underline(0),
	)
	l1 := lf.Label(
		Txt("フレーム内のラベル"),
	)
	Grid(l1, Column(0), Row(0))

	return lf
}

func TreeviewExample() *TTreeviewWidget {
	c := Columns(`Name Age`)
	tv := TTreeview(
		c,
	)
	tv.Column("#0", Width(0), Stretch("no")) // 非表示にする
	tv.Column("Name", Anchor("w"), Width(100))
	tv.Column("Age", Anchor("e"), Width(50))

	tv.Heading("Name", Txt("名前"), Command(func() {
		MessageBox(Icon("info"), Msg("名前をクリックしました"), Title("情報"))
	}))
	tv.Heading("Age", Txt("年齢"), Command(func() {
		// TODO: 年齢によって昇順・降順に変えたい
	}))

	tv.Insert("", "end", Values(`Alice 20`))
	tv.Insert("", "end", Values(`Bob 30`))
	tv.Insert("", "end", Values(`Seto 40`))

	return tv
}

func main() {
	ErrorMode = CollectErrors
	// ウィンドウタイトルの設定
	App.WmTitle("ウィンドウタイルの設定")
	// アプリアイコンの設定
	imgData := NewPhoto(Data(icon))
	App.IconPhoto(imgData)

	// ボタン
	btn := Button(Txt("Exit"), Command(func() {
		Destroy(App)
	}))
	Grid(btn, Column(0), Row(0))

	// ボタン2
	btn2 := TButton(Txt("TButton からの MsgBox"), Command(func() {
		MessageBox(Icon("error"), Msg("はろ〜"), Title("エラーです"))
	}))
	Grid(btn2, Column(1), Row(0))

	// NoteBook
	nb := NotebookExample()
	Grid(nb, Column(2), Row(0), Sticky("news"))

	// 画像
	imgLbl := Label(Image(imgData))
	Grid(imgLbl, Column(0), Row(1))

	// フレーム
	// フレームをTopLevelに配置
	frame := FrameExample()
	Grid(frame, Column(1), Row(1), Sticky("nsew"))

	// 関数が呼ばれた時点で色選択ダイアログが出る - Button などに Command で設定するのが一般的
	//color := Background(ChooseColor(Initialcolor("gray"), Title("色を選べよ")))
	//MessageBox(Msg(fmt.Sprintf("選ばれたのは %s でした", color)), Title("色だよ〜"))

	// TreeView
	tv := TreeviewExample()
	Grid(tv, Column(2), Row(1), Ipady("2m"), Padx("1m"), Pady("1m"))

	// TLabelframe
	lf := LabelFrameExample()
	Grid(lf, Column(0), Row(2))

	if Error != nil {
		fmt.Printf("Error: %s\n", Error)
	} else {
		App.Wait()
	}

}
