package lib

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/fs"
	"io/ioutil"
	"os"
)

func CopyFile(sourcePath, dstPath string, perm fs.FileMode) error {
	input, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}

	err = ioutil.WriteFile(dstPath, input, perm)
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	err = os.Chmod(dstPath, 0777)
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	return nil
}

func CopyFiles(sourceFilenameList []string, sourceDirPath, dstDirPath string, perm fs.FileMode) error {
	for _, filename := range sourceFilenameList {
		sourcePath := fmt.Sprintf("%s/%s", sourceDirPath, filename)
		dstPath := fmt.Sprintf("%s/%s", dstDirPath, filename)
		err := CopyFile(sourcePath, dstPath, perm)
		if err != nil {
			return err
		}
	}
	return nil
}
