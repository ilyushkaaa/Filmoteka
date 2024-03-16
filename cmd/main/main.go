package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ilyushkaaa/Filmoteka/config"
	actorDelivery "github.com/ilyushkaaa/Filmoteka/internal/actors/delivery"
	actorRepo "github.com/ilyushkaaa/Filmoteka/internal/actors/repo"
	actorUseCase "github.com/ilyushkaaa/Filmoteka/internal/actors/usecase"
	filmDelivery "github.com/ilyushkaaa/Filmoteka/internal/films/delivery"
	filmRepo "github.com/ilyushkaaa/Filmoteka/internal/films/repo"
	filmUseCase "github.com/ilyushkaaa/Filmoteka/internal/films/usecase"
	"github.com/ilyushkaaa/Filmoteka/internal/middleware"
	sessionRepo "github.com/ilyushkaaa/Filmoteka/internal/session/repo"
	sessionUseCase "github.com/ilyushkaaa/Filmoteka/internal/session/usecase"
	userDelivery "github.com/ilyushkaaa/Filmoteka/internal/users/delivery"
	userRepo "github.com/ilyushkaaa/Filmoteka/internal/users/repo"
	userUseCase "github.com/ilyushkaaa/Filmoteka/internal/users/usecase"
	"github.com/ilyushkaaa/Filmoteka/pkg/dbinit"
	"go.uber.org/zap"
)

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Printf("error in logger start")
		return
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in logger sync")
		}
	}()

	err = config.GetConfig()
	if err != nil {
		logger.Errorf("error in getting config")
		return
	}

	pgxDB, err := dbinit.GetPostgres()
	fmt.Println("eeee")

	if err != nil {
		logger.Errorf("error in connection to postgres: %s", err)
		return
	}
	logger.Infof("connected to postgres")
	defer func() {
		err = pgxDB.Close()
		if err != nil {
			logger.Errorf("error in close connection to mysql: %s", err)
		}
	}()

	redisConn, err := dbinit.GetRedis()
	if err != nil {
		logger.Infof("error on connection to redis: %s", err.Error())
	}
	defer func() {
		err = redisConn.Close()
		if err != nil {
			logger.Infof("error on redis close: %s", err.Error())
		}
	}()
	logger.Infof("connected to redis")

	sr := sessionRepo.NewSessionRepo(redisConn)
	su := sessionUseCase.NewSessionUseCase(sr)

	ur := userRepo.NewUserRepo(pgxDB)
	uu := userUseCase.NewUserUseCase(ur)
	uh := userDelivery.NewUserHandler(uu, su)

	fr := filmRepo.NewFilmRepo(pgxDB, logger)
	fu := filmUseCase.NewFilmUseCase(fr)
	fh := filmDelivery.NewFilmHandler(fu)

	ar := actorRepo.NewActorRepo(pgxDB, logger)
	au := actorUseCase.NewActorUseCase(ar)
	ah := actorDelivery.NewActorHandler(au)

	mw := middleware.NewMiddleware(su, uu)

	router := mux.NewRouter()
	adminRouter := mux.NewRouter()
	authRouter := mux.NewRouter()

	router.PathPrefix("/admin").Handler(adminRouter)
	router.PathPrefix("/logout").Handler(authRouter)

	router.HandleFunc("/actor/{ACTOR_ID}", ah.GetActorByID).Methods(http.MethodGet)
	router.HandleFunc("/actors", ah.GetActors).Methods(http.MethodGet)

	router.HandleFunc("/film/{FILM_ID}", fh.GetFilmByID).Methods(http.MethodGet)
	router.HandleFunc("/film/{FILM_ID}", fh.GetFilmByID).Methods(http.MethodGet)
	router.HandleFunc("/films", fh.GetFilms).Methods(http.MethodGet)
	router.HandleFunc("/film/search/{SEARCH_STR}", fh.GetFilmsBySearch).Methods(http.MethodGet)

	router.HandleFunc("/login", uh.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", uh.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/logout", uh.Logout).Methods(http.MethodPost)

	adminRouter.HandleFunc("/admin/actor/{ACTOR_ID}", ah.DeleteActor).Methods(http.MethodDelete)
	adminRouter.HandleFunc("/admin/actor", ah.UpdateActor).Methods(http.MethodPut)
	adminRouter.HandleFunc("/admin/actor", ah.AddActor).Methods(http.MethodPost)

	adminRouter.HandleFunc("/admin/film/{FILM_ID}", fh.DeleteFilm).Methods(http.MethodDelete)
	adminRouter.HandleFunc("/admin/film", fh.UpdateFilm).Methods(http.MethodPut)
	adminRouter.HandleFunc("/admin/film", fh.AddFilm).Methods(http.MethodPost)

	router.Use(mw.RequestInitMiddleware)
	router.Use(mw.AccessLog)

	adminRouter.Use(mw.AuthMiddleware)
	adminRouter.Use(mw.AdminMiddleware)

	authRouter.Use(mw.AuthMiddleware)

	port := os.Getenv("appPort")

	logger.Infow("starting server",
		"type", "START",
		"addr", port,
	)
	err = http.ListenAndServe(port, router)
	if err != nil {
		logger.Fatalf("errror in server start")
	}

}
