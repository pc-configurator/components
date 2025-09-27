package usecase

type Postgres interface{}

type Component struct {
	postgres Postgres
}

func New(p Postgres) *Component {
	return &Component{
		postgres: p,
	}
}
