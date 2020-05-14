package app

import (
	"errors"

	errs "github.com/pkg/errors"
)

type DBSuccess struct {
}

func (d *DBSuccess) SaveGopher(g *Gopher) (string, error) {
	return "successid", nil
}

func (d *DBSuccess) GetGopher(id string) (*Gopher, error) {
	return &Gopher{
		Id:   id,
		Name: "",
	}, nil
}

type DBFailure struct {
}

func (d *DBFailure) SaveGopher(g *Gopher) (string, error) {
	return "", errs.Wrap(errors.New("failure in saving to DB"), "failed in saving Gopher")
}

func (d *DBFailure) GetGopher(id string) (*Gopher, error) {
	return nil, errs.Wrap(errors.New("failure in getting from DB"), "failed in fetching Gopher")
}
