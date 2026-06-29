package cache

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type Cache interface {
	Get(key string, v any) error
	Set(key string, v any, ttl time.Duration) error
}

type FileCache struct {
	dir string
	log *slog.Logger
}

func NewFileCache(log *slog.Logger) *FileCache {
	cacheDir := filepath.Join(os.TempDir(), "finkit-cache")
	_ = os.MkdirAll(cacheDir, 0755)
	return &FileCache{dir: cacheDir, log: log}
}

func (c *FileCache) Get(key string, v any) error {

	path := filepath.Join(c.dir, key+".json")

	b, err := os.ReadFile(path)
	if err != nil {
		c.log.Error("error reading cache file", "err", err)
		return err
	}

	var wrapper WrapperGet

	if err := json.Unmarshal(b, &wrapper); err != nil {
		c.log.Error("error unmarshalling cache file", "err", err)
		return err
	}

	if time.Now().After(wrapper.ExpiresAt) {
		return errors.New("cache expired")
	}

	return json.Unmarshal(wrapper.Value, v)
}

func (c *FileCache) Set(key string, v any, ttl time.Duration) error {

	path := filepath.Join(c.dir, key+".json")

	data := WrapperSet{
		Value:     v,
		ExpiresAt: time.Now().Add(ttl),
	}

	b, err := json.Marshal(data)
	if err != nil {
		c.log.Error("error marshalling cache data", "err", err)
		return err
	}

	return os.WriteFile(path, b, 0644)
}
