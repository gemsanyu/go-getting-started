package provider

import (
	"github.com/heroku/go-getting-started/httphandler/extra"
	"github.com/heroku/go-getting-started/httphandler/ping"
	"github.com/heroku/go-getting-started/httphandler/profile"
	"github.com/heroku/go-getting-started/httphandler/user"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildHandlerProvider(c *dig.Container) *dig.Container {
	if err := c.Provide(ping.NewPingHandler); err != nil {
		panic(err)
	}

	if err := c.Provide(extra.NewP404Handler); err != nil {
		panic(err)
	}

	if err := c.Provide(user.NewAddHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewEditHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewGetHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewListHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewGetTokenHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewGetMyHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewVoidHandler); err != nil {
		panic(err)
	}

	if err := c.Provide(profile.NewAddHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(profile.NewEditHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(profile.NewGetHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(profile.NewUploadAvatarHandler); err != nil {
		panic(err)
	}

	return c
}
