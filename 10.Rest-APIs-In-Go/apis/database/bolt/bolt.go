package bolt

import "context"

type Bolt struct {

}

func New(ctx context.Context) *Bolt {
	return &Bolt {}
}