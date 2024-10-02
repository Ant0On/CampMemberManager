package storage

import (
	"fmt"
	"log"

	campmembermanager "github.com/Ant0On/CampMemberManager"
	"github.com/gocql/gocql"
)

const (
	allParticipantsTableName = "all_participants"
	groupsTableName          = "group_diplomas"
	roomTableName            = "room_diplomas"
)

type Options struct {
	Host         []string
	KeySpace     string
	ProtoVersion int
}

type Storage struct {
	session *gocql.Session
}

func NewStorage(opt Options) (*Storage, error) {
	cluster := gocql.NewCluster(opt.Host...)
	cluster.Keyspace = opt.KeySpace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = opt.ProtoVersion

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("cluster.CreateSession(): %v", err)
	}
	defer session.Close()

	return &Storage{session: session}, nil
}

func (s *Storage) Close() {
	s.session.Close()
}

func (s *Storage) CreateAllParticipantsTable() error {
	if err := s.session.Query(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (participant_id uuid, name text, last_name text, birth_date date, group_number text, room_number text, PRIMARY KEY (participant_id))`, allParticipantsTableName)).Exec(); err != nil {
		return fmt.Errorf("s.session.Query().Exec(CREATE %s): %w", allParticipantsTableName, err)
	}
	return nil
}

func (s *Storage) CreateGroupTable() error {
	if err := s.session.Query(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (group_number text, participant_id uuid, name text, last_name text, birth_date date, room_number text, PRIMARY KEY (group_number, participant_id))`, groupsTableName)).Exec(); err != nil {
		return fmt.Errorf("s.session.Query().Exec(CREATE %s): %w", groupsTableName, err)
	}
	return nil
}

func (s *Storage) CreateRoomTable() error {
	if err := s.session.Query(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (room_number text, participant_id uuid, name text, last_name text, birth_date date, group_number text, PRIMARY KEY (room_number, participant_id))`, roomTableName)).Exec(); err != nil {
		return fmt.Errorf("s.session.Query().Exec(CREATE %s): %w", roomTableName, err)
	}
	return nil
}

func (s *Storage) AddParticipant(participant campmembermanager.Participant) error {
	if err := s.session.Query(fmt.Sprintf("INSERT INTO %s (participant_id, name, last_name, birth_date, group_number, room_number) VALUES (?,?,?,?,?,?)", allParticipantsTableName), participant.ParticipantID, participant.Name, participant.LastName, participant.BirthDate, participant.GroupNumber, participant.RoomNumber).Exec(); err != nil {
		return fmt.Errorf("s.session.Query.Exec(INSERT INTO %s, name: %s, last name: %s): %w", allParticipantsTableName, participant.Name, participant.LastName, err)
	}
	if err := s.session.Query(fmt.Sprintf("INSERT INTO %s (participant_id, name, last_name, birth_date, group_number, room_number) VALUES (?,?,?,?,?,?)", groupsTableName), participant.ParticipantID, participant.Name, participant.LastName, participant.BirthDate, participant.GroupNumber, participant.RoomNumber).Exec(); err != nil {
		return fmt.Errorf("s.session.Query.Exec(INSERT INTO %s, name: %s, last name: %s): %w", groupsTableName, participant.Name, participant.LastName, err)
	}
	if err := s.session.Query(fmt.Sprintf("INSERT INTO %s (participant_id, name, last_name, birth_date, group_number, room_number) VALUES (?,?,?,?,?,?)", roomTableName), participant.ParticipantID, participant.Name, participant.LastName, participant.BirthDate, participant.GroupNumber, participant.RoomNumber).Exec(); err != nil {
		return fmt.Errorf("s.session.Query.Exec(INSERT INTO %s, name: %s, last name: %s): %w", roomTableName, participant.Name, participant.LastName, err)
	}
	return nil
}
