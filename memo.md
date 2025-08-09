// ポインタ データの住所を扱う仕組み
var number = 42          // 実際のデータ
var pointer = &number    // データの住所

pointer ──────> number
(住所録)         (42)


// ダブルポインタ　データの住所を順番に参照する仕組み
var number = 42           // 実際のデータ
var pointer = &number     // データの住所
var doublePointer = &pointer  // 住所の住所

doublePointer ──────> pointer ──────> number
(住所録の住所)        (住所録)         (42)



// スライス　ポインタ、長さ(length)、容量(capacity)の情報を持つ構造体

numbers := []int{1, 2, 3, 4, 5}
┌─────────────────┐
│ ポインタ ────────┼──> [1, 2, 3, 4, 5]  （実際のデータ）
│ 長さ: 5         │
│ 容量: 5         │
└─────────────────┘



// 1.スライスの共有
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1  // 同じデータを参照

slice1 ─────┐
            │
            ▼
            [1, 2, 3, 4, 5]
            ▲
            │
slice2 ─────┘



// 2.スライスの一部を参照
// 例1
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1[1:3]  // [2, 3] を参照

slice1 ────> [1, 2, 3, 4, 5]
                 ▲  ▲
slice2  ─────────┘  │
   (長さ: 2)  ──────┘

// 例2
numbers := []int{1, 2, 3, 4, 5}

slice1 := numbers[0:3]  // [1, 2, 3]
slice2 := numbers[1:4]  // [2, 3, 4]
slice3 := numbers[2:5]  // [3, 4, 5]
slice4 := numbers[1:3]  // [2, 3]

// 省略記法も使える
slice5 := numbers[:3]   // [1, 2, 3]  // 開始位置の0を省略
slice6 := numbers[2:]   // [3, 4, 5]  // 終了位置の5を省略
slice7 := numbers[:]    // [1, 2, 3, 4, 5]  // 全部

//   長さと容量の概念
numbers := []int{1, 2, 3, 4, 5}
slice := numbers[1:3]  // [2, 3]

fmt.Println(len(slice))  // 長さ: 2 実際に見えている部分
fmt.Println(cap(slice))  // 容量: 4 拡張可能な部分(元の配列の残り: 2,3,4,5の4つ)
// 長さ（len）は現在見えている要素の数
// 容量（cap）は開始位置から元配列の末尾までの要素数
// 容量の範囲内であれば、スライスを後ろに拡張可能
// 開始位置より前には戻れない