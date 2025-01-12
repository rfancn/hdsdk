package types

type Configer interface {
	GetLogConfig() interface{}          // 日志配置
	GetMysqlConfig() interface{}        // mysql数据库配置
	GetRedisConfig() interface{}        // redis缓存配置
	GetRabbitmqConfig() interface{}     // GetRabbitmqConfig 获取RabbitMq队列配置
	GetKafkaConfig() interface{}        // GetKafkaConfig 获取kafka消息队列配置
	GetNosqlConfig() interface{}        // NoSQL服务配置
	GetKvConfig() interface{}           // kv型服务配置
	GetMicroServiceConfig() interface{} // gokit微服务配置
}

// items under sdk config
type SdkConfigItem struct {
	Log          interface{} `mapstructure:"log"`     // 日志配置
	Mysql        interface{} `mapstructure:"mysql"`   // 数据库配置
	Redis        interface{} `mapstructure:"redis"`   // 缓存配置
	RabbitMq     interface{} `mapstructure:"aliyun"`  // rabbitmq消息队列配置
	Kafka        interface{} `mapstructure:"kafka"`   // kafka消息队列配置
	MicroService interface{} `mapstructure:"service"` // 基于gokit实现的微服务配置, ms means microservice
	Nosql        interface{} `mapstructure:"nosql"`   // NoSQL配置
	Kv           interface{} `mapstructure:"kv"`      // Key/Value数据库配置
}
