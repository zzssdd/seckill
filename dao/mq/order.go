package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"seckill/dao/mq/model"
	. "seckill/pkg/log"
)

type Order struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
	q    amqp.Queue
}

func (o *Order) SetUp(addr string) error {
	var err error
	o.Conn, err = amqp.Dial(addr)
	if err != nil {
		Log.Errorln("set up rabbitmq err:", err)
		return err
	}
	o.Ch, err = o.Conn.Channel()
	if err != nil {
		Log.Errorln("get rabbitmq channel err:", err)
		return err
	}
	o.q, err = o.Ch.QueueDeclare(
		"seckill-order",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		Log.Errorln("declare queue err:", err)
		return err
	}
	return nil
}

func (o *Order) Publish(id int64, uid int64, pid int, timeStamp int64) error {
	orderInfo := model.Order{
		Id:  id,
		Uid: uid,
		Pid: pid,
	}
	byte_data, err := json.Marshal(orderInfo)
	if err != nil {
		Log.Errorln("marshal order info err:", err)
		return err
	}
	msg := amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: 1,
		Body:         byte_data,
	}
	err = o.Ch.Publish(
		"",
		o.q.Name,
		false,
		false,
		msg,
	)
	if err != nil {
		Log.Errorln("publish msg err:", err)
		return err
	}
	return nil
}

func (o *Order) Consume() (<-chan amqp.Delivery, error) {
	msgChan, err := o.Ch.Consume(
		o.q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		Log.Errorln("consume msg err:", err)
		return nil, err
	}
	return msgChan, nil
}

func StartOrderTask() {

}
