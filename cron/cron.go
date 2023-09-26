package cron

import (
	"context"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

func Init(timeZone *time.Location) context.CancelFunc {
	ctx, cancelFunc := context.WithCancel(context.Background())
	if timeZone == nil {
		timeZone = time.UTC
	}

	cron = &Cron{
		scheduler: gocron.NewScheduler(timeZone),
		mutex:     sync.Mutex{},
		cancelCtx: ctx,
	}

	cron.scheduler.SingletonModeAll()

	return cancelFunc
}

func AddJob(timeCron *Cron, jobFun interface{}){
	timeCron.scheduler.Do(jobFun)
}

func Every(interval interface{}) *Cron {
	cron.scheduler.Every(interval)
	return cron
}

func Inst() *Cron {
	return cron
}

func (c *Cron) Millisecond() *Cron {
	c.scheduler.Millisecond()
	return c
}

func (c *Cron) Second() *Cron {
	c.scheduler.Second()
	return c
}

func (c *Cron) Day() *Cron {
	c.scheduler.Day()
	return c
}

func (c *Cron) Days() *Cron {
	c.scheduler.Days()
	return c
}

func (c *Cron) Month() *Cron {
	c.scheduler.Month()
	return c
}

func (c *Cron) Cron(cronExpression string) *Cron {
	c.scheduler.Cron(cronExpression)
	return c
}

func (c *Cron) Serve() {
	go cron.scheduler.StartBlocking()
}

func (c *Cron) Stop() {
	cron.scheduler.Stop()
}
