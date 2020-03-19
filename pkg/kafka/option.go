package kafka

type Options func(*Option)

type Option struct {
	Hosts     []string `json:"host"`
	Partition int      `json:"partition"`
	Topic     string   `json:"topic"`
}

func SetKafkaOptionHosts(host []string) Options {
	return func(o *Option) {
		o.Hosts = host
	}
}

func SetKafkaOptionPartition(partition int) Options {
	return func(o *Option) {
		o.Partition = partition
	}
}

func SetKafkaOptionUser(topic string) Options {
	return func(o *Option) {
		o.Topic = topic
	}
}
