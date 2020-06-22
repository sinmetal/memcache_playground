package main

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type SampleValue struct {
	Content   string
	Count     int64
	CreatedAt time.Time
}

type MemcacheService struct {
	client *memcache.Client
}

func (s *MemcacheService) Set(key string, value *SampleValue) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(value); err != nil {
		return err
	}
	item := &memcache.Item{
		Key:        key,
		Value:      buf.Bytes(),
		Expiration: int32(time.Now().Add(24 * time.Hour).Second()),
	}
	if err := s.client.Set(item); err != nil {
		return err
	}
	return nil
}

func (s *MemcacheService) Get(key string) (*SampleValue, error) {
	item, err := s.client.Get(key)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(item.Value)
	dec := gob.NewDecoder(buf)
	var v *SampleValue
	err = dec.Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
