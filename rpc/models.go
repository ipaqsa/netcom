package rpc

import "crypto/rsa"

type Options struct {
	encrypt      bool
	skey_size    uint
	rand_size    uint
	key_sender   *rsa.PrivateKey
	key_receiver *rsa.PublicKey
}
