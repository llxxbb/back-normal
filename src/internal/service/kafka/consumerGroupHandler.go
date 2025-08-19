package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// consumerGroupHandler 消费者组处理器
type consumerGroupHandler struct {
	messageHandler func(*Message) error
}

func (h *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			zap.L().Info("Received message",
				zap.String("topic", message.Topic),
				zap.Int32("partition", message.Partition),
				zap.Int64("offset", message.Offset),
			)

			// 解析消息
			var msg Message
			if err := json.Unmarshal(message.Value, &msg); err != nil {
				zap.L().Error("Failed to unmarshal message", zap.Error(err))
				session.MarkMessage(message, "")
				continue
			}

			// 处理消息
			if err := h.messageHandler(&msg); err != nil {
				zap.L().Error("Failed to handle message", zap.Error(err))
			}

			// 增加接收计数
			IncrementReceivedCount()

			// 标记消息已处理
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
