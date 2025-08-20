package kafka

import (
	"back/demo/config"
	"back/demo/tool"
	"sync"

	"go.uber.org/zap"
)

var (
	kafkaService *tool.KafkaService[Message]
	once         sync.Once
)

// InitKafka 初始化 Kafka 服务
func InitKafka(cfg *config.ProjectConfig) error {
	var err error
	once.Do(func() {
		kafkaService, err = tool.NewKafkaService[Message](&cfg.Kafka)
		if err != nil {
			zap.L().Error("Failed to initialize Kafka service", zap.Error(err))
			return
		}

		// 启动消费者
		err = kafkaService.StartConsumer(nil, &consumerGroupHandler[Message]{
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

// GetKafkaService 获取 Kafka 服务实例
func GetKafkaService() *tool.KafkaService[Message] {
	return kafkaService
}

// CloseKafka 关闭 Kafka 服务
func CloseKafka() error {
	if kafkaService != nil {
		return kafkaService.Close()
	}
	return nil
}
