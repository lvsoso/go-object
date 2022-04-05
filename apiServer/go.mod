module apiServer

go 1.17

require go-object/lib v1.0.0

require (
	github.com/klauspost/cpuid/v2 v2.0.6 // indirect
	github.com/klauspost/reedsolomon v1.9.16 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
)

replace go-object/lib v1.0.0 => ../lib
