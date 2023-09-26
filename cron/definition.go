package cron

import (
	"context"
	"sync"

	"github.com/go-co-op/gocron"
)

type Cron struct {
	scheduler *gocron.Scheduler
	mutex     sync.Mutex
	cancelCtx context.Context
}

var cron *Cron