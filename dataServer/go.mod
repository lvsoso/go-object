module dataServer

go 1.17

require (
	go-object/lib/rabbitmq v1.0.0
)

require github.com/streadway/amqp v1.0.0 // indirect

replace (
	go-object/lib/rabbitmq v1.0.0 => ../lib/rabbitmq
)
