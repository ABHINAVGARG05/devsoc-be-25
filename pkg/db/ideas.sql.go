// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ideas.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createIdea = `-- name: CreateIdea :one
INSERT INTO ideas (
  id, title, description, track, team_id, is_selected
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, title, description, track, team_id, is_selected, created_at, updated_at
`

type CreateIdeaParams struct {
	ID          uuid.UUID
	Title       string
	Description string
	Track       string
	TeamID      uuid.UUID
	IsSelected  bool
}

func (q *Queries) CreateIdea(ctx context.Context, arg CreateIdeaParams) (Idea, error) {
	row := q.db.QueryRow(ctx, createIdea,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Track,
		arg.TeamID,
		arg.IsSelected,
	)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Track,
		&i.TeamID,
		&i.IsSelected,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteIdea = `-- name: DeleteIdea :exec
DELETE FROM ideas
WHERE id = $1
`

func (q *Queries) DeleteIdea(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteIdea, id)
	return err
}

const getIdea = `-- name: GetIdea :one
SELECT id, title, description, track, team_id, is_selected, created_at, updated_at FROM ideas
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetIdea(ctx context.Context, id uuid.UUID) (Idea, error) {
	row := q.db.QueryRow(ctx, getIdea, id)
	var i Idea
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Track,
		&i.TeamID,
		&i.IsSelected,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getIdeaByTeamID = `-- name: GetIdeaByTeamID :one
SELECT id, title, description, track
FROM ideas
WHERE team_id = $1
LIMIT 1
`

type GetIdeaByTeamIDRow struct {
	ID          uuid.UUID
	Title       string
	Description string
	Track       string
}

func (q *Queries) GetIdeaByTeamID(ctx context.Context, teamID uuid.UUID) (GetIdeaByTeamIDRow, error) {
	row := q.db.QueryRow(ctx, getIdeaByTeamID, teamID)
	var i GetIdeaByTeamIDRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Track,
	)
	return i, err
}

const listIdeas = `-- name: ListIdeas :many
SELECT id, title, description, track, team_id, is_selected, created_at, updated_at FROM ideas
ORDER BY created_at DESC
`

func (q *Queries) ListIdeas(ctx context.Context) ([]Idea, error) {
	rows, err := q.db.Query(ctx, listIdeas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Idea
	for rows.Next() {
		var i Idea
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Track,
			&i.TeamID,
			&i.IsSelected,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateIdea = `-- name: UpdateIdea :exec
UPDATE ideas
SET title = $2,
    description = $3,
    track = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE team_id = $1
`

type UpdateIdeaParams struct {
	TeamID      uuid.UUID
	Title       string
	Description string
	Track       string
}

func (q *Queries) UpdateIdea(ctx context.Context, arg UpdateIdeaParams) error {
	_, err := q.db.Exec(ctx, updateIdea,
		arg.TeamID,
		arg.Title,
		arg.Description,
		arg.Track,
	)
	return err
}
