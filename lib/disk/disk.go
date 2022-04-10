package disk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err

		}
		if !info.IsDir() {
			size += info.Size()

		}
		return err

	})
	return size, err

}

func GetUserList() ([]string, error) {
	var ul []string

	files, err := ioutil.ReadDir("/data/mail")
	if err != nil {
		return ul, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			ul = append(paths, file.Name())
		}
	}

	return ul, nil

}

func Do() int {
	fmt.Println(GetUserList())
	fmt.Println(DirSize("/usr/bin"))
	return 0
}
