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

func NewFileCache() *FileCache {
	cacheDir := filepath.Join(os.TempDir(), "finkit-cache")
	_ = os.MkdirAll(cacheDir, 0755)
	return &FileCache{dir: cacheDir}
}

func (c *FileCache) Get(key string, v any) error {

	path := filepath.Join(c.dir, key+".json")

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var wrapper struct {
		Value     json.RawMessage `json:"value"`
		ExpiresAt time.Time       `json:"expires_at"`
	}

	if err := json.Unmarshal(b, &wrapper); err != nil {
		return err
	}

	if time.Now().After(wrapper.ExpiresAt) {
		return errors.New("cache expired")
	}

	return json.Unmarshal(wrapper.Value, v)
}

func (c *FileCache) Set(key string, v any, ttl time.Duration) error {

	path := filepath.Join(c.dir, key+".json")

	data := struct {
		Value     any       `json:"value"`
		ExpiresAt time.Time `json:"expires_at"`
	}{
		Value:     v,
		ExpiresAt: time.Now().Add(ttl),
	}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, 0644)
}
