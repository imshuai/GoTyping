package utils

import (
	"time"
)

//TimeNow return a formated time.Now() string
func TimeNow() string {
	return time.Now().Format(TimeLayout)
}

//FormatTime return a formated time.Time string
func FormatTime(t time.Time) string {
	return t.Format(TimeLayout)
}
