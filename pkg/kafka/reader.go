package kafka

import "github.com/Shopify/sarama"

// 放着先，后面实现

type Reader struct {
	option *Option
	config *sarama.Config
}
