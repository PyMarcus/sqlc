-- name: CreateTransfer :one
insert into transfers(from_account_id,
  to_account_id,
  amount)
values ($1, $2, $3) 
returning *;

-- name: GetTransfer :one
select * from transfers
where id = $1 limit 1;

-- name: ListTransfers :many
select * from transfers
order by from_account_id limit $1 offset $2;

