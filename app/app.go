package app

import errs "github.com/pkg/errors"

type GopherApp interface {
	SaveGopher(g *Gopher) (string, error)
	GetGopher(id string) (*Gopher, error)
}

type App struct {
	db GopherDB
}

func NewApp(gdb GopherDB) *App {
	return &App{
		db: gdb,
	}
}

func (a *App) SaveGopher(g *Gopher) (string, error) {
	gopherId, err := a.db.SaveGopher(g)
	if err != nil {
		return "", errs.Wrap(err, "failed in saving Gopher to DB")
	}
	return gopherId, nil
}

func (a *App) GetGopher(id string) (*Gopher, error) {
	gopher, err := a.db.GetGopher(id)
	if err != nil {
		return nil, errs.Wrap(err, "failed in getting Gopher from DB")
	}
	return gopher, nil
}
