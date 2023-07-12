package pathdigest

import (
	"crypto/md5"
	"io/fs"
	"os"
	"path/filepath"
)

func MD5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return m, nil
}
