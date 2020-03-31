package queue

//topic接口
type TopicReceivers interface {
	GetQueueName() string
	//GetRoutingKeys() []string
	Execute(routingKey string, data interface{}) error
}
