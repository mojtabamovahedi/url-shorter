package app

import (
	"fmt"
	"github.com/mojtabamovahedi/url-shorter/config"
	"github.com/mojtabamovahedi/url-shorter/internal/service"
	"github.com/mojtabamovahedi/url-shorter/pkg/cache"
	"github.com/mojtabamovahedi/url-shorter/pkg/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type App struct {
	cfg         config.Config
	db          *gorm.DB
	provider    cache.Provider
	linkService service.LinkService
}

func NewApp(cfg config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}
	a.setCache()

	err := a.setDB()

	if err != nil {
		return nil, err
	}
	return a, nil
}

func MustNewApp(cfg config.Config) *App {
	a, err := NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func (app *App) SetLinkService(linkService service.LinkService) {
	app.linkService = linkService
}

func (app *App) LinkService() service.LinkService {
	return app.linkService
}

func (app *App) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   app.cfg.DB.Username,
		Pass:   app.cfg.DB.Password,
		Host:   app.cfg.DB.Host,
		Port:   app.cfg.DB.Port,
		DBName: app.cfg.DB.Database,
		Schema: app.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	log.Println("Successfully connected to database.")

	app.db = db
	return nil
}

func (app *App) setCache() {
	app.provider = cache.NewRedisCacheConnection(
		fmt.Sprintf("%s:%d", app.cfg.Redis.Host, app.cfg.Redis.Port),
		time.Duration(app.cfg.Redis.TTL)*time.Hour)
	log.Println("Successfully connected to redis.")
}

func (app *App) Close() {
	if app.db != nil {
		db, _ := app.db.DB()
		_ = db.Close()
	}
}

func (app *App) DB() *gorm.DB {
	return app.db
}

func (app *App) Provider() cache.Provider {
	return app.provider
}

func (app *App) Config() config.Config {
	return app.cfg
}
