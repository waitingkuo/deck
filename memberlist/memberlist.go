package memberlist

import (
	"time"
)

//FIXME should redesign
type Memberlist struct {
	UDPBindIP   string
	UDPBindPort int
	udpListener *net.UDPConn
	members     []*net.UDPConn
	memberMap   map[string]*net.UDPConn
}

func New() *Memberlist {

}

func (m *Memberlist) Run() {
	//FIXME we need a stop func?
	time.Sleep(time.Second * 5)
}

func (m *Memberlist) Join() {
}
