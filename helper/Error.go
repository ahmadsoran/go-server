package helper

import "strings"

func IsSqlDuplicateError(err error) bool {
	if err != nil {
		return strings.Contains(strings.ToLower(err.Error()), "duplicate")
	}
	return false
}
