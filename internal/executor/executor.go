package executor

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func executeGoScriptAsString(code string) string {
	fmt.Println("Executing executeGoScriptAsString...")
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

func Run(code string) string {
	createContainerAndRunCode()
	return executeGoScriptAsString(code)
}

func createContainerAndRunCode() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	reader, err := cli.ImagePull(ctx, "94351554/go-runtime:v0.1", image.PullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	// cli.ImagePull is asynchronous.
	// The reader needs to be read completely for the pull operation to complete.
	// If stdout is not required, consider using io.Discard instead of os.Stdout.
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "94351554/go-runtime:v0.2",
		Cmd:   []string{"./main"},
		Tty:   false,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	// Remove container after it stops
	removeOptions := container.RemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	defer func() {
		err = cli.ContainerRemove(ctx, resp.ID, removeOptions)
		if err != nil {
			fmt.Println("Error removing container: ", err)
		}
	}()

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	// === Execute cli command ./images/go/main test ===
	// Create a command
	// cmd := exec.Command("./images/go/main", "process", "")
	// Run the command and capture the output
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// Print output
	// fmt.Println(string(output))
}
