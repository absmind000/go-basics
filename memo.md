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


// スライス　ポインタ、長さ、容量の情報を持つ構造体

numbers := []int{1, 2, 3, 4, 5}
┌─────────────────┐
│ ポインタ ────────┼──> [1, 2, 3, 4, 5]  （実際のデータ）
│ 長さ: 5         │
│ 容量: 5         │
└─────────────────┘

slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1  // 同じデータを参照

slice1 ─────┐
            │
            ▼
            [1, 2, 3, 4, 5]
            ▲
            │
slice2 ─────┘


