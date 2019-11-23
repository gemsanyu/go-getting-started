package lib

import (
	"net/smtp"

	"github.com/mmuflih/envgo/conf"
)

/**
 * Created by M. Muflih Kholidin
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

type Gmail struct {
	To      string
	Subject string
	Body    string
	Config  conf.Config
}

func (g *Gmail) Send() error {
	from := g.Config.GetString("gmail.username")
	pass := g.Config.GetString("gmail.password")
	name := g.Config.GetString("gmail.name")

	msg := "From: " + from + "\n" +
		"To: " + g.To + "\n" +
		"Subject: " + g.Subject + "\n\n" +
		g.Body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth(name, from, pass, "smtp.gmail.com"),
		from, []string{g.To}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}
