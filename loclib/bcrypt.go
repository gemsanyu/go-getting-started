package loclib

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/**
 * Created by M. Muflih Kholidin
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

func GeneratePassword(password string) string {
	bpass, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bpass)
}
