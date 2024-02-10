package bolt

import (
	"errors"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	"go.etcd.io/bbolt"
	"strconv"
)

type TokenRepository struct {
	db *bbolt.DB
}

func NewTokenRepository(db *bbolt.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) Save(chatID int64, token string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

func (s *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string

	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		token = string(b.Get(intToBytes(chatID)))
		return nil
	})

	if token == "" {
		return "", errors.New("not found")
	}

	return token, err
}

//func (r *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
//	var token string
//	err := r.db.View(func(tx *bbolt.Tx) error {
//		b := tx.Bucket([]byte(bucket))
//		data := b.Get(intToByte(chatID))
//		token = string(data)
//		return nil
//	})
//	if err != nil {
//		return "", errors.New("Not Found")
//	}
//
//	if token == "" {
//		return "", errors.New("Error in get method db, token not found")
//	}
//	return token, err
//}

func intToBytes(val int64) []byte {
	return []byte(strconv.FormatInt(val, 10))
}
