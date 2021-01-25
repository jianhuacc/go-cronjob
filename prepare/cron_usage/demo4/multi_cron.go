package main

import "github.com/gorhill/cronexpr"

type CronJob struct {
	expr *cronexpr.Expression
}

func main() {
	// 需要有1个调度协程，它定时检查所有的Cron任务，谁过期了就执行谁
}
