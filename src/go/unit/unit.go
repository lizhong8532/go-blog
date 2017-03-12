package unit

import (
	"../config"
)

func CreateApi(str string) string {
	return config.API_PREFIX + str
}
