package app

type GopherDB interface {
	SaveGopher(g *Gopher) (string, error)
	GetGopher(id string) (*Gopher, error)
}
