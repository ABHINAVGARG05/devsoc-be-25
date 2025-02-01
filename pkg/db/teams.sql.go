// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: teams.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addUserToTeam = `-- name: AddUserToTeam :exec
UPDATE users
SET team_id = $1
WHERE id = $2
`

type AddUserToTeamParams struct {
	TeamID uuid.NullUUID
	ID     uuid.UUID
}

func (q *Queries) AddUserToTeam(ctx context.Context, arg AddUserToTeamParams) error {
	_, err := q.db.Exec(ctx, addUserToTeam, arg.TeamID, arg.ID)
	return err
}

const banTeam = `-- name: BanTeam :exec
UPDATE users
SET is_banned = TRUE
WHERE id = $1
`

func (q *Queries) BanTeam(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, banTeam, id)
	return err
}

const countTeamMembers = `-- name: CountTeamMembers :one
SELECT COUNT(*) FROM users
WHERE team_id = $1
`

func (q *Queries) CountTeamMembers(ctx context.Context, teamID uuid.NullUUID) (int64, error) {
	row := q.db.QueryRow(ctx, countTeamMembers, teamID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams (
    id, name, number_of_people, round_qualified, code, is_banned
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, name, number_of_people, round_qualified, code, is_banned
`

type CreateTeamParams struct {
	ID             uuid.UUID
	Name           string
	NumberOfPeople int32
	RoundQualified pgtype.Int4
	Code           string
	IsBanned       bool
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRow(ctx, createTeam,
		arg.ID,
		arg.Name,
		arg.NumberOfPeople,
		arg.RoundQualified,
		arg.Code,
		arg.IsBanned,
	)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.NumberOfPeople,
		&i.RoundQualified,
		&i.Code,
		&i.IsBanned,
	)
	return i, err
}

const decreaseUserCountTeam = `-- name: DecreaseUserCountTeam :exec
UPDATE teams
SET number_of_people = number_of_people - 1
WHERE id = $1
`

func (q *Queries) DecreaseUserCountTeam(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, decreaseUserCountTeam, id)
	return err
}

const deleteTeam = `-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id = $1
`

func (q *Queries) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteTeam, id)
	return err
}

const findTeam = `-- name: FindTeam :one
SELECT id,name,code,round_qualified FROM teams
WHERE code = $1
LIMIT 1
`

type FindTeamRow struct {
	ID             uuid.UUID
	Name           string
	Code           string
	RoundQualified pgtype.Int4
}

func (q *Queries) FindTeam(ctx context.Context, code string) (FindTeamRow, error) {
	row := q.db.QueryRow(ctx, findTeam, code)
	var i FindTeamRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.RoundQualified,
	)
	return i, err
}

const getTeamById = `-- name: GetTeamById :one
SELECT teams.id, teams.name, teams.round_qualified, teams.code,teams.is_banned,
       score.design, score.implementation, score.presentation, score.round,
       submission.title, submission.description, submission.track, submission.github_link, submission.figma_link, submission.other_link,
       ideas.title, ideas.description, ideas.track, ideas.is_selected
FROM teams
INNER JOIN score ON score.team_id = teams.id
LEFT JOIN submission ON submission.team_id = teams.id
LEFT JOIN ideas ON ideas.team_id = teams.id
WHERE teams.id = $1
`

type GetTeamByIdRow struct {
	ID             uuid.UUID
	Name           string
	RoundQualified pgtype.Int4
	Code           string
	IsBanned       bool
	Design         int32
	Implementation int32
	Presentation   int32
	Round          int32
	Title          *string
	Description    *string
	Track          *string
	GithubLink     *string
	FigmaLink      *string
	OtherLink      *string
	Title_2        *string
	Description_2  *string
	Track_2        *string
	IsSelected     pgtype.Bool
}

func (q *Queries) GetTeamById(ctx context.Context, id uuid.UUID) (GetTeamByIdRow, error) {
	row := q.db.QueryRow(ctx, getTeamById, id)
	var i GetTeamByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.RoundQualified,
		&i.Code,
		&i.IsBanned,
		&i.Design,
		&i.Implementation,
		&i.Presentation,
		&i.Round,
		&i.Title,
		&i.Description,
		&i.Track,
		&i.GithubLink,
		&i.FigmaLink,
		&i.OtherLink,
		&i.Title_2,
		&i.Description_2,
		&i.Track_2,
		&i.IsSelected,
	)
	return i, err
}

const getTeamByTeamId = `-- name: GetTeamByTeamId :one
SELECT id, name, number_of_people, round_qualified, code, is_banned FROM teams WHERE id = $1
`

