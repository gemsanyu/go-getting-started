package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/heroku/go-getting-started/container"
	"github.com/heroku/go-getting-started/role"
	"github.com/mmuflih/go-httplib/httplib"
	"go.uber.org/dig"
)

var _ = dig.Name

func main() {
	myrole := make(map[string][]string)

	myrole[role.SA] = []string{role.SA}
	myrole[role.ADMIN] = []string{role.SA, role.ADMIN}
	myrole[role.OWNER] = []string{role.SA, role.ADMIN, role.OWNER}
	myrole[role.CASHIER] = []string{role.SA, role.ADMIN, role.OWNER, role.CASHIER}
	myrole[role.USER] = []string{role.SA, role.ADMIN, role.OWNER, role.CASHIER, role.USER}

	httplib.InitJWTMiddlewareWithRole([]byte("Go-DI-arch"), jwt.SigningMethodHS512, myrole)

	c := container.BuildContainer()

	if err := c.Invoke(container.InvokeRoute); err != nil {
		panic(err)
	}

	if err := c.Provide(container.NewRoute); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(s *container.ServerRoute) {
		s.Run()
	}); err != nil {
		fmt.Println(err)
	}
}
