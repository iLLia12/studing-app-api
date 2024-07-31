package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(code string) string {
	// Create a temporary file
	tmpFile, tempFileErr := os.CreateTemp("", "temp-*.go")
	if tempFileErr != nil {
		fmt.Println("Error creating temporary file:", tempFileErr)
		return "Error creating temporary file"
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println("Error removing temporary file: ", err)
		}
	}(tmpFile.Name()) // Clean up the temporary file

	// Write Go code to the temporary file
	if _, tmpFileWriteErr := tmpFile.Write([]byte(code)); tmpFileWriteErr != nil {
		fmt.Println("Error writing to temporary file:", tmpFileWriteErr)
		return "Error writing to temporary file"
	}
	if tmpFileCloseErr := tmpFile.Close(); tmpFileCloseErr != nil {
		fmt.Println("Error closing temporary file:", tmpFileCloseErr)
		return "Error closing temporary file"
	}

	// Run the Go code using `go run`
	cmd := exec.Command("go", "run", tmpFile.Name())
	output, сombinedOutputErr := cmd.CombinedOutput()
	if сombinedOutputErr != nil {
		fmt.Println("Error running Go code:", сombinedOutputErr)
	}

	return string(output)
}
