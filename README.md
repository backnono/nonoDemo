# nonoDemo
go的基础框架

## 项目的层级目录
* 主要的项目层级目录
```
  ├── adapter // Adapter层，适配各种框架及协议的接入，比如：Gin，tRPC，Echo，Fiber 等
  ├── application // App层，处理Adapter层适配过后与框架、协议等无关的业务逻辑
  │   ├── consumer //（可选）处理外部消息，比如来自消息队列的事件消费
  │   ├── dto // App层的数据传输对象，外层到达App层的数据，从App层出发到外层的数据都通过DTO传播
  │   ├── executor // 处理请求，包括command和query
  │   └── scheduler //（可选）处理定时任务，比如Cron格式的定时Job
  ├── domain // Domain层，最核心最纯粹的业务实体及其规则的抽象定义
  │   ├── gateway // 领域网关，model的核心逻辑以Interface形式在此定义，交由Infra层去实现
  │   └── model // 领域模型实体
  ├── infrastructure // Infra层，各种外部依赖，组件的衔接，以及domain/gateway的具体实现
  │   ├── cache //（可选）内层所需缓存的实现，可以是Redis，Memcached等
  │   ├── client //（可选）各种中间件client的初始化
  │   ├── config // 配置实现
  │   ├── database //（可选）内层所需持久化的实现，可以是MySQL，MongoDB，Neo4j等
  │   ├── distlock //（可选）内层所需分布式锁的实现，可以基于Redis，ZooKeeper，etcd等
  │   ├── log // 日志实现，在此接入第三方日志库，避免对内层的污染
  │   ├── mq //（可选）内层所需消息队列的实现，可以是Kafka，RabbitMQ，Pulsar等
  │   ├── node //（可选）服务节点一致性协调控制实现，可以基于ZooKeeper，etcd等
  │   └── rpc //（可选）广义上第三方服务的访问实现，可以通过HTTP，gRPC，tRPC等
  └── pkg // 各层可共享的公共组件代码
  * 其中infra层级目录样例
  * ├── infrastructure
    │   ├── cache
    │   │   └── redis.go // Redis 实现的缓存
    │   ├── client
    │   │   ├── kafka.go // 构建 Kafka client
    │   │   ├── mysql.go // 构建 MySQL client
    │   │   ├── redis.go // 构建 Redis client（cache和distlock中都会用到 Redis，统一在此构建）
    │   │   └── zookeeper.go // 构建 ZooKeeper client
    │   ├── config
    │   │   └── config.go // 配置定义及其解析
    │   ├── database
    │   │   ├── dataobject.go // 数据库操作依赖的数据对象
    │   │   └── mysql.go // MySQL 实现的数据持久化
    │   ├── distlock
    │   │   ├── distributed_lock.go // 分布式锁接口，在此是因为domain/gateway中没有直接需要此接口
    │   │   └── redis.go // Redis 实现的分布式锁
    │   ├── mq
    │   │   ├── dataobject.go // 消息队列操作依赖的数据对象
    │   │   └── kafka.go // Kafka 实现的消息队列
    │   ├── node
    │   │   └── zookeeper_client.go // ZooKeeper 实现的一致性协调节点客户端
    │   └── rpc
    │       ├── dataapi.go // 第三方服务访问功能封装
    │       └── dataobject.go // 第三方服务访问操作依赖的数据对象
```