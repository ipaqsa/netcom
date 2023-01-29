package netcom

import (
	"errors"
	"net/rpc"
	"netcom/packUtils"
)

func Send(addr, command string, pack *packUtils.Package, options *Options) (*packUtils.Package, error) {
	if pack == nil {
		return nil, errors.New("pack is nil")
	}
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	if options != nil {
		if options.encrypt {
			pack.Encrypt(options.key_sender, options.key_receiver, options.skey_size, options.rand_size)
		}
	}
	jpack, err := pack.Marshal()
	if err != nil {
		return nil, err
	}
	var response packUtils.Package
	err = client.Call(command, jpack, &response)
	if err != nil {
		return nil, err
	}
	if options != nil {
		if options.encrypt {
			err = response.Decrypt(options.key_sender)
			if err != nil {
				return nil, err
			}
			return &response, nil
		}
	}
	return &response, nil
}
