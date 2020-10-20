package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)

	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

// CreateDirectory creates a directory tree. Returns true if could create it and false in any other case
func CreateDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			logrus.Panicf("Error while creating dir %s. %s.", dirName, err)
		}
		return true
	}

	if src.Mode().IsRegular() {
		logrus.Warning(dirName, "already exist as a file!")
		return false
	}

	return false
}

// RemoveDirectory removes a directory tree. Returns true if could remove it and false if any errors were found
func RemoveDirectory(dirName string) bool {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		logrus.Warningf("The dir %s doesn't exists.", dirName)
		return false
	}

	err = os.RemoveAll(dirName)
	if err != nil {
		logrus.Errorf("Couldn't remove directory %s. %s.", dirName, err)
	}

	return true
}

// FileExists check if the file with name as argument exists
func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// RemoveFile removes a file. Returns true if coult remove it and false if any errors were found
func RemoveFile(fileName string) bool {
	if FileExists(fileName) {
		logrus.Warningf("The file %s doesn't exists.")
		return false
	}

	err := os.Remove(fileName)
	if err != nil {
		logrus.Errorf("Couldn't remove file %s. %s.", fileName, err)
		return false
	}

	return true
}
