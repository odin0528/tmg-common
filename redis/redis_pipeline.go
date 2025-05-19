package redis

import (
	"context"
	"time"
)

func (rp *PipelineWrapper) Set(key string, value interface{}, expiration time.Duration) {
	rp.pipe.Set(context.TODO(), key, value, expiration)
}

func (rp *PipelineWrapper) Get(key string) {
	rp.pipe.Get(context.TODO(), key)
}

func (rp *PipelineWrapper) Exec(ctx context.Context) ([]CmdResult, error) {
	cmds, err := rp.pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return processPipelineResults(cmds), nil
}

func (rp *PipelineWrapper) Discard() error {
	return nil
}
