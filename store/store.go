package store

import (
	"context"
	"fmt"

	"github.com/Asendar1/go-url-shortener/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool 	*pgxpool.Pool
	queries	*db.Queries
}

func Connect (connString string) (*Store, error) {
	pool, err := pgxpool.New(context.Background(), connString);
	if err != nil {
		return nil, fmt.Errorf("pgx error: %w", err)
	}
	return &Store{
		pool: pool,
		queries: db.New(pool),
	}, nil
}

func (s *Store) CreateURL (short, long string) error {
	_, err := s.queries.CreateURL(context.Background(), db.CreateURLParams{
		ShortCode:	short,
		LongUrl:	long,
	})
	return err
}

func (s *Store) GetByShortCode(shortCode string) (db.Url, error) {
	url, err := s.queries.GetByShortCode(context.Background(), shortCode)
	return url, err
}

func (s *Store) UpdateClicks(shortCode string) error {
	err := s.queries.UpdateClicks(context.Background(), shortCode)
	return err
}

func (s *Store) UpdateLongUrl(shortCode, longUrl string) error {
	err := s.queries.UpdateLongUrl(context.Background(), db.UpdateLongUrlParams{
		ShortCode: shortCode,
		LongUrl: longUrl,
	})
	return err
}

func (s *Store) DeleteByShortCode(shortCode string) error {
	err := s.queries.DeleteByShortCode(context.Background(), shortCode)
	return err
}

