package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// TemplateDir 模板文件夹
var TemplateDir = "templates"

// TemplatesFiles 模板文件
var TemplatesFiles []string

// CommonTemplatesFiles 通用模板
var CommonTemplatesFiles []string

func init() {
	log.Println("init template")
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
			} else {
				findTemplate(fileDirName, TemplateDir)
			}

		} else {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				TemplatesFiles = append(TemplatesFiles, TemplateDir+"/"+filename)
			}
		}

	}
}

func findTemplate(fileDirName string, baseDir string) {
	newDirFiles, err := ioutil.ReadDir(baseDir + "/" + fileDirName)
	if err != nil {
		log.Println(err)
		log.Println("Empty Dir:" + baseDir + "/" + fileDirName)
	}

	for _, newFile := range newDirFiles {
		newFilename := newFile.Name()
		if newFile.IsDir() {
			findTemplate(newFilename, baseDir+"/"+fileDirName)
		} else if strings.HasSuffix(newFilename, ".html") {
			TemplatesFiles = append(TemplatesFiles, baseDir+"/"+fileDirName+"/"+newFilename)
		}
	}
}
