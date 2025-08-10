package main

import (
	"fmt"
	"unsafe"
)

// Task構造体の定義
// Title: タスクのタイトル（文字列）
// Estimate: タスクの見積もり時間（整数）
type Task struct {
	Title    string
	Estimate int
}

func main() {
	// 値型での構造体の操作
	// task1を初期化
	task1 := Task{
		Title:    "Learn Golang",
		Estimate: 3,
	}
	// task1のTitleを変更
	task1.Title = "Learning Go"
	// task1の型情報(%T)、フィールド名と値(%+v)、Titleの値(%v)を出力
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)
	// 出力結果: main.Task {Title:Learning Go Estimate:3} Learning Go

	// 値型の代入（コピーが作成される）
	var task2 Task = task1
	task2.Title = "new"
	// task2のTitleを変更してもtask1には影響しない
	fmt.Printf("task1: %v task2: %v\n", task1.Title, task2.Title)
	// 出力結果: task1: Learning Go task2: new

	// ポインタ型での構造体の操作
	// task1pをポインタとして初期化
	task1p := &Task{
		Title:    "Learn concurrency",
		Estimate: 2,
	}
	// ポインタの型情報、値、サイズ（バイト）を出力
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	// 出力結果: task1p: *main.Task {Title:Learn concurrency Estimate:2} 8

	// ポインタ経由でTitleを変更
	task1p.Title = "Changed"
	fmt.Printf("task1p: %+v\n", *task1p)
	// 出力結果: task1p: {Title:Changed Estimate:2}

	// ポインタの代入（同じアドレスを参照）
	var task2p *Task = task1p
	task2p.Title = "Changed by Task2"
	// task2pを変更するとtask1pも変更される（同じメモリを参照しているため）
	fmt.Printf("task1: %+v\n", *task1p)
	// 出力結果: task1: {Title:Changed by Task2 Estimate:2}
	fmt.Printf("task2: %+v\n", *task2p)
	// 出力結果: task2: {Title:Changed by Task2 Estimate:2}

	// レシーバの違いによる動作の確認
	task1.extendEstimate()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate) // (&task1).extendEstimatePointer()は省略可能
	// 出力結果: task1 value receiver: 3

	task1.extendEstimatePointer()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)
	// 出力結果: task1 value receiver: 13
}

// メソッドの定義は main() の前後どちらでも可能
// 値レシーバのメソッド
// コピーに対して操作を行うため、元の値は変更されない
func (task Task) extendEstimate() {
	task.Estimate += 10
}

// ポインタレシーバのメソッド
// 実体に対して操作を行うため、元の値が変更される
func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10
}
