//go:build !linux

package device

import (
	"github.com/RunawayVPN/wireguard-go/conn"
	"github.com/RunawayVPN/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
