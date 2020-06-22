package main

import (
	"errors"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestMemcache(t *testing.T) {
	s := &MemcacheService{
		client: memcache.New("localhost:11211"),
	}

	key := "hoge"
	v := &SampleValue{
		Content:   "Hello World",
		Count:     100,
		CreatedAt: time.Now(),
	}
	// set
	if err := s.Set(key, v); err != nil {
		t.Fatal(err)
	}

	// get
	got, err := s.Get(key)
	if errors.Is(err, memcache.ErrCacheMiss) {
		t.Log("cache miss")
		return
	} else if err != nil {
		t.Fatal(err)
	}

	// check
	if e, g := v.Content, got.Content; e != g {
		t.Errorf("Content want %v but got %v", e, g)
	}
	if e, g := v.Count, got.Count; e != g {
		t.Errorf("Count want %v but got %v", e, g)
	}
	if v.CreatedAt.IsZero() {
		t.Errorf("CreatedAt is zero")
	}
}
