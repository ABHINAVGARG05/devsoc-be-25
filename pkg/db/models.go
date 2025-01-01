// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Idea struct {
	ID          uuid.UUID
	Title       string
	Description string
	Track       string
	TeamID      uuid.UUID
	IsSelected  bool
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Score struct {
	ID             uuid.UUID
	TeamID         uuid.UUID
	Design         int32
	Implementation int32
	Presentation   int32
	Round          int32
}

type Submission struct {
	ID         uuid.UUID
	GithubLink string
	FigmaLink  string
	PptLink    string
	OtherLink  string
	TeamID     uuid.UUID
}

type Team struct {
	ID             uuid.UUID
	Name           string
	NumberOfPeople int32
	RoundQualified pgtype.Int4
	Code           string
}

type User struct {
	ID         uuid.UUID
	Name       string
	TeamID     uuid.NullUUID
	Email      string
	IsVitian   bool
	RegNo      string
	Password   string
	PhoneNo    string
	Role       string
	IsLeader   bool
	College    string
	IsVerified bool
	IsBanned   bool
}