func (q *Queries) GetTeamByTeamId(ctx context.Context, id uuid.UUID) (Team, error) {
	row := q.db.QueryRow(ctx, getTeamByTeamId, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.NumberOfPeople,
		&i.RoundQualified,
		&i.Code,
		&i.IsBanned,
	)
	return i, err
}

const getTeamIDByCode = `-- name: GetTeamIDByCode :one
SELECT id FROM teams WHERE code = $1
`

func (q *Queries) GetTeamIDByCode(ctx context.Context, code string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getTeamIDByCode, code)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getTeamMembers = `-- name: GetTeamMembers :many
SELECT first_name , last_name, is_leader, github_profile, reg_no, phone_no
FROM users
Where team_id = $1
`

type GetTeamMembersRow struct {
	FirstName     string
	LastName      string
	IsLeader      bool
	GithubProfile *string
	RegNo         *string
	PhoneNo       pgtype.Text
}

func (q *Queries) GetTeamMembers(ctx context.Context, teamID uuid.NullUUID) ([]GetTeamMembersRow, error) {
	rows, err := q.db.Query(ctx, getTeamMembers, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTeamMembersRow
	for rows.Next() {
		var i GetTeamMembersRow
		if err := rows.Scan(
			&i.FirstName,
			&i.LastName,
			&i.IsLeader,
			&i.GithubProfile,
			&i.RegNo,
			&i.PhoneNo,
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

const getTeamUsers = `-- name: GetTeamUsers :many
SELECT first_name, last_name
From users
Where team_id = $1
`

type GetTeamUsersRow struct {
	FirstName string
	LastName  string
}

func (q *Queries) GetTeamUsers(ctx context.Context, teamID uuid.NullUUID) ([]GetTeamUsersRow, error) {
	rows, err := q.db.Query(ctx, getTeamUsers, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTeamUsersRow
	for rows.Next() {
		var i GetTeamUsersRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamUsersEmails = `-- name: GetTeamUsersEmails :many
SELECT email
FROM users
WHERE team_id = $1
`

func (q *Queries) GetTeamUsersEmails(ctx context.Context, teamID uuid.NullUUID) ([]string, error) {
	rows, err := q.db.Query(ctx, getTeamUsersEmails, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeams = `-- name: GetTeams :many
SELECT teams.id, teams.name, teams.number_of_people, teams.round_qualified, teams.code, teams.is_banned,ideas.title,ideas.description,ideas.track
FROM teams
LEFT JOIN ideas ON ideas.team_id = teams.id
WHERE teams.name ILIKE '%' || $1 || '%'
  AND teams.id > $2
ORDER BY teams.id
LIMIT $3
`

type GetTeamsParams struct {
	Column1 *string
	ID      uuid.UUID
	Limit   int32
}

type GetTeamsRow struct {
	ID             uuid.UUID
	Name           string
	NumberOfPeople int32
	RoundQualified pgtype.Int4
	Code           string
	IsBanned       bool
	Title          *string
	Description    *string
	Track          *string
}

func (q *Queries) GetTeams(ctx context.Context, arg GetTeamsParams) ([]GetTeamsRow, error) {
	rows, err := q.db.Query(ctx, getTeams, arg.Column1, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTeamsRow
	for rows.Next() {
		var i GetTeamsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.NumberOfPeople,
			&i.RoundQualified,
			&i.Code,
			&i.IsBanned,
			&i.Title,
			&i.Description,
			&i.Track,
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

const getUserByID = `-- name: GetUserByID :one
SELECT id, team_id, first_name, last_name, email, phone_no, gender, reg_no, github_profile, password, role, is_leader, is_verified, is_banned, is_profile_complete, is_starred, room_no, hostel_block FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
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
		&i.GithubProfile,
		&i.Password,
		&i.Role,
		&i.IsLeader,
		&i.IsVerified,
		&i.IsBanned,
		&i.IsProfileComplete,
		&i.IsStarred,
		&i.RoomNo,
		&i.HostelBlock,
	)
	return i, err
}

const increaseCountTeam = `-- name: IncreaseCountTeam :exec
UPDATE teams
SET number_of_people = number_of_people + 1
WHERE id = $1
`

func (q *Queries) IncreaseCountTeam(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, increaseCountTeam, id)
	return err
}

const infoQuery = `-- name: InfoQuery :many
SELECT teams.id, name, number_of_people, round_qualified, code, teams.is_banned, users.id, team_id, first_name, last_name, email, phone_no, gender, reg_no, github_profile, password, role, is_leader, is_verified, users.is_banned, is_profile_complete, is_starred, room_no, hostel_block FROM teams INNER JOIN users ON users.team_id = teams.id WHERE teams.id = $1
`

type InfoQueryRow struct {
	ID                uuid.UUID
	Name              string
	NumberOfPeople    int32
	RoundQualified    pgtype.Int4
	Code              string
	IsBanned          bool
	ID_2              uuid.UUID
	TeamID            uuid.NullUUID
	FirstName         string
	LastName          string
	Email             string
	PhoneNo           pgtype.Text
	Gender            string
	RegNo             *string
	GithubProfile     *string
	Password          string
	Role              string
	IsLeader          bool
	IsVerified        bool
	IsBanned_2        bool
	IsProfileComplete bool
	IsStarred         bool
	RoomNo            *string
	HostelBlock       *string
}

func (q *Queries) InfoQuery(ctx context.Context, id uuid.UUID) ([]InfoQueryRow, error) {
	rows, err := q.db.Query(ctx, infoQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InfoQueryRow
	for rows.Next() {
		var i InfoQueryRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.NumberOfPeople,
			&i.RoundQualified,
			&i.Code,
			&i.IsBanned,
			&i.ID_2,
			&i.TeamID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNo,
			&i.Gender,
			&i.RegNo,
			&i.GithubProfile,
			&i.Password,
			&i.Role,
			&i.IsLeader,
			&i.IsVerified,
			&i.IsBanned_2,
			&i.IsProfileComplete,
			&i.IsStarred,
			&i.RoomNo,
			&i.HostelBlock,
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

const kickMemeber = `-- name: KickMemeber :exec
UPDATE users
SET team_id = NULL
WHERE id = $1
`

func (q *Queries) KickMemeber(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, kickMemeber, id)
	return err
}

const leaveTeam = `-- name: LeaveTeam :exec
UPDATE users
SET team_id = NULL
WHERE id = $1
`

func (q *Queries) LeaveTeam(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, leaveTeam, id)
	return err
}

const removeTeamIDFromUsers = `-- name: RemoveTeamIDFromUsers :exec
UPDATE users
SET team_id = NULL
WHERE team_id = $1
`

func (q *Queries) RemoveTeamIDFromUsers(ctx context.Context, teamID uuid.NullUUID) error {
	_, err := q.db.Exec(ctx, removeTeamIDFromUsers, teamID)
	return err
}

const removeUserFromTeam = `-- name: RemoveUserFromTeam :exec
UPDATE users
SET team_id = NULL
WHERE team_id = $1 AND id = $2
`

type RemoveUserFromTeamParams struct {
	TeamID uuid.NullUUID
	ID     uuid.UUID
}

func (q *Queries) RemoveUserFromTeam(ctx context.Context, arg RemoveUserFromTeamParams) error {
	_, err := q.db.Exec(ctx, removeUserFromTeam, arg.TeamID, arg.ID)
	return err
}

const unBanTeam = `-- name: UnBanTeam :exec
UPDATE teams
SET is_banned = FALSE
WHERE id = $1
`

func (q *Queries) UnBanTeam(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, unBanTeam, id)
	return err
}

const updateLeader = `-- name: UpdateLeader :exec
UPDATE users
SET is_leader = $1
WHERE id = $2
`

type UpdateLeaderParams struct {
	IsLeader bool
	ID       uuid.UUID
}

func (q *Queries) UpdateLeader(ctx context.Context, arg UpdateLeaderParams) error {
	_, err := q.db.Exec(ctx, updateLeader, arg.IsLeader, arg.ID)
	return err
}

const updateTeamName = `-- name: UpdateTeamName :exec
UPDATE teams
SET name = $1
WHERE id = $2
`

type UpdateTeamNameParams struct {
	Name string
	ID   uuid.UUID
}

func (q *Queries) UpdateTeamName(ctx context.Context, arg UpdateTeamNameParams) error {
	_, err := q.db.Exec(ctx, updateTeamName, arg.Name, arg.ID)
	return err
}

const updateTeamRound = `-- name: UpdateTeamRound :exec
UPDATE teams
SET round_qualified = $1
WHERE id = $2
`

type UpdateTeamRoundParams struct {
	RoundQualified pgtype.Int4
	ID             uuid.UUID
}

func (q *Queries) UpdateTeamRound(ctx context.Context, arg UpdateTeamRoundParams) error {
	_, err := q.db.Exec(ctx, updateTeamRound, arg.RoundQualified, arg.ID)
	return err
}

const updateUserTeam = `-- name: UpdateUserTeam :exec
UPDATE users
SET team_id = $1, is_leader = $2
WHERE id = $3
`

type UpdateUserTeamParams struct {
	TeamID   uuid.NullUUID
	IsLeader bool
	ID       uuid.UUID
}

func (q *Queries) UpdateUserTeam(ctx context.Context, arg UpdateUserTeamParams) error {
	_, err := q.db.Exec(ctx, updateUserTeam, arg.TeamID, arg.IsLeader, arg.ID)
	return err
}
