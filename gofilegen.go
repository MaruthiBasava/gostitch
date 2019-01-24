package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type StitchConf struct {
	Files map[string]FileConf `yaml:"files"`
}

type FileConf struct {
	Extension string   `yaml:"extension"`
	Directory string   `yaml:"directory"`
	Yield     string   `yaml:"yield"`
	Exclude   []string `yaml:"exclude"`
}

func ParseStitchConf() {

	conf := StitchConf{}

	ctnts, err := ioutil.ReadFile("stitchconf.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(ctnts), &conf)
	if err != nil {
		panic(err)
	}

	for k := range conf.Files {
		YieldStitchedFile(conf.Files[k], k)
	}

}

func YieldStitchedFile(fileConf FileConf, filename string) {

	cp := FileCompletePath(fileConf.Yield, fileConf.Extension, filename)
	if _, err := os.Stat(cp); err == nil {
		if err := os.Remove(cp); err != nil {
			panic(err)
		}
	}

	files, err := ioutil.ReadDir("./" + fileConf.Directory)
	if err != nil {
		panic(err)
	}

	var stitchedString strings.Builder

	fmt.Fprint(&stitchedString, StitchedFileHeader())

	for f, _ := range FilterFiles(fileConf.Exclude, files, fileConf.Extension) {
		fmt.Print(f)
		cp := FileCompletePath(fileConf.Directory, "", f)
		ctnts, err := ioutil.ReadFile(cp)
		if err != nil {
			panic(err)
		}

		fmt.Fprint(&stitchedString, FileContent(f, string(ctnts[:])))
	}

	err = ioutil.WriteFile(cp, []byte(stitchedString.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func FileCompletePath(path string, ext string, filename string) string {
	return fmt.Sprintf("./%s/%s%s", path, filename, ext)
}

func StitchedFileHeader() string {
	return `/* GENERATED FILE DO NOT EDIT */` + lbrk() + lbrk()
}

func lbrk() string {
	return `
`
}

func FileContent(filename string, content string) string {
	return fmt.Sprintf(`/* %s */%s%s`, filename, lbrk(), content)
}

func FilterFiles(exclude []string, files []os.FileInfo, ext string) map[string]int {
	fileMap := make(map[string]int)

	for _, f := range files {
		fileMap[f.Name()] = 0
	}

	for _, ef := range exclude {
		if _, ok := fileMap[ef]; ok {
			delete(fileMap, ef)
		}
	}

	return fileMap
}
