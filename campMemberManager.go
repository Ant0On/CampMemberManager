package campembermanager

import "time"

type Participant struct {
	ParticipantID string
	Name          string
	LastName      string
	BirthDate     time.Time
	GroupNumber   string
	RoomNumber    string
}
