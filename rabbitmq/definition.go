package rabbitmq

import "github.com/streadway/amqp"

type QueueDeclareSetting struct {
	QueueName  string     // 隊列名稱
	Durable    bool       // 是否持久化
	AutoDelete bool       // 是否自動刪除
	Exclusive  bool       // 是否排他
	NoWait     bool       // 是否不等待
	Args       amqp.Table // 額外的設定
}

type PublishSetting struct {
	Exchange      string // 交換器名稱（默認空字符串表示默認交換器）
	QueueName     string // 隊列名稱
	Mandatory     bool   // 是否強制發送
	Immediate     bool   // 是否等待確認
	MsgPublishing amqp.Publishing
}

type ConsumeSetting struct {
	QueueName string     // 隊列名稱
	Consumer  string     // 消費者名稱
	AutoAck   bool       // 是否自動確認
	Exclusive bool       // 是否排他
	NoLocal   bool       // 是否持久
	NoWait    bool       // 是否不等待
	Args      amqp.Table // 額外的設定
}
