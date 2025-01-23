// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const banUser = `-- name: BanUser :exec
UPDATE users
SET is_banned = TRUE
WHERE email = $1
`

func (q *Queries) BanUser(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, banUser, email)
	return err
}

const completeProfile = `-- name: CompleteProfile :exec
UPDATE users
SET
    first_name = $2,
    last_name = $3,
    phone_no = $4,
    gender = $5,
    reg_no = $6,
    vit_email = $7,
    hostel_block = $8,
    room_no = $9,
    github_profile = $10,
    is_profile_complete = TRUE
WHERE email = $1
`

type CompleteProfileParams struct {
	Email         string
	FirstName     string
	LastName      string
	PhoneNo       pgtype.Text
	Gender        string
	RegNo         *string
	VitEmail      *string
	HostelBlock   string
	RoomNo        int32
	GithubProfile string
}

func (q *Queries) CompleteProfile(ctx context.Context, arg CompleteProfileParams) error {
	_, err := q.db.Exec(ctx, completeProfile,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNo,
		arg.Gender,
		arg.RegNo,
		arg.VitEmail,
		arg.HostelBlock,
		arg.RoomNo,
		arg.GithubProfile,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
    id,
    team_id,
    first_name,
    last_name,
    email,
    phone_no,
    gender,
    reg_no,
    vit_email,
    hostel_block,
    room_no,
    github_profile,
    password,
    role,
    is_leader,
    is_verified,
    is_banned,
    is_profile_complete
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
)
`

type CreateUserParams struct {
	ID                uuid.UUID
	TeamID            uuid.NullUUID
	FirstName         string
	LastName          string
	Email             string
	PhoneNo           pgtype.Text
	Gender            string
	RegNo             *string
	VitEmail          *string
	HostelBlock       string
	RoomNo            int32
	GithubProfile     string
	Password          string
	Role              string
	IsLeader          bool
	IsVerified        bool
	IsBanned          bool
	IsProfileComplete bool
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.ID,
		arg.TeamID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNo,
		arg.Gender,
		arg.RegNo,
		arg.VitEmail,
		arg.HostelBlock,
		arg.RoomNo,
		arg.GithubProfile,
		arg.Password,
		arg.Role,
		arg.IsLeader,
		arg.IsVerified,
		arg.IsBanned,
		arg.IsProfileComplete,
	)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users
WHERE id > $1
ORDER BY id ASC
LIMIT $2
`

type GetAllUsersParams struct {
	ID    uuid.UUID
	Limit int32
}

func (q *Queries) GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNo,
			&i.Gender,
			&i.RegNo,
			&i.VitEmail,
			&i.HostelBlock,
			&i.RoomNo,
			&i.GithubProfile,
			&i.Password,
			&i.Role,
			&i.IsLeader,
			&i.IsVerified,
			&i.IsBanned,
			&i.IsProfileComplete,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllVitians = `-- name: GetAllVitians :many
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users WHERE is_vitian = TRUE
`

func (q *Queries) GetAllVitians(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllVitians)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNo,
			&i.Gender,
			&i.RegNo,
			&i.VitEmail,
			&i.HostelBlock,
			&i.RoomNo,
			&i.GithubProfile,
			&i.Password,
			&i.Role,
			&i.IsLeader,
			&i.IsVerified,
			&i.IsBanned,
			&i.IsProfileComplete,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamLeader = `-- name: GetTeamLeader :one
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users WHERE team_id = $1 AND is_leader = TRUE
`

func (q *Queries) GetTeamLeader(ctx context.Context, teamID uuid.NullUUID) (User, error) {
	row := q.db.QueryRow(ctx, getTeamLeader, teamID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Gender,
		&i.RegNo,
		&i.VitEmail,
		&i.HostelBlock,
		&i.RoomNo,
		&i.GithubProfile,
		&i.Password,
		&i.Role,
		&i.IsLeader,
		&i.IsVerified,
		&i.IsBanned,
		&i.IsProfileComplete,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Gender,
		&i.RegNo,
		&i.VitEmail,
		&i.HostelBlock,
		&i.RoomNo,
		&i.GithubProfile,
		&i.Password,
		&i.Role,
		&i.IsLeader,
		&i.IsVerified,
		&i.IsBanned,
		&i.IsProfileComplete,
	)
	return i, err
}

const getUserAndTeamDetails = `-- name: GetUserAndTeamDetails :many

SELECT teams.name, teams.number_of_people, teams.round_qualified, teams.code,
	users.id, users.first_name, users.last_name, users.email, users.reg_no, users.phone_no, users.gender, users.vit_email, users.hostel_block, users.room_no, users.github_profile, users.is_leader
	FROM teams
	INNER JOIN users ON users.team_id = teams.id
	LEFT JOIN submission ON submission.team_id = teams.id
	LEFT JOIN ideas ON ideas.team_id = teams.id
