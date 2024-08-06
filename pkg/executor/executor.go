package executor

import (
	"fmt"
	"os/exec"
)

type CodeExecuteData struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

func Run(data CodeExecuteData) string {
	return execCommand("docker", "run", "--rm", "94351554/go-runtime:v0.3", "go", "run", "main.go", "-code", data.Code, "-lang", data.Lang)
}

func execCommand(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	output, сombinedOutputErr := cmd.CombinedOutput()
	if сombinedOutputErr != nil {
		fmt.Println("Error running Go code:", сombinedOutputErr)
	}
	fmt.Println(string(output))
	return string(output)
}
