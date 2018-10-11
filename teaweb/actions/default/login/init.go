package login

import "github.com/iwind/TeaGo"

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Prefix("/login").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
