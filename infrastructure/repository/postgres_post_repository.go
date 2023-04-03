package repository

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
	"database/sql"
)

type PostgresPostRepository struct {
	db *sql.DB
}

func NewPostgresPostRepository(url string) (*PostgresPostRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresPostRepository{db: db}, nil
}

func (p *PostgresPostRepository) Insert(ctx context.Context, post domain.Post) error {
	_, err := p.db.ExecContext(ctx,
		"INSERT INTO posts (id, post_content, user_id) VALUES ($1, $2, $3)",
		post.Id, post.PostContent, post.UserId)

	return err
}

func (p *PostgresPostRepository) Close() error {
	return p.db.Close()
}
