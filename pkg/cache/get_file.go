package cache

import (
	"os"
	"telegram-file-server/pkg/config"
)

const cacheFolder = ".cache"

func GetFile(f func(fileId string) (string, error)) func(fileId string) (string, error) {
	return func(fileId string) (string, error) {
		// Check if the file is already cached
		cachedFilePath, err := cacheStorage.Get(fileId)
		if err == nil {
			_, err := os.OpenFile(string(cachedFilePath), os.O_RDONLY, 0644)
			if err == nil {
				return string(cachedFilePath), nil
			} else {
				_ = cacheStorage.Delete(fileId)
			}
		}

		filePath, err := f(fileId)
		if err != nil {
			return "", err
		}
		// Move the file to the cache folder
		nonCachedTempFile, err := os.Open(filePath)
		if err != nil {
			return "", err
		}
		file, err := os.CreateTemp(cacheFolder, "file")
		if err != nil {
			return "", err
		}
		_, err = nonCachedTempFile.WriteTo(file)
		if err != nil {
			return "", err
		}
		defer func() {
			_ = nonCachedTempFile.Close()
			_ = file.Close()
			_ = os.RemoveAll(filePath)
		}()
		_ = cacheStorage.Set(fileId, []byte(file.Name()), config.CacheFileTTL)

		return file.Name(), nil
	}
}
