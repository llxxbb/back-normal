package kafka

import (
	"back/demo/config"
	"back/demo/tool"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

var (
	kafkaProducer *tool.KafkaProducer[Message]
	kafkaConsumer *tool.KafkaConsumer[Message]
	once          sync.Once
)

// InitKafka 初始化 Kafka 服务
func InitKafka(cfg *config.ProjectConfig) error {
	var err error
	once.Do(func() {
		// 初始化生产者
		kafkaProducer, err = tool.NewKafkaProducer[Message](&cfg.Kafka)
		if err != nil {
			zap.L().Error("Failed to initialize Kafka producer", zap.Error(err))
			return
		}

		// 初始化消费者
		kafkaConsumer, err = tool.NewKafkaConsumer[Message](&cfg.Kafka)
		if err != nil {
			zap.L().Error("Failed to initialize Kafka consumer", zap.Error(err))
			return
		}

		// 启动消费者
		err = kafkaConsumer.StartConsumer(nil, &consumerGroupHandler[Message]{
			messageHandler: func(msg Message) error {
				zap.L().Info("Processing message",
					zap.String("id", msg.ID),
					zap.String("content", msg.Content),
					zap.Time("time", msg.Time),
				)
				return nil
			},
		})
		if err != nil {
			zap.L().Error("Failed to start Kafka consumer", zap.Error(err))
			return
		}

		zap.L().Info("Kafka service initialized successfully")
	})

	return err
}

// GetKafkaProducer 获取 Kafka 生产者实例
func GetKafkaProducer() *tool.KafkaProducer[Message] {
	return kafkaProducer
}

// GetKafkaConsumer 获取 Kafka 消费者实例
func GetKafkaConsumer() *tool.KafkaConsumer[Message] {
	return kafkaConsumer
}

// CloseKafka 关闭 Kafka 服务
func CloseKafka() error {
	var errs []error

	if kafkaProducer != nil {
		if err := kafkaProducer.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if kafkaConsumer != nil {
		if err := kafkaConsumer.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("close errors: %v", errs)
	}
	return nil
}
