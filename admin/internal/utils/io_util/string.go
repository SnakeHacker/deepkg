package io_util

import "strings"

func CleanJsonStr(str string) (result string) {
	result = strings.TrimPrefix(str, "```json")
	result = strings.TrimSuffix(result, "```")

	return result
}
