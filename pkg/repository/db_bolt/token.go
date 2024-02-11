package db_bolt

import (
	"errors"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	"github.com/boltdb/bolt"
	"strconv"
)

type TokenRepo struct {
	db *bolt.DB
}

func NewTokenStorage(db *bolt.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

func (s *TokenRepo) Save(chatID int64, token string, bucket repository.Bucket) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

func (s *TokenRepo) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		token = string(b.Get(intToBytes(chatID)))
		return nil
	})

	if token == "" {
		return "", errors.New("not found")
	}

	return token, err
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
