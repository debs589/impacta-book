package application

import (
	"api/internal/repositories"
	"api/internal/router/routes"
	"api/internal/services"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type ConfigApplicationDefault struct {
	Db   *mysql.Config
	Addr string
}

func NewApplicationDefault(config *ConfigApplicationDefault) *ApplicationDefault {

	defaultCfg := &ConfigApplicationDefault{
		Db:   nil,
		Addr: ":5000",
	}

	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}

		if config.Addr != "" {
			defaultCfg.Addr = config.Addr
		}
	}

	return &ApplicationDefault{
		cfgDB:   defaultCfg.Db,
		cfgAddr: defaultCfg.Addr,
	}
}

type ApplicationDefault struct {
	cfgDB   *mysql.Config
	cfgAddr string
	db      *sql.DB
	router  *chi.Mux
}

func (a *ApplicationDefault) TearDown() {
	// close db
	if a.db != nil {
		a.db.Close()
	}
}

func (a *ApplicationDefault) SetUp() (err error) {
	a.db, err = sql.Open("mysql", a.cfgDB.FormatDSN())

	if err != nil {
		log.Fatalf("error opening db: %s", err.Error())
	}

	if err = a.db.Ping(); err != nil {
		log.Fatalf("error pinging db: %s", err.Error())
	}

	router := chi.NewRouter()

	userRepo := repositories.NewUserRepository(a.db)
	userService := services.NewUserService(userRepo)
	err = routes.NewUserRoutes(router, userService)

	if err != nil {
		panic(err)
	}

	err = routes.NewLoginRoutes(router, userService)
	if err != nil {
		panic(err)
	}

	a.router = router

	return nil
}

func (a *ApplicationDefault) Run() (err error) {
	defer a.db.Close()
	log.Printf("starting server at %s\n", a.cfgAddr)

	err = http.ListenAndServe(a.cfgAddr, a.router)
	return
}
