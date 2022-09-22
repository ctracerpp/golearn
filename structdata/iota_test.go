package structdata

import "testing"

func TestIota(t *testing.T) {

}

type Priority int

const (
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1
	bit1, mask1
	_, _
	bit3, mask3
)

