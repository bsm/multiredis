package multiredis

import "gopkg.in/redis.v5"

// Client is an abstract client interface which can be either a
// cluster or a sentinel-backed or a single-node client
type Client interface {
	redis.Cmdable

	Process(cmd redis.Cmder) error
	Close() error
}

// New returns a new client, depending on the following conditions:
//
// If a MasterName is passed a sentinel-backed client will be created.
// If the number of Addrs is tow or more, a cluster client will be launched.
// Otherwise, a single-node client will be returned.
func New(opts *Options) Client {
	if opts.MasterName != "" {
		return redis.NewFailoverClient(opts.failover())
	} else if len(opts.Addrs) > 1 {
		return redis.NewClusterClient(opts.cluster())
	}
	return redis.NewClient(opts.simple())
}

// Cluster always creates a cluster instance, ignoring hints
// from the passed opts.
//
// This is useful when you want to explicitely connect to a redis
// cluster but only have a single seed address to connect to.
func Cluster(opts *Options) Client {
	return redis.NewClusterClient(opts.cluster())
}

// --------------------------------------------------------------------
