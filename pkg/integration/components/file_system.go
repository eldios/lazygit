package components

import (
	"fmt"
	"os"
)

type FileSystem struct {
	*assertionHelper
}

// This does _not_ check the files panel, it actually checks the filesystem
func (self *FileSystem) PathPresent(path string) *FileSystem {
	self.assertWithRetries(func() (bool, string) {
		_, err := os.Stat(path)
		return err == nil, fmt.Sprintf("Expected path '%s' to exist, but it does not", path)
	})
	return self
}

// This does _not_ check the files panel, it actually checks the filesystem
func (self *FileSystem) PathNotPresent(path string) *FileSystem {
	self.assertWithRetries(func() (bool, string) {
		_, err := os.Stat(path)
		return os.IsNotExist(err), fmt.Sprintf("Expected path '%s' to not exist, but it does", path)
	})
	return self
}

// Asserts that the file at the given path has the given content
func (self *FileSystem) FileContent(path string, matcher *TextMatcher) *FileSystem {
	self.assertWithRetries(func() (bool, string) {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			return false, fmt.Sprintf("Expected path '%s' to exist, but it does not", path)
		}

		output, err := os.ReadFile(path)
		if err != nil {
			return false, fmt.Sprintf("Expected error when reading file content at path '%s': %s", path, err.Error())
		}

		strOutput := string(output)

		if ok, errMsg := matcher.context("").test(strOutput); !ok {
			return false, fmt.Sprintf("Unexpected content in file %s: %s", path, errMsg)
		}

		return true, ""
	})
	return self
}
