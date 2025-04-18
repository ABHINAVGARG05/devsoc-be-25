package models

import "github.com/jackc/pgx/v5/pgtype"

type UserData struct {
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Email         string      `json:"email"`
	RegNo         string      `json:"reg_no"`
	PhoneNo       pgtype.Text `json:"phone_no"`
	Gender        string      `json:"gender"`
	GithubProfile string      `json:"github_profile"`
	IsLeader      bool        `json:"is_leader"`
	HostelBlock   string      `json:"hostel_block"`
	RoomNo        string      `json:"room_no"`
}

type TeamMember struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	GithubProfile string `json:"github_profile"`
	IsLeader      bool   `json:"is_leader"`
}

type TeamData struct {
	Name           string       `json:"team_name"`
	NumberOfPeople int          `json:"number_of_people"`
	RoundQualified int          `json:"round_qualified"`
	Code           string       `json:"code"`
	Members        []TeamMember `json:"members"`
}

type ResponseData struct {
	User UserData `json:"user"`
	Team TeamData `json:"team"`
}

type UpdateUserRequest struct {
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	PhoneNo       string `json:"phone_no" validate:"required,len=10"`
	Gender        string `json:"gender" validate:"required,len=1"`
	RegNo         string `json:"reg_no" validate:"required"`
	HostelBlock   string `json:"hostel_block" validate:"required"`
	RoomNumber    string `json:"room_no" validate:"required"`
	GithubProfile string `json:"github_profile" validate:"required,url"`
}
