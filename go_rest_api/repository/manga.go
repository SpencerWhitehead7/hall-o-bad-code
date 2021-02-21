package repository

import (
	"context"
	"fmt"

	"github.com/SpencerWhitehead7/hall-o-bad-code/go_rest_api/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// MangaRepo is an interface for interacting with the DB to handle manga data.
type MangaRepo interface {
	FindAll() ([]Manga, error)
}

// Manga implements MangaRepo.
type Manga struct{ db *pgxpool.Pool }

// FindAll returns all manga in the manga table, sorted alphabetically
func (r *Manga) FindAll() ([]*model.Manga, error) {
	var mangas []*model.Manga

	rows, err := r.db.Query(context.Background(), `SELECT * FROM manga;`) // would need a more real query to get things like genres, chapter count, other real data, etc
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		manga := model.Manga{}

		err = rows.Scan(
			&manga.ID,
			&manga.Name,
			&manga.OtherNames,
			&manga.Description,
			&manga.Demo,
			&manga.StartDate,
			&manga.EndDate,
		)
		if err != nil {
			fmt.Printf("Manga row scan failed: %v\n", err)
		}

		mangas = append(mangas, &manga)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return mangas, nil
}

// MangaFactory creates new MangaRepositories.
func MangaFactory(db *pgxpool.Pool) *Manga {
	return &Manga{db: db}
}
