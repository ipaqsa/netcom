package rpc

import "crypto/rsa"

func CreateOptions(encrypt bool, skey_size uint, sender *rsa.PrivateKey, receiver *rsa.PublicKey) *Options {
	return &Options{
		encrypt:      encrypt,
		skey_size:    skey_size,
		rand_size:    1024,
		key_sender:   sender,
		key_receiver: receiver,
	}
}
