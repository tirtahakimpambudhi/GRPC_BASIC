package model

import "context"

type Operation func(ctx context.Context) error
type Operations map[string]Operation

type Ratings struct {
	Count uint32
	Sum   float64
}
