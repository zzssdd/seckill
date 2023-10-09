package mq

import (
	"github.com/streadway/amqp"
	. "seckill/pkg/log"
	"strconv"
)

type OrderStatus struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
	q    amqp.Queue
}

func (o *OrderStatus) SetUp(addr string) error {
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
	err = o.Ch.ExchangeDeclare(
		"seckill-delay",
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-delayed-type": "direct",
		},
	)
	if err != nil {
		Log.Errorln("declare exchange err:", err)
		return err
	}
	o.q, err = o.Ch.QueueDeclare(
		"msg_delay",
		true,
		false,
		true,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "seckill-delay",
			"x-dead-letter-routing-key": "delay-key",
		},
	)
	if err != nil {
		Log.Errorln("declare delay queue err:", err)
		return err
	}
	err = o.Ch.QueueBind(
		o.q.Name,
		"delay-key",
		"seckill-delay",
		false,
		nil,
	)
	return err
}

func (o *OrderStatus) DelayPublish(id int64) error {
	str_id := strconv.FormatInt(id, 64)
	return o.Ch.Publish(
		"seckill-delay",
		"delay-key",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(str_id),
			Headers: map[string]interface{}{
				"x-delay": "300000",
			},
		})
}

func (o *OrderStatus) DelayConsume() (<-chan amqp.Delivery, error) {
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
		Log.Errorln("delay consume err:", err)
		return nil, err
	}
	return msgChan, nil
}
