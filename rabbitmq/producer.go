package rabbitmq

import (
	"github.com/streadway/amqp"
)

// SendMessage 發送消息到 RabbitMQ
func SendMessageSample(ch *amqp.Channel, queueName string, ContentType string, msgBytes []byte) error {

	// 宣告一個隊列
	_, err := ch.QueueDeclare(
		queueName, // 隊列名稱
		true,      // 是否持久化
		false,     // 是否自動刪除
		false,     // 是否排他
		false,     // 是否不等待
		nil,       // 額外的設定
	)
	if err != nil {
		return err
	}

	// 發送消息到隊列
	err = ch.Publish(
		"",        // 交換器名稱（默認空字符串表示默認交換器）
		queueName, // 隊列名稱
		false,     // 是否強制發送
		false,     // 是否等待確認
		amqp.Publishing{
			ContentType: ContentType,
			Body:        msgBytes,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SendMessage(ch *amqp.Channel, declareSetting QueueDeclareSetting, publishSetting PublishSetting) error {

	// 宣告一個隊列
	_, err := ch.QueueDeclare(
		declareSetting.QueueName,  // 隊列名稱
		declareSetting.Durable,    // 是否持久化
		declareSetting.AutoDelete, // 是否自動刪除
		declareSetting.Exclusive,  // 是否排他
		declareSetting.NoWait,     // 是否不等待
		declareSetting.Args,       // 額外的設定
	)
	if err != nil {
		return err
	}

	// 發送消息到隊列
	err = ch.Publish(
		publishSetting.Exchange,  // 交換器名稱（默認空字符串表示默認交換器）
		publishSetting.QueueName, // 隊列名稱
		publishSetting.Mandatory, // 是否強制發送
		publishSetting.Immediate, // 是否等待確認
		publishSetting.MsgPublishing,
	)
	if err != nil {
		return err
	}

	return nil
}
