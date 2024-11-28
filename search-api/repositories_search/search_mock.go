package repositories_search

import (
	"search/dao_search"
)

type Mock struct {
	data map[int64]dao_search.Search
}

func NewMock() Mock {
	return Mock{
		data: make(map[int64]dao_search.Search),
	}
}
