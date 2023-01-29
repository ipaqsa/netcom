package packUtils

import (
	"io/ioutil"
)

func SaveFileFromByte(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0664)
	if err != nil {
		return err
	}
	return nil
}
