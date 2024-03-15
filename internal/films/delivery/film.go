package delivery

import (
	"net/http"

	"github.com/ilyushkaaa/Filmoteka/internal/films/usecase"
)

type FilmHandler struct {
	filmUseCase usecase.FilmUseCase
}

func NewFilmHandler(filmUseCase usecase.FilmUseCase) *FilmHandler {
	return &FilmHandler{
		filmUseCase: filmUseCase,
	}
}

func (h *FilmHandler) GetFilms(w http.ResponseWriter, r *http.Request) {

}

func (h *FilmHandler) GetFilmByID(w http.ResponseWriter, r *http.Request) {

}

func (h *FilmHandler) AddFilm(w http.ResponseWriter, r *http.Request) {

}

func (h *FilmHandler) UpdateFilm(w http.ResponseWriter, r *http.Request) {

}

func (h *FilmHandler) GetFilmsBySearch(w http.ResponseWriter, r *http.Request) {

}
