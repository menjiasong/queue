基于rabbitmq的 消息队列golang封装 

一、安装使用 

	go get github.com/menjiasong00/queue

或者 

	git clone github.com/menjiasong00/queue.git


二、主题订阅 topic 

1、执行demo

代码例子：

cd /你的目录/github.com/menjiasong00/queue/test_topic_push  

	go run main.go 

cd /你的目录/github.com/menjiasong00/queue/test_topic_listen  

	go run main.go 


2、代码解释

我们看主题 接收者的接口


	type TopicReceivers interface {

		GetQueueName() string
		
		Execute(routingKey string, data interface{}) error
		
	}


我们假设现在业务完成了 一条 邮件排版的提交。  推送[邮件已写完]的消息：

	que.New().TopicPush("emain.write.finish","i finish an email")

在消息Push进队列后 ， 接收者1号[邮件发送] 接收者2号[干点其他的事] 以 某规则订阅了[邮件已写完] 消息，则这些接收者都能收到该消息： 

只需要把接收者 接口实现出来：

接收者1号[邮件发送]：

	type MsgTopic struct {}

	func (c MsgTopic) GetQueueName() string {

		return "topic_email"
		
	}

	// 执行
	func (c MsgTopic) Execute(routingKey string,data interface{}) error {

		fmt.Println(routingKey)
		
		fmt.Println(data)
		
		return nil
		
	}

接收者2号[干点别的]：

	type TodoTopic struct {}

	func (c TodoTopic) GetQueueName() string {

		return "topic_todo"
		
	}

	// 执行
	func (c TodoTopic) Execute(routingKey string,data interface{}) error {

		fmt.Println(routingKey)
		
		fmt.Println(data)
		
		return nil
		
	}

并在运行的进程运行它们(可参考 queue/test_topic_push 和queue/test_topic_listen ) 

路由绑定规则：
	
	que.New().TopicQueueBind("topic_test",[]string{"xx.*","xx.22.xx"})
	
监听：

	 que.New().TopicListen(MsgTopic{})  

	 que.New().TopicListen(TodoTopic{}) 

两个进程都收到了消息并执行了对应的业务 Execute 。完成了 [邮件排版的提交] 和 [邮件发送]模块 [干点别的]模块 的解耦


三、工作模式 

1、测试demo

新起一个控制台

生产者:

	que.New().Push("TestJob","xxxxxx")

代码例子：

cd /你的目录/github.com/menjiasong00/queue/test_job_push  

go run main.go 

消费者：

	que.New().Listen(map[string]que.JobReceivers{"TestJob":MsgJob{}})

代码例子：

cd /你的目录/github.com/menjiasong00/queue/test_job_listen  

go run main.go 

2、代码解释

我们看工作的接口

	//Job 工作队列
	type JobReceivers interface {

		Execute(interface{}) error //执行任务
		
	}


推送发邮件的消息：

	que.New().Push("SendEmail","this is an email")

在消息Push进队列后 ，监听程序Listen到消息，解析出map里Job名称 ，并调用对应的Execute

因此，只需要把工作的 接口实现出来：

	type SendEmailJob struct {}

	func (c SendEmailJob) Execute(data interface{}) error {

		// 业务代码
		
		fmt.Println(data)
		
		return nil
		
	}

并在运行的进程Listen监听他 
 
	que.New().Listen(map[string]que.JobReceivers{"SendEmail":SendEmailJob{}}) 

把示例的生产者和消费者(可参考 queue/test_job_push 和queue/test_job_listen )go run main.go 。可以看到消费者中执行了SendEmailJob的 Execute


 

四、总结

不管是工作模式还是订阅模式 ，设计思路都是 留出接口 让业务代码进行水平扩展，这样在业务中只需要去实现一个个 job或者topic的接收者。而不用关心消息的流转过程和处理。





