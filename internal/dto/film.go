package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ilyushkaaa/Filmoteka/pkg/validator"
)

type FilmAdd struct {
	Name          string    `json:"name" valid:"required,length(1|150)"`
	Description   string    `json:"description" valid:"required,length(1|1000)"`
	DateOfRelease time.Time `json:"date_of_release" valid:"required"`
	Rating        float64   `json:"rating" valid:"required,range(0|10)"`
	ActorIDs      []uint64  `json:"actor_ids"`
}

func (f *FilmAdd) Validate() []string {
	_, err := govalidator.ValidateStruct(f)
	return validator.CollectErrors(err)
}
