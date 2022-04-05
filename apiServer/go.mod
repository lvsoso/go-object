module apiServer

go 1.17

require (
	go-object/lib v1.0.0
)

require github.com/streadway/amqp v1.0.0 // indirect

replace (
	go-object/lib v1.0.0 => ../lib
)
