package event

import (
	"github.com/step_ddd/domain/parking"
	"github.com/step_ddd/log"
)

var MQ = &QueueImpl{}

type QueueImpl struct {
	q chan interface{}
}

func init() {
	MQ.q = make(chan interface{}, 1024)
	go listen(MQ)
}

func (e *QueueImpl) Enqueue(v interface{}) {
	e.q <- v
}

func listen(mq *QueueImpl) {
	for v := range mq.q {
		log.G.Infof("event listen: %#v", v)
		switch v.(type) {
		case *parking.CheckInEvent:
		case *parking.CheckInFailEvent:
		case *parking.CheckOutEvent:
		case *parking.CheckOutFailEvent:
		case *parking.PaidEvent:
		default:
			log.G.Errorf("Invaild Event type:%#v", v)
		}
	}
}
