package utils

import "strings"

func ParsePath(path string) string {
	return strings.ReplaceAll(path, "/", "\\")
}
