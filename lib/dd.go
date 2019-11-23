package lib

import (
	"encoding/json"
	"fmt"
)

/**
 * Created by M. Muflih Kholidin
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

func DD(data interface{}) {
	o, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(o))
}
