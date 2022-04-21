package util

import (
	"encoding/json"
	"fmt"
)

func DebugPrint(v interface{}) {
	res, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println("==================")
	fmt.Println(string(res))
	fmt.Println("==================")
}
