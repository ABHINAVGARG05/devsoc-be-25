-- name: GetTeamIDByCode :one
SELECT id FROM teams WHERE code = $1;

-- name: GetTeams :many
SELECT * FROM teams;

-- name: GetTeamById :one
SELECT * FROM teams WHERE id = $1;

-- name: FindTeam :one
SELECT id,name,code,round_qualified FROM teams 
WHERE code = $1 
LIMIT 1;

-- name: KickMemeber :exec
UPDATE users
SET team_id = NULL
WHERE id = $1;

-- name: LeaveTeam :exec
UPDATE users 
SET team_id = NULL
WHERE id = $1;

-- name: CreateTeam :one
INSERT INTO teams (
    id, name, number_of_people, round_qualified, code
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM teams 
WHERE id = $1;

-- name: CountTeamMembers :one
SELECT COUNT(*) FROM users
WHERE team_id = $1;

-- name: AddUserToTeam :exec
UPDATE users
SET team_id = $1
WHERE id = $2;

-- name: RemoveUserFromTeam :exec
UPDATE users
SET team_id = NULL
WHERE team_id = $1 AND id = $2;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUserTeam :exec
UPDATE users
SET team_id = $1, is_leader = $2
WHERE id = $3;

-- name: IncreaseCountTeam :exec
UPDATE teams
SET number_of_people = number_of_people + 1
WHERE id = $1;

-- name: DecreaseUserCountTeam :exec
UPDATE teams
SET number_of_people = number_of_people - 1
WHERE id = $1;

-- name: RemoveTeamIDFromUsers :exec
UPDATE users
SET team_id = NULL
WHERE team_id = $1;

-- name: UpdateLeader :exec
UPDATE users
SET is_leader = $1
WHERE id = $2;

-- name: UpdateTeamName :exec
UPDATE teams
SET name = $1
WHERE id = $2;