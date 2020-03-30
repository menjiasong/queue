package main

import (
	"fmt"
	que "github.com/menjiasong00/queue"
)

func main() {

	err :=que.NewConfig([]string{"127.0.0.1","5672","guest","guest"}).Listen(map[string]que.JobReceivers{"TestJob":MsgJob{}})
	fmt.Println(err.Error())

}

type MsgJob struct {}

// 执行发邮件
func (c MsgJob) Execute(data interface{}) error {
	fmt.Println(data)
	return nil
}
