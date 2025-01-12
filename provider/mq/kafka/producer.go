package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/hdget/hdsdk/types"
)

type Producer struct {
	Logger types.LogProvider
	Option *PublishOption
	Client *ProducerClient
}

var _ types.MqProducer = (*Producer)(nil)

// CreateProducer 创造一个生产者
func (k *Kafka) CreateProducer(name string, args ...map[types.MqOptionType]types.MqOptioner) (types.MqProducer, error) {
	options := k.GetDefaultOptions()
	if len(args) > 0 {
		options = args[0]
	}

	// 初始化rabbitmq client
	client, err := k.newProducerClient(name, options)
	if err != nil {
		return nil, err
	}

	// 客户端连接
	err = client.connect(k.Config.Brokers)
	if err != nil {
		return nil, err
	}

	p := &Producer{
		Logger: k.Logger,
		Client: client,
	}

	return p, nil
}

func (p *Producer) GetLastConfirmedId() uint64 {
	return 0
}

func (p *Producer) Close() {
	p.Client.close()
}

func (p Producer) Publish(data []byte, args ...interface{}) error {
	msgs := make([]*sarama.ProducerMessage, 0)
	for _, topic := range p.Client.Config.Topics {
		m := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(data),
		}
		msgs = append(msgs, m)
	}

	// We are not setting a message key, which means that all messages will
	// be distributed randomly over the different partitions.
	return p.Client.handler.SendMessages(msgs)
}
