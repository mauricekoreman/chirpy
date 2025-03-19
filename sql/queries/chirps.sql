-- name: CreateChirp :one
insert into chirps (id, created_at, updated_at, body, user_id)
values (gen_random_uuid(), NOW(), NOW(), $1, $2)
returning *;

-- name: GetAllChirps :many
select *
from chirps
order by created_at asc;

-- name: GetChirp :one
select *
from chirps
where id = $1;