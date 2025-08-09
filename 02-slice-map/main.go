package main

import "fmt"

func main() {
	// 配列の宣言と初期化の例
	var a1 [3]int               // 要素数3の配列を宣言（初期値は0）
	var a2 = [3]int{10, 20, 30} // 要素数3の配列を値を指定して初期化
	a3 := [...]int{10, 20}      // 要素数を値から推論（この場合は2）
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	// 出力: [0 0 0] [10 20 30] [10 20]

	fmt.Printf("%v %v\n", len(a3), cap(a3))
	// 出力: 2 2（長さと容量が同じ）

	fmt.Printf("%T %T\n", a2, a3)
	// 出力: [3]int [2]int（型が異なる - 要素数が型の一部）

	// スライスの宣言と初期化の例
	var s1 []int  // nilスライスの宣言
	s2 := []int{} // 空のスライスの初期化
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// 出力: s1: []int [] 0 0（型、値、長さ、容量）

	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	// 出力: s2: []int [] 0 0

	fmt.Println(s1 == nil) // nilスライスはnilと等しい
	// 出力: true

	fmt.Println(s2 == nil) // 空スライスはnilではない
	// 出力: false

	// スライスへの要素追加
	s1 = append(s1, 1, 2, 3) // 複数要素を追加
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// 出力: s1: []int [1 2 3] 3 4（長さ3、容量4）

	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...) // スライスの要素を展開して追加
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	// 出力: s1: []int [1 2 3 4 5 6] 6 8（容量が自動的に拡張）

	// makeを使用したスライスの作成
	s4 := make([]int, 0, 2) // 長さ0、容量2のスライス
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	// 出力: s4: []int [] 0 2

	s4 = append(s4, 1, 2, 3, 4) // 容量を超えて追加
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	// 出力: s4: []int [1 2 3 4] 4 4（容量が自動的に拡張）

	// スライスの参照と部分スライス
	s5 := make([]int, 4, 6) // 長さ4、容量6のスライス
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5: [0 0 0 0] 4 6

	s6 := s5[1:3] // インデックス1から2までの部分スライス
	s6[1] = 10    // s6の変更はs5にも影響する（同じ配列を参照）
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5: [0 0 10 0] 4 6

	fmt.Printf("s6: %v %v %v\n", s6, len(s6), cap(s6))
	// 出力: s6: [0 10] 2 5

	s6 = append(s6, 2) // 元の配列の容量内での追加
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5: [0 0 10 2] 4 6（s5も影響を受ける）

	fmt.Printf("s6 appended: %v %v %v\n", s6, len(s6), cap(s6))
	// 出力: s6 appended: [0 10 2] 3 5

	// スライスのコピー
	sc6 := make([]int, len(s5[1:3])) // コピー先のスライスを作成
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5 source of copy: [0 0 10 2] 4 6

	fmt.Printf("sc6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// 出力: sc6 dst copy before: [0 0] 2 2

	copy(sc6, s5[1:3]) // スライスのコピー
	fmt.Printf("sc6 dst of copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// 出力: sc6 dst of copy after: [0 10] 2 2

	sc6[1] = 12 // コピーしたスライスの変更は元に影響しない
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5: [0 0 10 2] 4 6

	fmt.Printf("sc6: %v %v %v\n", sc6, len(sc6), cap(sc6))
	// 出力: sc6: [0 12] 2 2

	// 完全スライス式（容量を制限）
	s5 = make([]int, 4, 6)
	fs6 := s5[1:3:3] // インデックス1から2まで、容量を3に制限
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5: [0 0 0 0] 4 6

	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	// 出力: fs6: [0 0] 2 2

	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8) // 容量制限を超えての追加は新しい配列を作成
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5: [0 6 7 0] 4 6

	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	// 出力: fs6: [6 7 8] 3 4

	s5[3] = 9
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))
	// 出力: s5: [0 6 7 9] 4 6

	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6))
	// 出力: fs6: [6 7 8] 3 4（fs6は独立した配列を参照）

	// マップの宣言と操作
	var m1 map[string]int  // nilマップの宣言
	m2 := map[string]int{} // 空のマップの初期化
	fmt.Printf("%v %v \n", m1, m1 == nil)
	// 出力: map[] true

	fmt.Printf("%v %v \n", m2, m2 == nil)
	// 出力: map[] false

	m2["A"] = 10 // 要素の追加
	m2["B"] = 20
	m2["C"] = 0
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])
	// 出力: map[A:10 B:20 C:0] 3 10

	delete(m2, "A") // 要素の削除
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"])
	// 出力: map[B:20 C:0] 2 0

	v, ok := m2["A"] // 要素の存在確認（存在しない場合）
	fmt.Printf("%v %v\n", v, ok)
	// 出力: 0 false

	v, ok = m2["C"] // 要素の存在確認（値が0の場合）
	fmt.Printf("%v %v\n", v, ok)
	// 出力: 0 true

	// マップの繰り返し処理
	for k, v := range m2 { // キーと値を取得
		fmt.Printf("%v %v\n", k, v)
	}
	// 出力:
	// B 20
	// C 0
}
