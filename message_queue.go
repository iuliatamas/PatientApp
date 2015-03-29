package main

import (
	"container/list"
	"sync"
	"time"

	pkg "github.com/njern/gonexmo"
)

type MsgQueue struct {
	*list.List
	lock sync.Mutex
}

func (mq *MsgQueue) Push(msg *pkg.SMSMessage) {
	mq.lock.Lock()
	mq.PushBack(msg)
	mq.lock.Unlock()
}

func (mq *MsgQueue) Stream() <-chan *pkg.SMSMessage {
	time.Sleep(1 * time.Second)
	msgs := make(chan *pkg.SMSMessage)
	go func() {
		for {
			mq.lock.Lock()
			e := mq.Front()
			if e == nil {
				mq.lock.Unlock()
				time.Sleep(500 * time.Millisecond)
				continue
			}
			defer mq.lock.Unlock()
			msgs <- mq.Remove(e).(*pkg.SMSMessage)
		}
	}()

	return msgs
}

func NewMessageQueue() *MsgQueue {
	return &MsgQueue{list.New(), sync.Mutex{}}
}
