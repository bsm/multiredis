package multiredis

import "gopkg.in/redis.v5"

// Client is an abstract client interface which can be either a
// cluster or a sentinel-backed or a single-node client
type Client interface {
	Cmdable
	Close() error
	Pipeline() Pipeline
	PoolStats() *redis.PoolStats
	Sync()
	Watch(fn func(*redis.Tx) error, keys ...string) error
}

// Pipeline is a client-neutral pipeline
type Pipeline interface {
	Cmdable
	Close() error
	Discard() error
	Exec() ([]redis.Cmder, error)
	ReadOnly() *redis.StatusCmd
	ReadWrite() *redis.StatusCmd
	Select(index int) *redis.StatusCmd
	Sync()
}

// New returns a new client, depending on the following conditions:
//
// If a MasterName is passed a sentinel-backed client will be created.
// If the number of Addrs is tow or more, a cluster client will be launched.
// Otherwise, a single-node client will be returned.
func New(opts *Options) Client {
	if opts.MasterName != "" {
		return simpleClient{Client: redis.NewFailoverClient(opts.failover())}
	} else if len(opts.Addrs) > 1 {
		return clusterClient{ClusterClient: redis.NewClusterClient(opts.cluster())}
	}
	return simpleClient{Client: redis.NewClient(opts.simple())}
}

// Cluster always creates a cluster instance, ignoring hints
// from the passed opts.
//
// This is useful when you want to explicitely connect to a redis
// cluster but only have a single seed address to connect to.
func Cluster(opts *Options) Client {
	return clusterClient{ClusterClient: redis.NewClusterClient(opts.cluster())}
}

// --------------------------------------------------------------------

type simpleClient struct{ *redis.Client }

func (c simpleClient) Pipeline() Pipeline { return c.Client.Pipeline() }

type clusterClient struct{ *redis.ClusterClient }

func (c clusterClient) Pipeline() Pipeline { return c.ClusterClient.Pipeline() }
