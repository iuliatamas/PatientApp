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
	msgs := make(chan *pkg.SMSMessage)
	go func() {
		for {
			time.Sleep(1 * time.Second)

			mq.lock.Lock()
			e := mq.Front()
			if e == nil {
				mq.lock.Unlock()
				continue
			}
			m := mq.Remove(e).(*pkg.SMSMessage)
			mq.lock.Unlock()
			msgs <- m
		}
	}()

	return msgs
}

func NewMessageQueue() *MsgQueue {
	return &MsgQueue{list.New(), sync.Mutex{}}
}
