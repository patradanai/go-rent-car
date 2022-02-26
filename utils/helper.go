package utils

import (
	"regexp"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuidNew := uuid.New()

	reg := regexp.MustCompile("[-]")

	return reg.ReplaceAllString(uuidNew.String(), "")
}

func Contains(items []string, item string) bool {
	for _, raw := range items {
		if raw == item {
			return true
		}
	}
	return false
}
