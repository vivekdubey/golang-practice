package fileexplorer

import (
	"fmt"
	"os"
)

func CreateFile(f string) (err error) {
	info, err := os.Stat(f)
	if err != nil {
		if os.IsNotExist(err) {
			if _, err = os.Create(f); err != nil {
				return err
			}
		}
	} else {
		if info.IsDir() {
			return fmt.Errorf("Dir exists with the file name: %v", f)
		} else {
			return fmt.Errorf("File exists with the file name: %v", f)
		}
	}
	return nil
}
