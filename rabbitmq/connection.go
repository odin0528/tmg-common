package rabbitmq

import (
	"fmt"
	"mgmt/common/configs"

	"github.com/streadway/amqp"
)

func Init() (*amqp.Connection, *amqp.Channel, error) {
	return Connection()
}

func Connection() (*amqp.Connection, *amqp.Channel, error) {
	var url = fmt.Sprintf("amqps://%s:%s@%s:%s/",
		configs.Get(configs.SECTION_RABBIT_MQ, configs.MQ_USERNAME, "guest"),
		configs.Get(configs.SECTION_RABBIT_MQ, configs.MQ_PASSWORD, "guest"),
		configs.Get(configs.SECTION_RABBIT_MQ, configs.MQ_HOST, "localhost"),
		configs.Get(configs.SECTION_RABBIT_MQ, configs.MQ_PORT, "5672"))

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}
