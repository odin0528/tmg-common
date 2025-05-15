package redis

import (
	"time"
)

func (rp *redisPipelineWrapper) Set(key string, value interface{}, expiration time.Duration) {
	cmd := rp.pipe.Set(rp.ctx, key, value, expiration)
	rp.cmds = append(rp.cmds, cmd)
}

func (rp *redisPipelineWrapper) Get(key string) {
	cmd := rp.pipe.Get(rp.ctx, key)
	rp.cmds = append(rp.cmds, cmd)
}
