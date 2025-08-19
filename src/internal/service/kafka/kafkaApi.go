package kafka

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/llxxbb/platform-common/access"
	"github.com/llxxbb/platform-common/def"
	"go.uber.org/zap"
)

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	Topic   string `json:"topic" binding:"required"`   // 主题
	Content string `json:"content" binding:"required"` // 消息内容
}

// SendMessageResponse 发送消息响应
type SendMessageResponse struct {
	MessageID string `json:"message_id"` // 消息ID
	Topic     string `json:"topic"`      // 主题
	Status    string `json:"status"`     // 状态
	Time      string `json:"time"`       // 发送时间
}

// MessageStatusResponse 消息状态响应
type MessageStatusResponse struct {
	TotalSent     int64 `json:"total_sent"`     // 总发送数
	TotalReceived int64 `json:"total_received"` // 总接收数
}

// SendKafkaMessage 发送 Kafka 消息
func SendKafkaMessage(c *gin.Context) {
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, access.GetErrorResultD[SendMessageResponse](def.ET_ENV, def.E_ENV.Code, "Invalid request parameters: "+err.Error(), nil))
		return
	}

	// 获取 Kafka 服务
	kafkaService := GetKafkaService()
	if kafkaService == nil {
		c.JSON(http.StatusOK, access.GetErrorResultD[SendMessageResponse](def.ET_ENV, def.E_ENV.Code, "Kafka service not initialized", nil))
		return
	}

	// 创建消息
	message := &Message{
		ID:      uuid.New().String(),
		Content: req.Content,
		Time:    time.Now(),
	}

	// 发送消息
	if err := kafkaService.SendMessage(req.Topic, message); err != nil {
		zap.L().Error("Failed to send Kafka message", zap.Error(err))
		c.JSON(http.StatusOK, access.GetErrorResultD[SendMessageResponse](def.ET_ENV, def.E_ENV.Code, "Failed to send message: "+err.Error(), nil))
		return
	}

	// 返回响应
	response := SendMessageResponse{
		MessageID: message.ID,
		Topic:     req.Topic,
		Status:    "sent",
		Time:      message.Time.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, access.GetSuccessResult(response))
}

// GetKafkaStatus 获取 Kafka 状态
func GetKafkaStatus(c *gin.Context) {
	sent, received := GetMessageCounts()
	response := MessageStatusResponse{
		TotalSent:     sent,
		TotalReceived: received,
	}

	c.JSON(http.StatusOK, access.GetSuccessResult(response))
}
