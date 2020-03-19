package kafka

import (
	"bytes"
	"io"

	"github.com/Shopify/sarama"
)

var _ io.WriteCloser = &Writer{}

type Writer struct {
	option   *Option
	config   *sarama.Config
	producer sarama.AsyncProducer
}

func NewKafkaWriter(o *Option) (*Writer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_0_0

	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner

	kafkaWriter := &Writer{
		option: o,
		config: config,
	}

	return kafkaWriter, nil
}

func (k *Writer) send(key []byte, p []byte) error {
	msg := &sarama.ProducerMessage{
		Topic:     k.option.Topic,
		Partition: int32(k.option.Partition),
		Key:       sarama.ByteEncoder(key),
		Value:     sarama.ByteEncoder(p),
	}
	select {
	case k.producer.Input() <- msg:
	case <-k.producer.Errors():
	case <-k.producer.Successes():
	}

	return nil
}

func (k *Writer) Start() error {
	producer, err := sarama.NewAsyncProducer(k.option.Hosts, k.config)
	if err != nil {
		return err
	}
	k.producer = producer

	return nil
}

func (k *Writer) Write(p []byte) (int, error) {
	bytesSlice := bytes.Split(p, []byte(" "))
	// discard error command
	if len(bytesSlice) < 2 {
		return 0, nil
	}
	return len(p), k.send(bytesSlice[1], p)
}

func (k *Writer) Close() error {
	return k.producer.Close()
}
