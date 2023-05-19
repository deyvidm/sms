package utils

import (
	"encoding/json"
	"fmt"
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