WHERE teams.id = $1
`

type GetUserAndTeamDetailsRow struct {
	Name           string
	NumberOfPeople int32
	RoundQualified pgtype.Int4
	Code           string
	ID             uuid.UUID
	FirstName      string
	LastName       string
	Email          string
	RegNo          *string
	PhoneNo        pgtype.Text
	Gender         string
	VitEmail       *string
	HostelBlock    string
	RoomNo         int32
	GithubProfile  string
	IsLeader       bool
}

// Goofy ahh query hai, but kaam karega if decoded
// SELECT
//
//	(json_build_object(
//	  'user', json_strip_nulls(json_build_object(
//	    'first_name', u.first_name,
//	    'last_name', u.last_name,
//	    'email', u.email,
//	    'phone_no', u.phone_no,
//	    'gender', u.gender,
//	    'reg_no', u.reg_no,
//	    'vit_email', u.vit_email,
//	    'hostel_block', u.hostel_block,
//	    'room_no', u.room_no,
//	    'github_profile', u.github_profile,
//	    'role', u.role
//	  )),
//	  'team', json_build_object(
//	    'team_name', t.name,
//	    'number_of_people', t.number_of_people,
//	    'round_qualified', t.round_qualified,
//	    'code', t.code,
//	    'members', (
//	      SELECT json_agg(json_strip_nulls(json_build_object(
//	        'first_name', members.first_name,
//	        'last_name', members.last_name,
//	        'email', members.email,
//	        'phone_no', members.phone_no,
//	        'github_profile', members.github_profile,
//	        'role', members.role,
//	        'is_leader', members.is_leader
//	      )))
//	      FROM users members
//	      WHERE members.team_id = t.id AND members.id != u.id
//	    )
//	  )
//	))::json AS result
//
// FROM
//
//	users u
//
// JOIN
//
//	teams t ON u.team_id = t.id
//
// WHERE
//
//	u.id = $1;
func (q *Queries) GetUserAndTeamDetails(ctx context.Context, id uuid.UUID) ([]GetUserAndTeamDetailsRow, error) {
	rows, err := q.db.Query(ctx, getUserAndTeamDetails, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserAndTeamDetailsRow
	for rows.Next() {
		var i GetUserAndTeamDetailsRow
		if err := rows.Scan(
			&i.Name,
			&i.NumberOfPeople,
			&i.RoundQualified,
			&i.Code,
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.RegNo,
			&i.PhoneNo,
			&i.Gender,
			&i.VitEmail,
			&i.HostelBlock,
			&i.RoomNo,
			&i.GithubProfile,
			&i.IsLeader,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Gender,
		&i.RegNo,
		&i.VitEmail,
		&i.HostelBlock,
		&i.RoomNo,
		&i.GithubProfile,
		&i.Password,
		&i.Role,
		&i.IsLeader,
		&i.IsVerified,
		&i.IsBanned,
		&i.IsProfileComplete,
	)
	return i, err
}

const getUserByPhoneNo = `-- name: GetUserByPhoneNo :one
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users WHERE phone_no = $1
`

func (q *Queries) GetUserByPhoneNo(ctx context.Context, phoneNo pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByPhoneNo, phoneNo)
	var i User
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Gender,
		&i.RegNo,
		&i.VitEmail,
		&i.HostelBlock,
		&i.RoomNo,
		&i.GithubProfile,
		&i.Password,
		&i.Role,
		&i.IsLeader,
		&i.IsVerified,
		&i.IsBanned,
		&i.IsProfileComplete,
	)
	return i, err
}

const getUserByRegNo = `-- name: GetUserByRegNo :one
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users WHERE reg_no = $1
`

func (q *Queries) GetUserByRegNo(ctx context.Context, regNo *string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByRegNo, regNo)
	var i User
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Gender,
		&i.RegNo,
		&i.VitEmail,
		&i.HostelBlock,
		&i.RoomNo,
		&i.GithubProfile,
		&i.Password,
		&i.Role,
		&i.IsLeader,
		&i.IsVerified,
		&i.IsBanned,
		&i.IsProfileComplete,
	)
	return i, err
}

const getUserByVitEmail = `-- name: GetUserByVitEmail :one
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users WHERE vit_email = $1
`

func (q *Queries) GetUserByVitEmail(ctx context.Context, vitEmail *string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByVitEmail, vitEmail)
	var i User
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNo,
		&i.Gender,
		&i.RegNo,
		&i.VitEmail,
		&i.HostelBlock,
		&i.RoomNo,
		&i.GithubProfile,
		&i.Password,
		&i.Role,
		&i.IsLeader,
		&i.IsVerified,
		&i.IsBanned,
		&i.IsProfileComplete,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, vit_email, hostel_block, room_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNo,
			&i.Gender,
			&i.RegNo,
			&i.VitEmail,
			&i.HostelBlock,
			&i.RoomNo,
			&i.GithubProfile,
			&i.Password,
			&i.Role,
			&i.IsLeader,
			&i.IsVerified,
			&i.IsBanned,
			&i.IsProfileComplete,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unbanUser = `-- name: UnbanUser :exec
UPDATE users
SET is_banned = FALSE
WHERE email = $1
`

func (q *Queries) UnbanUser(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, unbanUser, email)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users
SET password = $2
WHERE email = $1
`

type UpdatePasswordParams struct {
	Email    string
	Password string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.Exec(ctx, updatePassword, arg.Email, arg.Password)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET first_name = $2,
    last_name = $3,
    email = $4,
    phone_no = $5,
    gender = $6,
    reg_no = $7,
    vit_email = $8,
    hostel_block = $9,
    room_no = $10,
    github_profile = $11
WHERE id = $1
`

type UpdateUserParams struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	Email         string
	PhoneNo       pgtype.Text
	Gender        string
	RegNo         *string
	VitEmail      *string
	HostelBlock   string
	RoomNo        int32
	GithubProfile string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNo,
		arg.Gender,
		arg.RegNo,
		arg.VitEmail,
		arg.HostelBlock,
		arg.RoomNo,
		arg.GithubProfile,
	)
	return err
}

const verifyUser = `-- name: VerifyUser :exec
UPDATE users
SET is_verified = TRUE
WHERE email = $1
`

func (q *Queries) VerifyUser(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, verifyUser, email)
	return err
}
