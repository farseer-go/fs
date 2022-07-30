package exec

import (
	"context"
	"log"
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
			log.Println(output)
		}
		//for {
		//	select {
		//	case output := <-receiveOutput:
		//		log.Println(output)
		//	case <-ctx.Done():
		//		return
		//	}
		//}
	}()

	exitCode := RunShellContext("go env", receiveOutput, env, "", ctx)
	log.Println(exitCode)
}
