package stun

import (
	"context"
	"github.com/pion/stun"
	"net"
	"sync"
)

func Dial(ctx context.Context, network, address string) (net.IP, error) {
	c, err := stun.Dial(network, address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	type Result struct {
		IP  net.IP
		Err error
	}
	resultChan := make(chan Result, 1)

	once := sync.Once{}
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Sending request to STUN server, waiting for response message.
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			once.Do(func() {
				resultChan <- Result{
					Err: res.Error,
				}
			})
			return
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			once.Do(func() {
				resultChan <- Result{
					Err: res.Error,
				}
			})
			return
		}
		once.Do(func() {
			resultChan <- Result{
				IP: xorAddr.IP,
			}
		})
	}); err != nil {
		return nil, err
	}
	select {
	case res := <-resultChan:
		return res.IP, res.Err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
