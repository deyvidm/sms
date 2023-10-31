package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unicode"
)

func Ptr[T any](v T) *T {
	return &v
}

func JSONDump(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func IsAllNumeric(term string) bool {
	_, err := strconv.Atoi(term)
	return err == nil
}

func IsAllAlpha(term string) bool {
	for _, ch := range term {
		if !unicode.IsLetter(ch) {
			return false
		}
	}
	return true
}

func ParseFloat64(str string) (float64, error) {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
