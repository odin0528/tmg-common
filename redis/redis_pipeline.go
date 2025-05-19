package redis

import (
	"context"
	"time"
)

func (rp *PipelineWrapper) Set(key string, value interface{}, expiration time.Duration) {
	cmd := rp.pipe.Set(context.TODO(), key, value, expiration)
	rp.cmds = append(rp.cmds, cmd)
}

func (rp *PipelineWrapper) Get(key string) {
	cmd := rp.pipe.Get(context.TODO(), key)
	rp.cmds = append(rp.cmds, cmd)
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

func (rp *PipelineWrapper) HSet(key string, field string, value interface{}) {
	cmd := rp.pipe.HSet(context.TODO(), key, field, value)
	rp.cmds = append(rp.cmds, cmd)
}

func (rp *PipelineWrapper) HGet(key string, field string) {
	cmd := rp.pipe.HGet(context.TODO(), key, field)
	rp.cmds = append(rp.cmds, cmd)
}
