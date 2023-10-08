package etcd

import (
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewClient() (*clientv3.Client, error) {
	cfg := clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD_CLIENT_URL")},
		DialTimeout: 3 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client, nil
}
