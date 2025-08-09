package main

import (
	"fmt"
	"unsafe" // メモリサイズを取得するためのパッケージ
)

// 定数の宣言（値の変更不可）
const secret = "abc"

// カスタム型の定義：Os という新しい型を int 型をベースに作成
type Os int

// 列挙型の定義：定数のグループを作成
const (
	Mac     Os = iota + 1 // iota は 0 から始まる連番を生成。+1 しているので Mac = 1
	Windows               // 自動で Windows = 2 が設定される
	Linux                 // 自動で Linux = 3 が設定される
)

// パッケージレベルでの変数宣言（初期値は型の初期値が設定される）
var (
	i int    // 初期値: 0
	s string // 初期値: ""（空文字）
	b bool   // 初期値: false
)

func main() {
	// 変数宣言の4つの方法（コメントアウトされている3つと使用中の1つ）
	//var i int        // 1. 変数宣言のみ（初期値は0）
	//var i int = 2    // 2. 型を明示して初期化
	//var i = 2        // 3. 型推論で初期化
	i := 1 // 4. 短縮形で宣言と初期化（関数内でのみ使用可能）

	ui := uint16(2) // uint16型（符号なし16ビット整数）に型変換して代入

	fmt.Println(i) // 単純な値の出力
	// 出力例: 1
	fmt.Printf("i: %v %T\n", i, i) // %v は値を出力、%T は型を出力
	// 出力例: "i: 1 int"
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui) // %[1] は1番目の引数、%[2]は2番目の引数を参照
	// 出力例: "i: 1 int ui: 2 uint16"

	// 異なる型の変数宣言と初期化
	f := 1.23456                      // float64型（浮動小数点数）
	s := "hello"                      // string型（文字列）
	b := true                         // bool型（真偽値）
	fmt.Printf("f: %[1]v %[1]T\n", f) // 浮動小数点数の値と型を出力
	// 出力例: "f: 1.23456 float64"
	fmt.Printf("s: %[1]v %[1]T\n", s) // 文字列の値と型を出力
	// 出力例: "s: hello string"
	fmt.Printf("b: %[1]v %[1]T\n", b) // 真偽値の値と型を出力
	// 出力例: "b: true bool"

	// 複数の変数を一度に宣言と初期化
	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title) // 出力例: "pi: 3.14 title: Go"

	// 型変換の例
	x := 10             // int型
	y := 1.23           // float64型
	z := float64(x) + y // int型をfloat64型に変換してから計算
	fmt.Println(z)      // 出力例: 11.23

	// 列挙型の値を出力
	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux) // 出力例: "Mac:1 Windows:2 Linux:3"

	// 変数の値を変更する演算
	i = 2                    // 値の代入
	fmt.Printf("i: %v\n", i) // 出力例: "i: 2"
	i += 1                   // 加算して代入（i = i + 1 と同じ）
	fmt.Printf("i: %v\n", i) // 出力例: "i: 3"
	i *= 2                   // 乗算して代入（i = i * 2 と同じ）
	fmt.Printf("i: %v\n", i) // 出力例: "i: 6"

	// ポインタの基本
	var ui1 uint16                                  // uint16型の変数を宣言
	fmt.Printf("memory address of ui1: %p\n", &ui1) // ui1のメモリアドレスを出力（&は参照演算子）
	// 出力例: "memory address of ui1: 0xc000018390"
	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2) // 出力例: "memory address of ui2: 0xc000018392"

	var p1 *uint16                                           // uint16型のポインタを宣言（初期値はnil）
	fmt.Printf("value of p1: %v\n", p1)                      // 出力例: "value of p1: <nil>"
	p1 = &ui1                                                // ui1のアドレスをp1に代入
	fmt.Printf("value of p1: %v\n", p1)                      // 出力例: "value of p1: 0xc000018390"
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1)) // ポインタのサイズを出力
	// 出力例: "size of p1: 8[bytes]"
	fmt.Printf("memory address of p1: %p\n", &p1)      // 出力例: "memory address of p1: 0xc00000e028"
	fmt.Printf("value of ui1(dereference): %v\n", *p1) // p1が指す値を取得（*は逆参照演算子）
	// 出力例: "value of ui1(dereference): 0"
	*p1 = 1                               // ポインタ経由で値を変更
	fmt.Printf("value of ui1: %v\n", ui1) // 出力例: "value of ui1: 1"

	// ポインタのポインタ
	var pp1 **uint16 = &p1                                     // ポインタのポインタを宣言と初期化
	fmt.Printf("value of pp1: %v\n", pp1)                      // 出力例: "value of pp1: 0xc00000e028"
	fmt.Printf("memory address of pp1: %p\n", &pp1)            // 出力例: "memory address of pp1: 0xc00000e030"
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1)) // 出力例: "size of pp1: 8[bytes]"
	fmt.Printf("value of p1(dereference): %v\n", *pp1)         // 出力例: "value of p1(dereference): 0xc000018390"
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)       // 二重の逆参照
	// 出力例: "value of ui1(dereference): 1"
	**pp1 = 10                            // ポインタのポインタ経由で値を変更
	fmt.Printf("value of ui1: %v\n", ui1) // 出力例: "value of ui1: 10"

	// スコープの例
	ok, result := true, "A"                               // 変数を宣言と初期化
	fmt.Printf("memory address of result: %p\n", &result) // 出力例: "memory address of result: 0xc000010240"
	if ok {
		result = "B"                                          // 同じスコープの変数を変更
		fmt.Printf("memory address of result: %p\n", &result) // 出力例: "memory address of result: 0xc000010240"
		println(result)                                       // 出力例: B
	} else {
		result = "C"
		println(result) // この行は実行されない（okがtrueのため）
	}
	println(result) // 出力例: B
}
