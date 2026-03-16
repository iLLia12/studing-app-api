package runner

import (
	"fmt"
)

type Payload struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

func Run(data Payload) []byte {

	switch data.Lang {
	case "go":
		fmt.Println("Go lang")
		return []byte("Go lang")
	case "python":
		fmt.Println("Python lang")
		return []byte("Python lang")
	}
  return []byte("Undefined lang")
}

