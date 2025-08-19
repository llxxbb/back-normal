package config

import (
	"strings"

	"go.uber.org/zap"
)

type KafkaConfig struct {
	Brokers string `yaml:"brokers"`  // Kafka 服务器地址，多个用逗号分隔
	Topic   string `yaml:"topic"`    // 默认主题
	GroupID string `yaml:"group_id"` // 消费者组ID
}

func (c *KafkaConfig) AppendFieldMap(fm map[string]string) {
	// 添加 Kafka 配置字段映射
	fm["kafka.brokers"] = "Kafka.Brokers"
	fm["kafka.topic"] = "Kafka.Topic"
	fm["kafka.group_id"] = "Kafka.GroupID"
}

func (c *KafkaConfig) Print() {
	zap.L().Info("Kafka config",
		zap.String("brokers", c.Brokers),
		zap.String("topic", c.Topic),
		zap.String("group_id", c.GroupID),
	)
}

// GetBrokers 获取 broker 列表
func (c *KafkaConfig) GetBrokers() []string {
	return strings.Split(c.Brokers, ",")
}
