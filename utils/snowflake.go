package utils

import (
	"math/rand"
	"seckill/conf"
	"sync"
	"time"
)

const sequenceMask = (16<<1 - 1)

var (
	lastTime    int64
	mutex       sync.Mutex
	curSequence int64
)

func GenID() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	curHostID := conf.HostID
	timeStamp := time.Now().UnixNano()/1000000 - conf.StartTimeStamp
	if timeStamp > lastTime {
		curSequence = 0
	}
	switch timeStamp < lastTime {
	case lastTime-timeStamp <= 1000:
		timeStamp = lastTime
	default:
		curHostID += rand.Int63()
	}
	if timeStamp == lastTime {
		curSequence = (curSequence + 1) % sequenceMask
		if curSequence == 0 {
			timeStamp += 1
		}
	}
	lastTime = timeStamp
	return curSequence + curHostID<<12 + timeStamp<<22
}
