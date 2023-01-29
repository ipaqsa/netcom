package packUtils

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"netcom/cryptoUtils"
)

func (pack *Package) Marshal() ([]byte, error) {
	jpack, err := json.Marshal(pack)
	if err != nil {
		return nil, err
	}
	return jpack, nil
}

func Unmarshal(data []byte) *Package {
	var pack Package
	err := json.Unmarshal(data, &pack)
	if err != nil {
		return nil
	}
	return &pack
}

func (pack *Package) Encrypt(sender *rsa.PrivateKey, receiver *rsa.PublicKey, skeySize, randSize uint) {
	var (
		session = cryptoUtils.GenerateBytes(skeySize)
		rand    = cryptoUtils.GenerateBytes(randSize)
		hash    = cryptoUtils.HashSum(bytes.Join(
			[][]byte{
				rand,
				cryptoUtils.Base64Decode(cryptoUtils.StringPublic(&sender.PublicKey)),
				cryptoUtils.Base64Decode(cryptoUtils.StringPublic(receiver)),
				[]byte(pack.Head.Title),
				[]byte(pack.Body.Data),
				[]byte(pack.Body.Date),
			},
			[]byte{},
		))
		sign = cryptoUtils.Sign(sender, hash)
	)
	*pack = Package{
		Head: HeadPackage{
			Rand:    cryptoUtils.Base64Encode(cryptoUtils.EncryptAES(session, rand)),
			Title:   cryptoUtils.Base64Encode(cryptoUtils.EncryptAES(session, []byte(pack.Head.Title))),
			Sender:  cryptoUtils.Base64Encode(cryptoUtils.EncryptAES(session, cryptoUtils.Base64Decode(cryptoUtils.StringPublic(&sender.PublicKey)))),
			Session: cryptoUtils.Base64Encode(cryptoUtils.EncryptRSA(receiver, session)),
			Meta:    cryptoUtils.Base64Encode(cryptoUtils.EncryptAES(session, []byte(pack.Head.Meta))),
		},
		Body: BodyPackage{
			Date: cryptoUtils.Base64Encode(cryptoUtils.EncryptAES(session, []byte(pack.Body.Date))),
			Data: cryptoUtils.Base64Encode(cryptoUtils.EncryptAES(session, []byte(pack.Body.Data))),
			Hash: cryptoUtils.Base64Encode(hash),
			Sign: cryptoUtils.Base64Encode(sign),
		}}
}

func (pack *Package) Decrypt(priv *rsa.PrivateKey) error {
	session := cryptoUtils.DecryptRSA(priv, cryptoUtils.Base64Decode(pack.Head.Session))
	if session == nil {
		return errors.New("session decrypt fail")
	}
	publicBytes := cryptoUtils.DecryptAES(session, cryptoUtils.Base64Decode(pack.Head.Sender))
	if publicBytes == nil {
		return errors.New("sender decrypt fail")
	}
	public := cryptoUtils.ParsePublic(cryptoUtils.Base64Encode(publicBytes))
	if public == nil {
		return errors.New("parse sender fail")
	}
	titleBytes := cryptoUtils.DecryptAES(session, cryptoUtils.Base64Decode(pack.Head.Title))
	if titleBytes == nil {
		return errors.New("title decrypt fail")
	}
	metaBytes := cryptoUtils.DecryptAES(session, cryptoUtils.Base64Decode(pack.Head.Meta))
	if metaBytes == nil {
		return errors.New("meta decrypt fail")
	}
	dataBytes := cryptoUtils.DecryptAES(session, cryptoUtils.Base64Decode(pack.Body.Data))
	if dataBytes == nil {
		return errors.New("data decrypt fail")
	}
	dateBytes := cryptoUtils.DecryptAES(session, cryptoUtils.Base64Decode(pack.Body.Date))
	if dateBytes == nil {
		return errors.New("date decrypt fail")
	}
	rand := cryptoUtils.DecryptAES(session, cryptoUtils.Base64Decode(pack.Head.Rand))
	hash := cryptoUtils.HashSum(bytes.Join(
		[][]byte{
			rand,
			publicBytes,
			cryptoUtils.Base64Decode(cryptoUtils.StringPublic(&priv.PublicKey)),
			titleBytes,
			dataBytes,
			dateBytes,
		},
		[]byte{},
	))
	if cryptoUtils.Base64Encode(hash) != pack.Body.Hash {
		return errors.New("wrong hash")
	}
	err := cryptoUtils.Verify(public, hash, cryptoUtils.Base64Decode(pack.Body.Sign))
	if err != nil {
		return errors.New("wrong sign")
	}
	*pack = Package{
		Head: HeadPackage{
			Rand:    cryptoUtils.Base64Encode(rand),
			Title:   string(titleBytes),
			Sender:  cryptoUtils.Base64Encode(publicBytes),
			Session: cryptoUtils.Base64Encode(session),
			Meta:    string(metaBytes),
		},
		Body: BodyPackage{
			Date: string(dateBytes),
			Data: string(dataBytes),
			Hash: pack.Body.Hash,
			Sign: pack.Body.Sign,
		},
	}
	return nil
}
