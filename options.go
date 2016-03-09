package multiredis

import (
	"time"

	"gopkg.in/redis.v3"
)

type Options struct {
	// Addrs is the list of addresses.
	// If only one address is passed, a single-node client will be created,
	// if multiple addresses are passed, then a cluster client will be returned.
	Addrs []string

	// DB only applies to single-node clients
	DB int64

	// MasterName only accepted for sentinel backed clients
	MasterName string

	// Password private token, if any
	Password string

	// PoolSize is the connection pool size
	PoolSize int

	DialTimeout, ReadTimeout, WriteTimeout, PoolTimeout, IdleTimeout time.Duration
}

func (o *Options) cluster() *redis.ClusterOptions {
	if len(o.Addrs) == 0 {
		o.Addrs = []string{"127.0.0.1:6379"}
	}

	return &redis.ClusterOptions{
		Addrs:    o.Addrs,
		PoolSize: o.PoolSize,
		Password: o.Password,

		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
		PoolTimeout:  o.PoolTimeout,
		IdleTimeout:  o.IdleTimeout,
	}
}

func (o *Options) failover() *redis.FailoverOptions {
	if len(o.Addrs) == 0 {
		o.Addrs = []string{"127.0.0.1:26379"}
	}

	return &redis.FailoverOptions{
		SentinelAddrs: o.Addrs,
		MasterName:    o.MasterName,
		PoolSize:      o.PoolSize,
		Password:      o.Password,
		DB:            o.DB,

		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
		PoolTimeout:  o.PoolTimeout,
		IdleTimeout:  o.IdleTimeout,
	}
}

func (o *Options) simple() *redis.Options {
	addr := "127.0.0.1:6379"
	if len(o.Addrs) > 0 {
		addr = o.Addrs[0]
	}

	return &redis.Options{
		Addr:     addr,
		PoolSize: o.PoolSize,
		Password: o.Password,
		DB:       o.DB,

		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
		PoolTimeout:  o.PoolTimeout,
		IdleTimeout:  o.IdleTimeout,
	}
}
