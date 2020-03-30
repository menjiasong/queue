package queue


//Job 工作队列
type JobReceivers interface {
	Execute(interface{}) error //执行任务
}
