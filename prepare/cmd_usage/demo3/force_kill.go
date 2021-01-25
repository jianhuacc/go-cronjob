package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	// 执行1个cmd,让它在一个协程里去执行，执行2秒：sleep 2;echo hello;
	// 然后在1秒时杀死cmd
	resultChan := make(chan *result, 1000)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	go func() {
		cmd := exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2;echo hello;")

		output, err := cmd.CombinedOutput()
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	// 继续往下走,睡1秒
	time.Sleep(1 * time.Second)
	// 取消上下文，即可杀死
	cancelFunc()

	// 在main协程里等待子协程的退出，并打印任务执行结果
	res := <-resultChan
	fmt.Println(res.err, string(res.output))
}
