package postgres

type repository struct{}

func New() (*repository, error) {
	return &repository{}, nil
}
