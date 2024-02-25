package cache

import (
	"os"
	"telegram-file-server/pkg/log"
)

func clearCache() {
	log.Info("Start cleaning up cache")
	keys, err := cacheStorage.Keys()
	if err != nil {
		log.Error("Clean cache folder error: %v", err)
		return
	}

	// Build a map for faster search
	cacheStorageMap := make(map[string]bool)
	for _, key := range keys {
		value, err := cacheStorage.Get(string(key))
		if err == nil {
			cacheStorageMap[string(value)] = true
		}
	}

	// Scan cache folder
	// Remove files is not in the cacheStorage
	files, err := os.ReadDir(cacheFolder)
	if err != nil {
		log.Error("Clean cache folder error: %v", err)
		return
	}
	for _, file := range files {
		if _, ok := cacheStorageMap[file.Name()]; !ok {
			_ = os.RemoveAll(file.Name())
		}
	}
	log.Info("Cron job is ended")
}
