package model

import "context"

type Operation func(ctx context.Context) error
type Operations map[string]Operation
