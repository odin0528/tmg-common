package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (rp *PipelineWrapper) Exec(ctx context.Context) ([]CmdResult, error) {
	cmds, err := rp.pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return processPipelineResults(cmds), nil
}

func (rp *PipelineWrapper) Set(key string, value interface{}, expiration time.Duration) {
	rp.pipe.Set(context.TODO(), key, value, expiration)
}

func (rp *PipelineWrapper) Get(key string) {
	rp.pipe.Get(context.TODO(), key)
}

func (rp *PipelineWrapper) HSet(key string, field string, value interface{}) {
	rp.pipe.HSet(context.TODO(), key, field, value)
}

func (rp *PipelineWrapper) HGet(key string, field string) {
	rp.pipe.HGet(context.TODO(), key, field)
}

func (rp *PipelineWrapper) HGetAll(key string) {
	rp.pipe.HGetAll(context.TODO(), key)
}

func (rp *PipelineWrapper) ZAdd(key string, members ...*redis.Z) {
	rp.pipe.ZAdd(context.TODO(), key, members...)
}

func (rp *PipelineWrapper) Publish(channel string, message interface{}) {
	rp.pipe.Publish(context.TODO(), channel, message)
}

func (rp *PipelineWrapper) Unlink(keys []string) ([]CmdResult, error) {
	ctx := context.Background()
	rp.pipe.Unlink(ctx, keys...)
	return rp.Exec(ctx)
}

func (rp *PipelineWrapper) Del(key string) {
	rp.pipe.Del(context.TODO(), key)
}

func (rp *PipelineWrapper) DelHashMap(key string, field []string) {
	rp.pipe.HDel(context.TODO(), key, field...)
}
