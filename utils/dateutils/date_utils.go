package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func getNow() time.Time {
	return time.Now().UTC()
}

//GetNowString gets string time of now:apiDateLayout
func GetNowString() string {
	return getNow().Format(apiDateLayout)
}

//GetNowDBFormat gets string time of now :apiDbLayout
func GetNowDBFormat() string {
	return getNow().Format(apiDbLayout)
}
