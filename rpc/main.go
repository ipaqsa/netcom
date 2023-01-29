package rpc

import (
	"errors"
	"github.com/ipaqsa/netcom/cryptoUtils"
	"github.com/ipaqsa/netcom/packUtils"
)

func ReadPack(data []byte, response *packUtils.Package, opt *Options) (*packUtils.Package, *packUtils.Package, error) {
	ans := packUtils.CreatePack("answer", "")
	pack := packUtils.Unmarshal(data)
	if pack == nil {
		return nil, nil, errors.New("unmarshal err")
	}
	if opt.encrypt {
		err := pack.Decrypt(opt.key_sender)
		if err != nil {
			return nil, nil, errors.New(err.Error())
		}
		sender := cryptoUtils.ParsePublic(pack.Head.Sender)
		if sender == nil {
			return nil, nil, errors.New("parse sender key fail")
		}
		opt.key_receiver = sender
	}
	return pack, ans, nil
}

func SendAnswer(answer *packUtils.Package, opt *Options, response *packUtils.Package) {
	if opt.encrypt {
		answer.Encrypt(opt.key_sender, opt.key_receiver, opt.skey_size, opt.rand_size)
	}
	*response = *answer
}
