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

### 环境变量

可以通过环境变量覆盖配置：

```bash
export KAFKA_BROKERS=localhost:9092,localhost:9093
export KAFKA_TOPIC=my-topic
export KAFKA_GROUP_ID=my-group
```

## API 接口

### 1. 发送消息

**接口地址：** `POST /kafka/send`

**请求参数：**
```json
{
  "topic": "test-topic",
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
    "topic": "test-topic",
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
curl -X POST http://localhost:19000/kafka/send \
  -H "Content-Type: application/json" \
  -d '{
    "topic": "test-topic",
    "content": "Hello from curl!"
  }'

# 查看状态
curl http://localhost:19000/kafka/status
```

### 使用 Python 测试

```python
import requests
import json

# 发送消息
url = "http://localhost:19000/kafka/send"
data = {
    "topic": "test-topic",
    "content": "Hello from Python!"
}
response = requests.post(url, json=data)
print(response.json())

# 查看状态
status_url = "http://localhost:19000/kafka/status"
status_response = requests.get(status_url)
print(status_response.json())
```

## 代码结构

```
src/
├── config/
│   └── kafkaConfig.go          # Kafka 配置
├── internal/service/kafka/
│   ├── kafka.go               # Kafka 服务核心逻辑
│   ├── init.go                # 服务初始化
│   └── kafka_test.go          # 单元测试
├── api/
│   ├── kafkaApi.go            # Web API 接口
│   └── kafkaApi_test.go       # API 测试
└── cmd/
    ├── main.go                # 主程序（已集成 Kafka 初始化）
    └── config_default.yml     # 配置文件
```

## 功能特性

1. **自动初始化**：应用启动时自动初始化 Kafka 生产者和消费者
2. **消息发送**：支持同步发送消息到指定主题
3. **消息接收**：自动消费消息并记录日志
4. **状态监控**：提供消息发送和接收的统计信息
5. **错误处理**：完善的错误处理和重试机制
6. **配置灵活**：支持多 broker 和自定义配置

## 注意事项

1. **Kafka 服务要求**：需要本地或远程 Kafka 服务运行
2. **主题创建**：确保目标主题已存在，否则发送消息会失败
3. **网络连接**：确保应用能够连接到 Kafka 服务器
4. **权限设置**：确保有足够的权限进行消息发送和接收

## 故障排除

### 常见问题

1. **连接失败**
   - 检查 Kafka 服务是否运行
   - 验证 broker 地址和端口
   - 检查网络连接

2. **消息发送失败**
   - 确认主题存在
   - 检查权限设置
   - 查看应用日志

3. **消费者不工作**
   - 检查消费者组配置
   - 确认主题有消息
   - 查看消费者日志

### 日志查看

应用会记录详细的 Kafka 操作日志，包括：
- 连接状态
- 消息发送结果
- 消息接收处理
- 错误信息

查看日志文件：`/web/logs/back-normal.log`
