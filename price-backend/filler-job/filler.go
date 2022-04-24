package fillerjob

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type priceFiller struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	mx sync.Mutex
}

func NewPriceFiller() *priceFiller {
	return &priceFiller{}
}

func (s *priceFiller) StartListening(jobPeriod time.Duration) error {
	s.mx.Lock()
	s.ctx, s.cancelCtx = context.WithCancel(context.Background())
	s.mx.Unlock()
	for {
		select {
		case <-time.After(jobPeriod):
			s.startFilling()
		case <-s.ctx.Done():
			return fmt.Errorf("context canceled")
		}
	}
}

func (s *priceFiller) startFilling() {
	// I will get here all currency pairs which exists in table assets and not acquired
	// Will loop over currency pairs
	// Will lock every currency pair and in worker pool parallel get from api and insert into asset table
	// Will unlock after finish each currency pair insert
	// Lock I will do based on postgres current time
	// But I thinks redis is more comfortable and fast to do distributed locking
}

func (s *priceFiller) ShutDown() {
	s.cancelCtx()
}
