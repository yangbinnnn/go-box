package core

import (
	"go-box/common"
	"time"
)

func MSGLoop() {
	for {
		common.WSBroadcast([]byte(time.Now().Format(time.RFC3339)))
		time.Sleep(1 * time.Second)
	}
}
