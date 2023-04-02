package repository

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(url string) (*PostgresUserRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresUserRepository{db: db}, nil
}

func (p *PostgresUserRepository) Insert(ctx context.Context, user domain.User) error {
	_, err := p.db.ExecContext(ctx,
		"INSERT INTO users (id, email, password) VALUES ($1, $2, $3)",
		user.Id, user.Email, user.Password)

	return err
}

func (p *PostgresUserRepository) GetUserById(ctx context.Context, id domain.UserId) (domain.User, error) {
	response, err := p.db.QueryContext(ctx,
		"SELECT id email FROM user WHERE id = $1", id)
	defer func() {
		err = response.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	var user = domain.User{}

	for response.Next() {
		if err = response.Scan(&user.Id, &user.Email); err == nil {
			return user, nil
		}
	}

	if err = response.Err(); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (p *PostgresUserRepository) Close() error {
	return p.db.Close()
}
