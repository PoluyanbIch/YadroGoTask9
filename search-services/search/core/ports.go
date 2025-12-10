package core

import (
	"context"
)

type Searcher interface {
	Search(context.Context, string, int) ([]Comic, error)
	ISearch(context.Context, string, int) ([]Comic, error)
}

type DB interface {
	Read(context.Context) ([]DBComic, error)
}

type Words interface {
	Norm(ctx context.Context, phrase string) ([]string, error)
}

type Subscriber interface {
	Subscribe(context.Context, string, func(context.Context, []byte) error) error
}

type Initiator interface {
	BuildIndex(context.Context) error
}
