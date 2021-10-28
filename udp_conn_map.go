package golwip

import (
	"sync"
)

var udpConns sync.Map

type udpConnId struct {
	src string
}
