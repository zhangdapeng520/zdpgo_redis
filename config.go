package zdpgo_redis

// Config redis配置对象
type Config struct {
	Host      string `yaml:"host" json:"host"`             // 主机地址
	Port      int    `yaml:"port" json:"port"`             // 端口号
	Database  int    `yaml:"database" json:"database"`     // 数据库
	Username  string `yaml:"username" json:"username"`     // 用户名
	Password  string `yaml:"password" json:"password"`     // 密码
	PoolSize  int    `yaml:"pool_size" json:"pool_size"`   // 连接池连接数
	StreamTag string `yaml:"stream_tag" json:"stream_tag"` // 流标签
}

// 获取默认的配置信息
func getDefaultConfig(config Config) Config {
	// 初始化配置
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 6379
	}
	if config.StreamTag == "" {
		config.StreamTag = "zdpgo_redis_config_stream_tag"
	}
	// 初始化Redis连接
	if config.PoolSize == 0 {
		config.PoolSize = 33 // 默认是33个
	}
	return config
}
