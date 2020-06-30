package common

import "fmt"

const (
	Major        = "0.1"
	Minor        = "0.1"
	ConfigPrefix = "/go/micro"
)

func Version(appVer string) string {
	return fmt.Sprintf("%s-%s-%s", Major, Minor, appVer)
}
