package multiredis

import "gopkg.in/redis.v3"

// Client is an abstract client interface which can be either a
// cluster or a sentinel-backed or a single-node client
type Client interface {
	commands
	Close() error
	Pipeline() Pipeline
}

// Pipeline is a client-neutral pipeline
type Pipeline interface {
	commands
	Close() error
	Discard() error
	Exec() ([]redis.Cmder, error)
}

// New returns a new client, depending on the following conditions:
//
// If the number of Addrs is one, a single-node client will be returned.
// If a MasterName is passed a sentinel-backed client will be created.
// Otherwise, a cluster client will be launched.
func New(opts *Options) Client {
	if len(opts.Addrs) == 1 {
		return simpleClient{Client: redis.NewClient(opts.simple())}
	} else if opts.MasterName != "" {
		return simpleClient{Client: redis.NewFailoverClient(opts.failover())}
	}
	return clusterClient{ClusterClient: redis.NewClusterClient(opts.cluster())}
}

// --------------------------------------------------------------------

type simpleClient struct{ *redis.Client }

func (c simpleClient) Pipeline() Pipeline { return c.Client.Pipeline() }

type clusterClient struct{ *redis.ClusterClient }

func (c clusterClient) Pipeline() Pipeline { return c.ClusterClient.Pipeline() }
