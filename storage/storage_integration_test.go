package storage

import (
	"testing"

	cassandratest "github.com/Ant0On/CampMemberManager/pkg"
	"github.com/Ant0On/CampMemberManager/testhelpers"
	"github.com/stretchr/testify/require"
)

func createStorage(t *testing.T, cfg testhelpers.TestOptions) *Storage {
	cassandratest.RecreateKeyspace(t, cfg.Hosts, cfg.ProtoVersion, cfg.KeySpace)

	store, err := NewStorage(Options{
		Host:         cfg.Hosts,
		KeySpace:     cfg.KeySpace,
		Consistency:  cfg.Consistency,
		ProtoVersion: cfg.ProtoVersion,
	})
	require.NoErrorf(t, err, "createStorage: NewStorage fail")
	require.NoErrorf(t, store.CreateAllParticipantsTable(), "createStorage: store.CreateAllParticipantsTable() fail")
	require.NoErrorf(t, store.CreateGroupTable(), "createStorage: store.CreateGroupTable() fail")
	require.NoErrorf(t, store.CreateRoomTable(), "createStorage: store.CreateRoomTable() fail")

	return store
}
