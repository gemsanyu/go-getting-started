package lib

import (
	"time"
)

/**
 * Created by M. Muflih Kholidin
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

var YMDHMS string = "2006-01-02 15:04:05"
var YMD string = "2006-01-02"
var HMS string = "15:04:05"
var TIMEPARSER string = time.RFC3339

func ZEROTIME() time.Time {
	t, _ := time.Parse(YMDHMS, "0001-01-01 00:00:00")
	return t
}
