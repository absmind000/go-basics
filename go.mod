// go.mod は依存関係を管理する重要なファイルです。
// 役割：
// - プロジェクトの依存関係（外部パッケージ）を管理
// - モジュール名とGoのバージョンを定義
// - 必要なパッケージとそのバージョンを指定
// - 通常は `go get` コマンドで自動的に更新される
// - npm の package.json に相当

module go-basics

go 1.20

require (
	github.com/joho/godotenv v1.5.1
	go.uber.org/goleak v1.2.1
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2
	golang.org/x/sync v0.1.0
)
