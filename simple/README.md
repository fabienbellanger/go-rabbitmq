# go-rabbitmq-simple

RabbitMQ tests - Simple ([Part 1](https://www.rabbitmq.com/tutorials/tutorial-one-go.html))

1.  `docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management`
2.  Producer
    -  `cd send`
    -  `go run .`
3.  Consumer
    -  `cd receive`
    -  `go run .`
