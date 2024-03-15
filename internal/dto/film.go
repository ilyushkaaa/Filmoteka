package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ilyushkaaa/Filmoteka/internal/films/entity"
	"github.com/ilyushkaaa/Filmoteka/pkg/validator"
)

type (
	FilmAdd struct {
		Name          string    `json:"name" valid:"required,length(1|150)"`
		Description   string    `json:"description" valid:"required,length(1|1000)"`
		DateOfRelease time.Time `json:"date_of_release" valid:"required"`
		Rating        float64   `json:"rating" valid:"required,range(0|10)"`
		ActorIDs      []uint64  `json:"actor_ids"`
	}
	FilmUpdate struct {
		ID            uint64    `json:"id" valid:"required"`
		Name          string    `json:"name" valid:"required,length(1|150)"`
		Description   string    `json:"description" valid:"required,length(1|1000)"`
		DateOfRelease time.Time `json:"date_of_release" valid:"required"`
		Rating        float64   `json:"rating" valid:"required,range(0|10)"`
		ActorIDs      []uint64  `json:"actor_ids"`
	}
)

func (f *FilmAdd) Validate() []string {
	_, err := govalidator.ValidateStruct(f)
	return validator.CollectErrors(err)
}

func (f *FilmAdd) GetFilmAndActorIDs() (entity.Film, []uint64) {
	film := entity.Film{
		Name:          f.Name,
		Description:   f.Description,
		DateOfRelease: f.DateOfRelease,
		Rating:        f.Rating,
	}
	IDs := make([]uint64, len(f.ActorIDs))
	for i := range f.ActorIDs {
		IDs[i] = f.ActorIDs[i]
	}
	return film, IDs
}

func (f *FilmUpdate) Validate() []string {
	_, err := govalidator.ValidateStruct(f)
	return validator.CollectErrors(err)
}

func (f *FilmUpdate) GetFilmAndActorIDs() (entity.Film, []uint64) {
	film := entity.Film{
		ID:            f.ID,
		Name:          f.Name,
		Description:   f.Description,
		DateOfRelease: f.DateOfRelease,
		Rating:        f.Rating,
	}
	IDs := make([]uint64, len(f.ActorIDs))
	for i := range f.ActorIDs {
		IDs[i] = f.ActorIDs[i]
	}
	return film, IDs
}
