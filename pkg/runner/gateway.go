package gateway

import (
	"fmt"
	"os/exec"
	"strings"
)

type Payload struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

var images = map[string]string{
		"go": "golang:1.22",
		"pythob": "python:3.12",
		"node": "node:20",
	}

func Run(data Payload) []byte {

	image := ""

	switch data.Lang {
	case "go":
		image = images[data.Lang]
		fmt.Println("Go lang\n" + string(data.Code))
		return []byte("Go lang\n" + string(data.Code))
	case "python":
		image = images[data.Lang]
		fmt.Println("Python lang\n" + string(data.Code))
		return []byte("Python lang\n" + string(data.Code))
	}
  return []byte("Undefined lang\n" + string(data.Code) + image)
}

func PythonRuntime(code string) {
	cmd := exec.Command("docker", "run", "--rm", "-i", "python:3.12", "python")
	cmd.Stdin = strings.NewReader(code)
}

