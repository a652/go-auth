package store

import "errors"

type CredentialStore interface {
	/* add methods */
	GetSecretByAppID(appID string) (string, error)
}

var _ CredentialStore = (*DefaultCredentialStore)(nil)

type DefaultCredentialStore struct {
	storage map[string]string
}

func NewDefaultCredentialStore(storage map[string]string) *DefaultCredentialStore {
	return &DefaultCredentialStore{storage}
}

func (dcs *DefaultCredentialStore) GetSecretByAppID(appID string) (string, error) {
	if secret, ok := dcs.storage[appID]; ok {
		return secret, nil
	}
	return "", errors.New("secret not found")
}
