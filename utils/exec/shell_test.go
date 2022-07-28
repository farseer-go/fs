package exec

import (
	"context"
	"fmt"
	"testing"
)

func TestRunShell(t *testing.T) {
	receiveOutput := make(chan string, 100)
	ctx, cancel := context.WithCancel(context.Background())
	env := map[string]string{
		"a": "b",
	}
	RunShell("go env", receiveOutput, env, "", ctx, cancel)
	for {
		select {
		case output := <-receiveOutput:
			fmt.Println(output)
		case <-ctx.Done():
			return
		}
	}
}
