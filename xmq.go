package smpkg

import (
	"github.com/streadway/amqp"
	"log"
)

type XMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewXmq(mqUrl string) (xmq XMQ) {
	var err error
	//创建rabbitmq连接
	xmq.Conn, err = amqp.Dial(mqUrl)
	checkErr(err, "创建连接失败")

	//创建Channel
	xmq.Channel, err = xmq.Conn.Channel()
	checkErr(err, "创建channel失败")
	return
}

func (x *XMQ) ReleaseRes() {
	_ = x.Conn.Close()
	_ = x.Channel.Close()
}

func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err.Error())
	}
}
