package netcom

import "crypto/rsa"

func CreateOptions(encrypt bool, skey_size uint, sender *rsa.PrivateKey, receiver *rsa.PublicKey) *Options {
	return &Options{
		Encrypt:      encrypt,
		Skey_size:    skey_size,
		Rand_size:    1024,
		Key_sender:   sender,
		Key_receiver: receiver,
	}
}
