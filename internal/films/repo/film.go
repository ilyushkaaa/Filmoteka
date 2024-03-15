package repo

import (
	"database/sql"
	"errors"

	entityActor "github.com/ilyushkaaa/Filmoteka/internal/actors/entity"
	entityFilm "github.com/ilyushkaaa/Filmoteka/internal/films/entity"
	"go.uber.org/zap"
)

type FilmRepo interface {
	GetFilms(sortParam string) ([]entityFilm.Film, error)
	GetFilmByID(filmID uint64) (*entityFilm.Film, error)
	AddFilm(film entityFilm.Film, actors []entityActor.Actor) (uint64, error)
	UpdateFilm(film entityFilm.Film, actors []entityActor.Actor) (bool, error)
	GetFilmsBySearch(searchStr string) ([]entityFilm.Film, error)
}

type FilmRepoPG struct {
	db        *sql.DB
	zapLogger *zap.SugaredLogger
}

func NewFilmRepo(db *sql.DB, zapLogger *zap.SugaredLogger) *FilmRepoPG {
	return &FilmRepoPG{
		db:        db,
		zapLogger: zapLogger,
	}
}

func (r *FilmRepoPG) GetFilms(sortParam string) ([]entityFilm.Film, error) {
	if sortParam == "" {
		sortParam = "rating"
	}

	rows, err := r.db.Query("SELECT f.id, f.name, f.description, f.date_of_release, f.rating FROM films f ORDER BY " + sortParam + " DESC")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			r.zapLogger.Errorf("error in closing query rows: %s", err)
		}
	}(rows)
	films := make([]entityFilm.Film, 0)
	for rows.Next() {
		film := entityFilm.Film{}
		err = rows.Scan(&film.ID, &film.Name, &film.Description, &film.DateOfRelease, &film.Rating)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return films, nil
}

func (r *FilmRepoPG) GetFilmByID(filmID uint64) (*entityFilm.Film, error) {
	film := &entityFilm.Film{}
	err := r.db.
		QueryRow("SELECT id, name, description, date_of_release, rating FROM films WHERE id = ?", filmID).
		Scan(&film.ID, &film.Name, &film.Description, &film.DateOfRelease, &film.Rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return film, nil
}
func (r *FilmRepoPG) AddFilm(film entityFilm.Film, actors []entityActor.Actor) (uint64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				r.zapLogger.Errorf("error in transaction rollback")
			}
		}
	}()

	var lastInsertId uint64
	err = tx.QueryRow("INSERT INTO films (name, description, date_of_release, rating) VALUES ($1, $2, $3, $4) RETURNING id",
		film.Name, film.Description, film.DateOfRelease, film.Rating).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	for _, actor := range actors {
		var actorID uint64
		err = tx.QueryRow("SELECT id FROM actors WHERE id = $1", actor.ID).Scan(&actorID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return 0, nil
			}
			return 0, err
		}

		_, err = tx.Exec("INSERT INTO film_actors (film_id, actor_id) VALUES ($1, $2)",
			lastInsertId, actor.ID)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}
func (r *FilmRepoPG) UpdateFilm(film entityFilm.Film, actors []entityActor.Actor) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				r.zapLogger.Errorf("error in transaction rollback")
			}
		}
	}()

	_, err = tx.Exec("UPDATE films SET name = $1, description = $2, date_of_release = $3, rating = $4 WHERE id = $5",
		film.Name, film.Description, film.DateOfRelease, film.Rating, film.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	_, err = tx.Exec("DELETE FROM film_actors WHERE film_id = $1", film.ID)
	if err != nil {
		return false, err
	}

	for _, actor := range actors {
		var actorID uint64
		err = tx.QueryRow("SELECT id FROM actors WHERE id = $1", actor.ID).Scan(&actorID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return false, nil
			}
			return false, err
		}
		_, err = tx.Exec("INSERT INTO film_actors (film_id, actor_id) VALUES ($1, $2)", film.ID, actor.ID)
		if err != nil {
			return false, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FilmRepoPG) GetFilmsBySearch(searchStr string) ([]entityFilm.Film, error) {
	rows, err := r.db.Query(`
    SELECT DISTINCT f.id, f.name, f.description, f.date_of_release, f.rating
	FROM films f
	LEFT JOIN film_actors fa ON f.id = fa.film_id
	LEFT JOIN actors a ON fa.actor_id = a.id
	WHERE LOWER(f.name) LIKE LOWER('%' || $1 || '%')
    OR LOWER(a.name || ' ' || a.surname) LIKE LOWER('%' || $1 || '%');
`,
		searchStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			r.zapLogger.Errorf("error in closing query rows: %s", err)
		}
	}(rows)
	films := make([]entityFilm.Film, 0)
	for rows.Next() {
		film := entityFilm.Film{}
		err = rows.Scan(&film.ID, &film.Name, &film.Description, &film.DateOfRelease, &film.Rating)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return films, nil
}
