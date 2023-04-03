package repository

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
	"database/sql"
	"log"
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

func (p *PostgresPostRepository) GetPostById(ctx context.Context, postId domain.PostId) (domain.Post, error) {
	response, err := p.db.QueryContext(ctx,
		"SELECT id, post_content, create_at, user_id FROM posts WHERE id = $1", postId)
	defer func() {
		err = response.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	var post = domain.Post{}
	for response.Next() {
		if err = response.Scan(&post.Id, &post.PostContent, &post.CreatedAt, &post.UserId); err != nil {
			return post, nil
		}
	}

	if err = response.Err(); err != nil {
		return post, err
	}
	return post, nil
}

func (p *PostgresPostRepository) Update(ctx context.Context, post domain.Post) error {
	_, err := p.db.ExecContext(ctx,
		"UPDATE posts SET post_content = $1 WHERE id = $2 and user_id = $3",
		post.PostContent, post.Id, post.UserId)

	return err
}

func (p *PostgresPostRepository) Delete(ctx context.Context, postId domain.PostId, userId domain.UserId) error {
	_, err := p.db.ExecContext(ctx,
		"DELETE FROM posts WHERE id = $1 AND user_id = $2",
		postId, userId)

	return err
}

func (p *PostgresPostRepository) Close() error {
	return p.db.Close()
}
