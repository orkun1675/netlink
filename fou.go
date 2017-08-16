package netlink

import (
	"fmt"
)

// Rule represents a netlink rule.
type Fou struct {
	Port    uint16
	IpProto uint8
	Gue     bool
	Ipv6    bool
}

func (r Fou) String() string {
	return fmt.Sprintf("ip fou port: %d, ipProto: %d, gue: %v", r.Port, r.IpProto, r.Gue)
}

// NewFou return default Fou.
func NewFou() *Fou {
	return &Fou{
		Port:    0,
		IpProto: 0,
		Gue:     false,
		Ipv6:    false,
	}
}
