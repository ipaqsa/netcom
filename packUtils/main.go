package packUtils

import (
	"crypto/rsa"
	"errors"
	"github.com/ipaqsa/netcom/cryptoUtils"
	"strings"
	"time"
)

func CreatePack(title string, data string) *Package {
	return &Package{
		Head: HeadPackage{
			Title:   title,
			Session: "SESSION",
			Sender:  "SENDER",
			Meta:    "META",
		},
		Body: BodyPackage{
			Date: time.Now().Format("2006-01-02 15:04:05"),
			Data: data,
			Hash: string(cryptoUtils.HashSum([]byte(data))),
		},
	}
}

func CreateEncryptPack(sender *rsa.PrivateKey, receiver *rsa.PublicKey, title, data string, skey_size, rand_size uint) ([]byte, error) {
	pack := CreatePack(title, data)
	pack.Encrypt(sender, receiver, skey_size, rand_size)
	if pack == nil {
		return nil, errors.New("encrypt fail")
	}
	jpack, err := pack.Marshal()
	if err != nil {
		return nil, err
	}
	return jpack, nil
}

func CreateFilePackage(path string) *Package {
	splited := strings.Split(path, "/")
	filename := splited[len(splited)-1]
	bytes, _ := cryptoUtils.GetFileBytes(path)
	return &Package{
		Head: HeadPackage{
			Title:   filename,
			Session: "SESSION",
			Sender:  "SENDER",
			Meta:    "META",
		},
		Body: BodyPackage{
			Date: time.Now().Format("2006-01-02 15:04:05"),
			Data: cryptoUtils.Base64Encode(bytes),
		},
	}
}
