package cassandratest

import (
	"fmt"
	"testing"
	"time"

	"github.com/gocql/gocql"
)

func RecreateKeyspace(t *testing.T, hosts []string, protoVer int, keyspace string) {
	clusterConfig := gocql.NewCluster(hosts...)
	clusterConfig.Timeout = time.Second * 5
	clusterConfig.ProtoVersion = protoVer
	clusterConfig.Consistency = gocql.LocalQuorum

	session, err := clusterConfig.CreateSession()
	for i := 1; i < 12 && err != nil; i++ {
		t.Logf("Can't reach cassandra. Try again after 30s")
		time.Sleep(time.Second * 30)
		session, err = clusterConfig.CreateSession()
	}
	if err != nil {
		t.Fatalf("clusterConfig.CreateSession(): %s", err)
	}
	defer session.Close()

	if err := session.Query(fmt.Sprintf("DROP KEYSPACE IF EXISTS %s;", keyspace)).Exec(); err != nil {
		t.Fatalf("session.Query().Exec(): %s", err)
	}

	if err := session.Query(fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1};", keyspace)).Exec(); err != nil {
		t.Fatalf("session.Query().Exec() %s", err)
	}
}
