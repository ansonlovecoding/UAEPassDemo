# Cryptocurrency Exchange Server

## About Architecture

- Web Framework: GIN or GraphQL
- Authentication: JWT
- Configuration Framework: Viper
- Service Discovery & Configuration: Consul
- Key Management Services ( KMS ): HashiCorp Vault
- Remote Procedure Call: gRPC
- Caching: Redis
- Message Queue: Pulsar
- Relation Database (main database): CockroachDB OR PostgreSQL
  - 存储用户数据和资产数据
- ORM Library: GORM
- Time Series Database (messages and historical orders): InfluxDB
  - 存储定序消息和行情数据和历史订单
- Notifications: SSE
  - 更轻量，单向通信
- Real-time Data Warehouse: Apache Link or Pulsar Golang Client
  - 用于实时统计处理行情数据
- Offline Data Warehouse: Apache Spark
  - 用于处理分析历史数据
- API Document: Swagger
- Log: Logrus + ELK
- Monitor: Prometheus or Grafna
- 



go run github.com/99designs/gqlgen generate --config pkg/common/graph/gqlgen_admin.yml

go run github.com/99designs/gqlgen generate --config pkg/common/graph/gqlgen_web.yml