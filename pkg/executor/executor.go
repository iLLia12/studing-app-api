package executor

import (
	"os/exec"
)

type Payload struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

func Run(data Payload) []byte {
	return execCommand("docker", "run", "--rm", "hello-world")
	//c := `docker run --rm 94351554/go-runtime go run main.go -code 'package main; import "fmt"; func main(){ fmt.Println("Hello world") }' -lang go`
	//return execCommand("docker", "run", "--rm", "94351554/go-runtime", "go", "run", "main.go", "-code", c, "-lang", data.Lang)
}

func execCommand(name string, arg ...string) []byte {
	cmd := exec.Command(name, arg...)
	output, _ := cmd.Output()
	return output
}
