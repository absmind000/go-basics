package calculator

// Multiply は2つの浮動小数点数を受け取り、その積とoffsetを加えた値を返します
// この関数は公開関数（大文字で始まる）なので、他のパッケージから呼び出せます
func Multiply(a float64, b float64) float64 {
	return (a * b) + offset
}

// multiply は2つの浮動小数点数を受け取り、その積とoffsetを加えた値を返します
// この関数は非公開関数（小文字で始まる）なので、このパッケージ内でのみ使用可能です
func multiply(a float64, b float64) float64 {
	return (a * b) + offset
}
