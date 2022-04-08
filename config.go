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
