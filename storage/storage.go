package storage

import (
	"log"

	"github.com/gocql/gocql"
)

type Options struct {
	Host         string
	KeySpace     string
	ProtoVersion int
}

type Storage struct {
	session gocql.Session
}

func NewStorage(opt Options) (*Storage, error) {
	cluster := gocql.NewCluster(opt.Host)
	cluster.Keyspace = opt.KeySpace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = opt.ProtoVersion

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("cluster.CreateSession(): %w", err)
		return nil, err
	}
	defer session.Close()

	return nil, err
}
