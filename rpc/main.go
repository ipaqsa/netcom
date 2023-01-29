package netcom

import (
	"errors"
	"netcom/packUtils"
)

func ReadPack(data []byte, response *packUtils.Package, opt Options) (*packUtils.Package, *packUtils.Package, error) {
	ans := packUtils.CreatePack("answer", "")
	pack := packUtils.Unmarshal(data)
	if pack == nil {
		ans.Body.Data = "serialise err"
		*response = *ans
		return nil, nil, errors.New(ans.Body.Data)
	}
	if opt.encrypt {
		err := pack.Decrypt(opt.key_sender)
		if err != nil {
			ans.Body.Data = err.Error()
			*response = *ans
			return nil, nil, errors.New(ans.Body.Data)
		}
	}
	return pack, ans, nil
}

func SendAnswer(answer *packUtils.Package, opt *Options, response *packUtils.Package) {
	if opt.encrypt {
		answer.Encrypt(opt.key_sender, opt.key_receiver, opt.skey_size, opt.rand_size)
	}
	*response = *answer
}
