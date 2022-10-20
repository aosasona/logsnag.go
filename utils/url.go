package utils

import "strings"

func makeEndpoint(baseUrl string, path string) string {
	return strings.Trim(baseUrl, "/") + "/" + strings.Trim(path, "/")
}
