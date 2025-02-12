// Code generated from Pkl module `Control`. DO NOT EDIT.
package control

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Control struct {
	Rl *Controller `pkl:"rl"`

	Cb *Controller `pkl:"cb"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Control
func LoadFromPath(ctx context.Context, path string) (ret *Control, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Control
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Control, error) {
	var ret Control
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
