package usecase

import (
	entityActor "github.com/ilyushkaaa/Filmoteka/internal/actors/entity"
	entityFilm "github.com/ilyushkaaa/Filmoteka/internal/films/entity"
	"github.com/ilyushkaaa/Filmoteka/internal/films/repo"
)

type FilmUseCase interface {
	GetFilms(sortParam string) ([]entityFilm.Film, error)
	GetFilmByID(filmID uint64) (*entityFilm.Film, error)
	AddFilm(film entityFilm.Film, actors []entityActor.Actor) (*entityFilm.Film, error)
	UpdateFilm(film entityFilm.Film, actors []entityActor.Actor) (bool, error)
	GetFilmsBySearch(searchStr string) ([]entityFilm.Film, error)
}

type FilmUseCaseApp struct {
	filmRepo repo.FilmRepo
}

func NewFilmUseCase(filmRepo repo.FilmRepo) *FilmUseCaseApp {
	return &FilmUseCaseApp{
		filmRepo: filmRepo,
	}
}

func (r *FilmUseCaseApp) GetFilms(sortParam string) ([]entityFilm.Film, error) {
	films, err := r.filmRepo.GetFilms(sortParam)
	if err != nil {
		return nil, err
	}
	if films == nil || len(films) == 0 {
		return nil, ErrFilmsNotFound
	}
	return films, nil
}

func (r *FilmUseCaseApp) GetFilmByID(filmID uint64) (*entityFilm.Film, error) {
	film, err := r.filmRepo.GetFilmByID(filmID)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, ErrFilmNotFound
	}
	return film, nil
}

func (r *FilmUseCaseApp) AddFilm(film entityFilm.Film, actors []entityActor.Actor) (*entityFilm.Film, error) {
	filmID, err := r.filmRepo.AddFilm(film, actors)
	if err != nil {
		return nil, err
	}
	film.ID = filmID
	return &film, nil
}

func (r *FilmUseCaseApp) UpdateFilm(film entityFilm.Film, actors []entityActor.Actor) (bool, error) {
	wasUpdated, err := r.filmRepo.UpdateFilm(film, actors)
	if err != nil {
		return false, nil
	}
	if !wasUpdated {
		return false, ErrBadFilmUpdateData
	}
	return true, nil
}

func (r *FilmUseCaseApp) GetFilmsBySearch(searchStr string) ([]entityFilm.Film, error) {
	films, err := r.filmRepo.GetFilmsBySearch(searchStr)
	if err != nil {
		return nil, err
	}
	if films == nil || len(films) == 0 {
		return nil, ErrFilmsNotFound
	}
	return films, nil
}
