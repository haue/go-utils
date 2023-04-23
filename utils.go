package utils

import "os"

func GetenvDefault(name string, defaultValue string) string {
	s := os.Getenv(name)
	if s == "" {
		s = defaultValue
	}
	return s
}
