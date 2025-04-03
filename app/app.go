package app

import (
	"github.com/mojtabamovahedi/url-shorter/config"
	"github.com/mojtabamovahedi/url-shorter/internal/service"
	"github.com/mojtabamovahedi/url-shorter/pkg/postgres"
	"gorm.io/gorm"
)

type App struct {
	cfg         config.Config
	db          *gorm.DB
	linkService service.LinkService
}

func NewApp(cfg config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}
	err := a.setDB()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func MustNewApp(cfg config.Config) *App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
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

	app.db = db
	return nil
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

func (app *App) Config() config.Config {
	return app.cfg
}
