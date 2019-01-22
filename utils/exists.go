package utils

import (
	"os"
)

func FileExists(f string) (bool, error) {
	if _, err := os.Stat(f); !os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}
