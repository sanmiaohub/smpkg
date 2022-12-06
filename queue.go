package smpkg

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
)

func NewBroker() *Broker {
	return &Broker{
		queueList: make(map[string]*Worker),
	}
}

type Broker struct {
	queueList map[string]*Worker
}

func (b *Broker) Start() {
	for queueName, worker := range b.queueList {
		go worker.Start(queueName)
	}
}

func (b *Broker) RegisterQueue(queueName string, worker *Worker) {
	if _, ok := b.queueList[queueName]; !ok {
		b.queueList[queueName] = worker
	}
}

func (b *Broker) waitSignedStop() {
	waiting := make(chan bool, 1)
	inputSigned := make(chan os.Signal, 1)
	signal.Notify(inputSigned, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		if s := <-inputSigned; len(s.String()) > 0 {
			waiting <- true
			os.Exit(1)
		}
	}()
	<-waiting
}

type Worker struct {
	handle    func(message string)
	mqUrl     string
	queueName string
}

func (w *Worker) Start(name string) {
	xmq := NewXmq(w.mqUrl)
	defer xmq.ReleaseRes()

	q, err := xmq.Channel.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		fmt.Println("声明消息队列异常:" + err.Error())
	}

	msgList, err := xmq.Channel.Consume(q.Name, "", true, false, false, true, nil)
	if err != nil {
		fmt.Println("获取消息通道异常:" + err.Error())
	}
	for d := range msgList {
		w.DoHandle(d)
	}
}

func (w *Worker) DoHandle(qBody amqp.Delivery) {
	w.handle(string(qBody.Body))
}
