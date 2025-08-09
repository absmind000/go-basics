package main

import (
	"fmt"
	"go-basics/00-module-package/calculator"
	"os"

	"github.com/joho/godotenv"
)

// このプログラムは、環境変数の読み込みと計算機能の使用例を示します
func main() {
	// .envファイルから環境変数を読み込みます
	godotenv.Load()

	// GO_ENV環境変数の値を表示します
	// 出力例: development
	fmt.Println(os.Getenv("GO_ENV"))

	// calculatorパッケージで定義された定数Offsetを表示します
	// 出力例: 100 (calculator/sum.goで定義された値)
	fmt.Println(calculator.Offset)

	// calculatorパッケージのSum関数を使用して1 + 2の計算を行います
	// 出力例: 3
	fmt.Println(calculator.Sum(1, 2))

	// calculatorパッケージのMultiply関数を使用して1 × 2の計算を行います
	// 出力例: 2
	fmt.Println(calculator.Multiply(1, 2))
}
