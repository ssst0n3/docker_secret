package lib

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
)

func CopyFile(sourcePath, dstPath string) error {
	input, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}

	err = ioutil.WriteFile(dstPath, input, 0644)
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	return nil
}

func CopyFiles(sourceFilenameList []string, sourceDirPath, dstDirPath string) error {
	for _, filename := range sourceFilenameList {
		sourcePath := fmt.Sprintf("%s/%s", sourceDirPath, filename)
		dstPath := fmt.Sprintf("%s/%s", dstDirPath, filename)
		err := CopyFile(sourcePath, dstPath)
		if err != nil {
			return err
		}
	}
	return nil
}
