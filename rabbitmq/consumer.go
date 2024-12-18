package rabbitmq

import (
	"github.com/streadway/amqp"
)

func GetConsumeDelivery(ch *amqp.Channel, consumeSetting ConsumeSetting) (<-chan amqp.Delivery, error) {
	delivery, err := ch.Consume(
		consumeSetting.QueueName, // 隊列名稱
		consumeSetting.Consumer,  // 消費者名稱
		consumeSetting.AutoAck,   // 是否自動確認
		consumeSetting.Exclusive, // 是否排他
		consumeSetting.NoLocal,   // 是否持久
		consumeSetting.NoWait,    // 是否不等待
		consumeSetting.Args,      // 額外的設定
	)

	if err != nil {
		return nil, err
	}

	return delivery, nil
}
