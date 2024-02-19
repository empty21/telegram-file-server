package cache

import (
	"github.com/gofiber/storage/memory/v2"
	"os"
	"time"
)

var cacheStorage *memory.Storage

func init() {
	_ = os.MkdirAll(cacheFolder, os.ModeDir)
	cacheStorage = memory.New(memory.Config{
		GCInterval: 10 * time.Second,
	})
}
