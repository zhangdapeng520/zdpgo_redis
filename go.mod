module github.com/zhangdapeng520/zdpgo_redis

go 1.17

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-redsync/redsync/v4 v4.5.0
	github.com/zhangdapeng520/zdpgo_random v0.1.0
	github.com/zhangdapeng520/zdpgo_zap v0.1.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.20.0 // indirect
)

replace github.com/zhangdapeng520/zdpgo_zap v0.1.0 => ../zdpgo_zap

replace github.com/zhangdapeng520/zdpgo_random v0.1.0 => ../zdpgo_random
