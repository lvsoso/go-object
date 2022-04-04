module apiServer

go 1.17

require (
	go-object/lib/objectstream v1.0.0
	go-object/lib/rabbitmq v1.0.0
)

require github.com/streadway/amqp v1.0.0 // indirect

replace (
	go-object/lib/objectstream v1.0.0 => ../lib/objectstream
	go-object/lib/rabbitmq v1.0.0 => ../lib/rabbitmq
)
