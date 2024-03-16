package dto

import (
	entityActor "github.com/ilyushkaaa/Filmoteka/internal/actors/entity"
	entityFilm "github.com/ilyushkaaa/Filmoteka/internal/films/entity"
)

type ActorWithFilms struct {
	Actor entityActor.Actor
	Films []entityFilm.Film
}
