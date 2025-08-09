// calculatorパッケージは基本的な計算機能を提供します
package calculator

import "fmt"

// offsetは小文字で始まるため、このパッケージ内でのみ使用可能な非公開変数です
var offset float64 = 1

// Offsetは大文字で始まるため、他のパッケージからも参照可能な公開変数です
var Offset float64 = 1

// Sum は2つの浮動小数点数を受け取り、その和とoffsetを加えた値を返します
// この関数は公開関数（大文字で始まる）なので、他のパッケージから呼び出せます
func Sum(a float64, b float64) float64 {
	// 内部で非公開関数multiplyと公開関数Multiplyを呼び出してその結果を表示します
	fmt.Println("multiply: ", multiply(a, b))
	fmt.Println("Multiply: ", Multiply(a, b))
	// a + b + offsetを計算して返します
	return a + b + offset
}
