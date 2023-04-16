package wmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// InitRabbitMQClient 初始化rabbitmq连接
func InitRabbitMQClient(user, password, host, port string) (mq *amqp.Connection, err error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port))
}
