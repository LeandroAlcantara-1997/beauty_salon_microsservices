package repository

type RepositoryI interface {
	Querier

	Execer
}

type Querier interface {
}

type Execer interface {
}
