package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

var placeholder = "generating"

type BlobStore struct {
	TTL      time.Duration
	LockTime time.Duration
	Client   *redis.Client
}

func (s *BlobStore) WLock(name string) (bool, error) {
	return s.Client.SetNX(name, placeholder, s.LockTime).Result()
}

func (s *BlobStore) Write(name string, blob []byte) error {
	return s.Client.Set(name, blob, s.TTL).Err()
}

func (s *BlobStore) Read(name string) ([]byte, error) {
	blob, err := s.Client.Get(name).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "Get")
	} else if blob == placeholder {
		return nil, nil
	}
	return []byte(blob), nil
}
