package kafka

import (
	"cdel/demo/Normal/config"
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type KafkaService struct {
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup
	config   *config.KafkaConfig
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewKafkaService 创建 Kafka 服务实例
func NewKafkaService(cfg *config.KafkaConfig) (*KafkaService, error) {
	ctx, cancel := context.WithCancel(context.Background())

	// 创建生产者配置
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Return.Errors = true

	// 创建生产者
	producer, err := sarama.NewSyncProducer(cfg.GetBrokers(), producerConfig)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	// 创建消费者配置
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	// 创建消费者组
	consumer, err := sarama.NewConsumerGroup(cfg.GetBrokers(), cfg.GroupID, consumerConfig)
	if err != nil {
		producer.Close()
		cancel()
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	return &KafkaService{
		producer: producer,
		consumer: consumer,
		config:   cfg,
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// SendMessage 发送消息到指定主题
func (k *KafkaService) SendMessage(topic string, message *Message) error {
	if topic == "" {
		topic = k.config.Topic
	}

	// 序列化消息
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// 创建 Kafka 消息
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(messageBytes),
	}

	// 发送消息
	partition, offset, err := k.producer.SendMessage(kafkaMessage)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	zap.L().Info("Message sent successfully",
		zap.String("topic", topic),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
		zap.String("message_id", message.ID),
	)

	// 增加发送计数
	IncrementSentCount()

	return nil
}

// StartConsumer 启动消费者
func (k *KafkaService) StartConsumer(topics []string, messageHandler func(*Message) error) error {
	if len(topics) == 0 {
		topics = []string{k.config.Topic}
	}

	handler := &consumerGroupHandler{
		messageHandler: messageHandler,
	}

	go func() {
		for {
			err := k.consumer.Consume(k.ctx, topics, handler)
			if err != nil {
				zap.L().Error("Error from consumer", zap.Error(err))
			}
			if k.ctx.Err() != nil {
				return
			}
		}
	}()

	zap.L().Info("Kafka consumer started", zap.Strings("topics", topics))
	return nil
}

// Close 关闭 Kafka 服务
func (k *KafkaService) Close() error {
	k.cancel()

	var errs []error

	if err := k.producer.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close producer: %w", err))
	}

	if err := k.consumer.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close consumer: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing kafka service: %v", errs)
	}

	return nil
}
