package zdpgo_redis

// RedisConfig redis配置对象
type RedisConfig struct {
	Host        string // 主机地址
	Port        int    // 端口号
	Database    int    // 数据库
	Username    string // 用户名
	Password    string // 密码
	LogFilePath string // 日志路径
	Debug       bool   // 是否为debug模式
	PoolSize    int    // 连接池连接数
	StreamTag   string // 流标签
}

// PubStreamConfig 发布流配置
type PubStreamConfig struct {
	Subject string                 `json:"subject"` // 发布主题
	MaxLen  int64                  `json:"max_len"`
	ID      string                 `json:"id"`     // 发布id
	Values  map[string]interface{} `json:"values"` // 发布数据
}

// SubStreamConfig 订阅流配置
type SubStreamConfig struct {
	Subject           string         `json:"subject"`             // 主题
	ConsumerGroupName string         `json:"consumer_group_name"` // 消费者组名称
	HandStreamFunc    HandStreamFunc `json:"hand_stream_func"`    // 处理流中Values数据的方法
}
