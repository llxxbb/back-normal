# Kafka 功能使用说明

## 概述

本项目已集成 Kafka 消息队列功能，支持消息的发送和接收。Kafka 服务在应用启动时自动初始化，并提供 Web API 接口用于测试和监控。

## 配置

### 配置文件

在 `src/cmd/config_default.yml` 中添加了 Kafka 配置：

```yaml
kafka:
  brokers: localhost:9092          # Kafka 服务器地址，多个用逗号分隔
  topic: test-topic                # 默认主题
  group_id: back-normal-group      # 消费者组ID
  version: 2.8.0                   # Kafka 版本
  max_retries: 3                   # 最大重试次数
```

## API 接口

### 1. 发送消息

**接口地址：** `POST /kafka/send`

**请求参数：**
```json
{
  "topic": "test-lxb",
  "content": "Hello Kafka!"
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message_id": "550e8400-e29b-41d4-a716-446655440000",
    "topic": "test-lxb",
    "status": "sent",
    "time": "2024-01-01T12:00:00Z"
  }
}
```

### 2. 获取状态

**接口地址：** `GET /kafka/status`

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "total_sent": 10,
    "total_received": 8
  }
}
```

## 使用示例

### 使用 curl 发送消息

```bash
# 发送消息到默认主题
curl -X POST http://localhost:8080/kafka/send \
  -H "Content-Type: application/json" \
  -d '{
    "topic": "test-lxb",
    "content": "Hello from curl!"
  }'

# 查看状态
curl http://localhost:8080/kafka/status
```


