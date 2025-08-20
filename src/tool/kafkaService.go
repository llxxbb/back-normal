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

// KafkaProducer 泛型 Kafka 生产者
type KafkaProducer[T any] struct {
	producer sarama.SyncProducer
	config   KafkaConfig
}

// NewKafkaProducer 创建 Kafka 生产者实例
func NewKafkaProducer[T any](cfg *KafkaConfig) (*KafkaProducer[T], error) {
	// 生产者配置
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(cfg.GetBrokers(), producerConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &KafkaProducer[T]{
		producer: producer,
		config:   *cfg,
	}, nil
}

// SendMessage 发送消息
func (p *KafkaProducer[T]) SendMessage(topic string, message T) error {
	if topic == "" {
		topic = p.config.Topic
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(messageBytes),
	}

	partition, offset, err := p.producer.SendMessage(kafkaMessage)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	zap.L().Info("Message sent", zap.String("topic", topic), zap.Int32("partition", partition), zap.Int64("offset", offset))
	return nil
}

// Close 关闭生产者
func (p *KafkaProducer[T]) Close() error {
	return p.producer.Close()
}

// KafkaConsumer 泛型 Kafka 消费者
type KafkaConsumer[T any] struct {
	consumer sarama.ConsumerGroup
	config   KafkaConfig
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewKafkaConsumer 创建 Kafka 消费者实例
func NewKafkaConsumer[T any](cfg *KafkaConfig) (*KafkaConsumer[T], error) {
	ctx, cancel := context.WithCancel(context.Background())

	// 消费者配置
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumerGroup(cfg.GetBrokers(), cfg.GroupID, consumerConfig)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	return &KafkaConsumer[T]{
		consumer: consumer,
		config:   *cfg,
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// StartConsumer 启动消费者
func (c *KafkaConsumer[T]) StartConsumer(topics []string, handler sarama.ConsumerGroupHandler) error {
	if len(topics) == 0 {
		topics = []string{c.config.Topic}
	}

	go func() {
		for {
			err := c.consumer.Consume(c.ctx, topics, handler)
			if err != nil {
				zap.L().Error("Consumer error", zap.Error(err))
			}
			if c.ctx.Err() != nil {
				return
			}
		}
	}()

	zap.L().Info("Consumer started", zap.Strings("topics", topics))
	return nil
}

// Close 关闭消费者
func (c *KafkaConsumer[T]) Close() error {
	c.cancel()
	return c.consumer.Close()
}
