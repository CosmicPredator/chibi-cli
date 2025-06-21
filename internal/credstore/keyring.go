package credstore

import "github.com/zalando/go-keyring"

var SRV_NAME = "chibi"

// add a key value pair to config table
func SetCredential(key string, value string) error {
	err := keyring.Set(SRV_NAME, key, value)
	return err
}

func GetCredential(key string) (*string, error) {
	value, err := keyring.Get(SRV_NAME, key)
	return &value, err
}

func DeleteCredentials() error {
	return keyring.DeleteAll(SRV_NAME)
}