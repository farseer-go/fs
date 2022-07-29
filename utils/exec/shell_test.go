package exec

import (
	"context"
	"fmt"
	"testing"
)

func TestRunShell(t *testing.T) {
	receiveOutput := make(chan string, 100)
	ctx, _ := context.WithCancel(context.Background())
	env := map[string]string{
		"a": "b",
	}

	go func() {
		for output := range receiveOutput {
			fmt.Println(output)
		}
		//for {
		//	select {
		//	case output := <-receiveOutput:
		//		fmt.Println(output)
		//	case <-ctx.Done():
		//		return
		//	}
		//}
	}()

	exitCode := RunShell("go env", receiveOutput, env, "", ctx)
	fmt.Println(exitCode)
}
