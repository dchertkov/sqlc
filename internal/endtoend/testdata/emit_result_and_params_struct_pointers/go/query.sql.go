// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const getAll = `-- name: GetAll :many
SELECT a, b FROM foo
`

func (q *Queries) GetAll(ctx context.Context) ([]*Foo, error) {
	rows, err := q.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.A, &i.B); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllAByB = `-- name: GetAllAByB :many
SELECT a FROM foo WHERE b = ?
`

func (q *Queries) GetAllAByB(ctx context.Context, b sql.NullInt32) ([]sql.NullInt32, error) {
	rows, err := q.db.QueryContext(ctx, getAllAByB, b)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullInt32
	for rows.Next() {
		var a sql.NullInt32
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOne = `-- name: GetOne :one
SELECT a, b FROM foo WHERE a = ? LIMIT 1
`

func (q *Queries) GetOne(ctx context.Context, a sql.NullInt32) (*Foo, error) {
	row := q.db.QueryRowContext(ctx, getOne, a)
	var i Foo
	err := row.Scan(&i.A, &i.B)
	return &i, err
}