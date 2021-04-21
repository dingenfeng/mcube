package nats_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mcube/bus/broker/nats"
	"github.com/infraboard/mcube/bus/event"
	"github.com/infraboard/mcube/logger/zap"
)

func TestPubSub(t *testing.T) {
	should := assert.New(t)
	b, err := nats.NewBroker(nats.NewDefaultConfig())
	should.NoError(err)

	log := zap.L().Named("Nats Bus")
	b.Debug(log)
	oe := &event.OperateEventData{
		Session:   "xxx",
		Account:   "test",
		IpAddress: "172.16.16.1",
	}
	sourceEvent, err := event.NewOperateEvent(oe)
	should.NoError(err)

	should.NoError(b.Connect())
	err = b.Sub("test", func(topic string, e *event.Event) error {
		should.Equal(sourceEvent.Id, e.Id)
		target := &event.OperateEventData{}
		err := e.ParseData(target)
		should.NoError(err)
		should.Equal(oe.IpAddress, target.IpAddress)
		log.Info(target)
		return nil
	})
	should.NoError(err)

	should.NoError(b.Pub("test", sourceEvent))
	should.NoError(b.Disconnect())
}

func init() {
	zap.DevelopmentSetup()
}
