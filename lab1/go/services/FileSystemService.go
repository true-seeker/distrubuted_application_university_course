package services

import (
	"os"
)

func WriteFile(data []byte) error {
	err := os.WriteFile(CredentialsFileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile() ([]byte, error) {
	data, err := os.ReadFile(CredentialsFileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DeleteFile() error {
	err := os.Remove(CredentialsFileName)
	if err != nil {
		return err
	}
	return nil
}
