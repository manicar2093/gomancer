package templutils

import (
	"net/http"
	"slices"
)

func IsOverwritableMethod(method string) bool {
	return slices.Contains(
		[]string{
			http.MethodConnect,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodPost,
			http.MethodPut,
			http.MethodTrace,
		}, method)
}

func IsPathStyles(path, expected, trueStyles, falseStyles string) string {
	if path == expected {
		return trueStyles
	}
	return falseStyles
}
