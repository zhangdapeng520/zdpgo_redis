package zdpgo_redis

// Config redis配置对象
type Config struct {
	Host      string // 主机地址
	Port      int    // 端口号
	Database  int    // 数据库
	Username  string // 用户名
	Password  string // 密码
	PoolSize  int    // 连接池连接数
	StreamTag string // 流标签
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
