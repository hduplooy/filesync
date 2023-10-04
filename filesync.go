package filesync

import (
	"io/fs"
	"os"
	"strings"
)

type Config struct {
	SourcePath      string
	DestinationPath string
}

func SyncFolders(config *Config) error {
	if !strings.HasSuffix(config.SourcePath, string(os.PathSeparator)) {
		config.SourcePath += string(os.PathSeparator)
	}
	if !strings.HasSuffix(config.DestinationPath, string(os.PathSeparator)) {
		config.DestinationPath += string(os.PathSeparator)
	}

	return fs.WalkDir(os.DirFS(config.SourcePath), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || path == "." {
			return nil
		}
		srcInfo, _ := d.Info()
		destInfo, err := os.Lstat(config.DestinationPath + path)
		if d.IsDir() {
			if err != nil {
				// Make folder
				err := os.MkdirAll(config.DestinationPath+path, 0777)
				if err != nil {
					return err
				}
			}
		} else {
			if err != nil || destInfo.ModTime().Unix() < srcInfo.ModTime().Unix() {
				buf, err := os.ReadFile(config.SourcePath + path)
				if err != nil {
					return err
				}
				err = os.WriteFile(config.DestinationPath+path, buf, 0777)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	return nil
}
