package main

import (
	"time"
)

/**
Functional Options Patternと言い、任意のオプションを引数で渡したいケースがある場合に使用します。
Go言語にはデフォルト引数という機能が存在しないため、既存のコンストラクタに新しく任意の引数を追加したいケースの場合に参照先のソースコードを壊してしまいます。

修正前
func NewHoge(int hoge1) {...}
NewHoge(10);

修正後
func NewHoge(hoge1 int, hoge2 int) {...}

上記の例では修正後にhoge2の引数が加わったことによって修正前のNewHoge(10)が引数が足りずにエラーを起こすとこになり既存のソースコードを壊してしまいます。

Functional Options Patternでは引数を可変長に受け取れるため、例えばWithIsDebugを追加して使用した場合でも①には影響は受けず後方互換性が高く拡張性が高いことがわかるかと思います。
**/

type Server struct {
	Addr    string
	Port    int
	Timeout time.Duration
	IsDebug bool
}

func DefaultServer() *Server {
	return &Server{
		Addr:    "localhost",
		Port:    8080,
		Timeout: 10 * time.Second,
	}
}

type Option func(*Server)

func WithTimeout(d time.Duration) Option {
	return func(c *Server) {
		c.Timeout = d
	}
}

// 今回の仕様追加に伴ってデバッグフラグが追加された
func WithIsDebug(is_debug bool) Option {
	return func(c *Server) {
		c.IsDebug = is_debug
	}
}

func NewServer(opts ...Option) *Server {
	s := DefaultServer()
	for _, f := range opts {
		f(s)
	}

	return s
}

func FunctionalOptionsPattern() {
	NewServer(WithTimeout(1000)) // ①

	NewServer(WithTimeout(1000), WithIsDebug(true)) // ②
}
