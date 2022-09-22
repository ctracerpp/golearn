package structdata

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"
)

func TestBallast(t *testing.T) {
	ballast := make([]byte, 10*1024*1024*1024) // 10G
	// <-time.After(time.Duration(math.MaxInt64))
	runtime.KeepAlive(ballast)

	http.ListenAndServe("localhost:6060", nil)

}
