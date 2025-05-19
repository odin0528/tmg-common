package redis

import (
	"context"
	"time"
)

func (tp *TxPipelineWrapper) Set(key string, value interface{}, expiration time.Duration) {
	cmd := tp.pipe.Set(context.TODO(), key, value, expiration)
	tp.cmds = append(tp.cmds, cmd)
}

func (tp *TxPipelineWrapper) Get(key string) {
	cmd := tp.pipe.Get(context.TODO(), key)
	tp.cmds = append(tp.cmds, cmd)
}

func (tp *TxPipelineWrapper) Exec(ctx context.Context) ([]CmdResult, error) {
	cmds, err := tp.pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return processPipelineResults(cmds), nil
}

func (tp *TxPipelineWrapper) Discard() error {
	tp.cmds = tp.cmds[:0]
	return tp.pipe.Discard()
}

func (tp *TxPipelineWrapper) HSet(key string, field string, value interface{}) {
	cmd := tp.pipe.HSet(context.TODO(), key, field, value)
	tp.cmds = append(tp.cmds, cmd)
}

func (tp *TxPipelineWrapper) HGet(key string, field string) {
	cmd := tp.pipe.HGet(context.TODO(), key, field)
	tp.cmds = append(tp.cmds, cmd)
}
