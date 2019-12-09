package config

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// TemplateDir 模板文件夹
var TemplateDir = "templates"

// TemplatesFiles 模板文件
var TemplatesFiles []string

// CommonTemplatesFiles 通用模板
var CommonTemplatesFiles []string

func init() {
	lookupFiles()
}

func lookupFiles() {
	files, err := ioutil.ReadDir(TemplateDir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fileDirName := file.Name()
			if fileDirName == "public" {
				newDirFiles, err := ioutil.ReadDir(TemplateDir + "/" + fileDirName)
				if err != nil {
					fmt.Println(err)
					return
				}

				for _, newFile := range newDirFiles {
					newFilename := newFile.Name()
					if strings.HasSuffix(newFilename, ".html") {
						CommonTemplatesFiles = append(CommonTemplatesFiles, TemplateDir+"/"+fileDirName+"/"+newFilename)
					}
				}
			}

		} else {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				TemplatesFiles = append(TemplatesFiles, TemplateDir+"/"+filename)
			}
		}

	}
}
