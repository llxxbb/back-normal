package tool

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
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

// KafkaService 泛型 Kafka 服务
type KafkaService[T any] struct {
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup
	config   KafkaConfig
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewKafkaService 创建 Kafka 服务实例
func NewKafkaService[T any](cfg *KafkaConfig) (*KafkaService[T], error) {
	ctx, cancel := context.WithCancel(context.Background())

	// 生产者配置
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(cfg.GetBrokers(), producerConfig)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	// 消费者配置
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumerGroup(cfg.GetBrokers(), cfg.GroupID, consumerConfig)
	if err != nil {
		producer.Close()
		cancel()
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	return &KafkaService[T]{
		producer: producer,
		consumer: consumer,
		config:   *cfg,
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// SendMessage 发送消息
func (k *KafkaService[T]) SendMessage(topic string, message T) error {
	if topic == "" {
		topic = k.config.Topic
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(messageBytes),
	}

	partition, offset, err := k.producer.SendMessage(kafkaMessage)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	zap.L().Info("Message sent", zap.String("topic", topic), zap.Int32("partition", partition), zap.Int64("offset", offset))
	return nil
}

// StartConsumer 启动消费者
func (k *KafkaService[T]) StartConsumer(topics []string, handler sarama.ConsumerGroupHandler) error {
	if len(topics) == 0 {
		topics = []string{k.config.Topic}
	}

	go func() {
		for {
			err := k.consumer.Consume(k.ctx, topics, handler)
			if err != nil {
				zap.L().Error("Consumer error", zap.Error(err))
			}
			if k.ctx.Err() != nil {
				return
			}
		}
	}()

	zap.L().Info("Consumer started", zap.Strings("topics", topics))
	return nil
}

// Close 关闭服务
func (k *KafkaService[T]) Close() error {
	k.cancel()
	var errs []error
	if err := k.producer.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := k.consumer.Close(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("close errors: %v", errs)
	}
	return nil
}
