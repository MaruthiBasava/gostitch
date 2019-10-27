package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// StitchConf wraps around a map of filenames of each FileConf struct
type StitchConf struct {
	Files map[string]FileConf `yaml:"stitch_files"`
}

// FileConf holds information for each directory that needs a stitched file
type FileConf struct {
	Extension string   `yaml:"extension"`
	Directory string   `yaml:"directory"`
	Yield     string   `yaml:"yield"`
	Exclude   []string `yaml:"exclude"`
}

// StitchInit is called every update command on the cli
func StitchInit() error {

	conf := StitchConf{}
	filename := "stitchconf.yml"

	if _, err := os.Stat("./" + filename); os.IsNotExist(err) {
		if _, err := os.Create("./" + filename); err != nil {
			return err
		}
		return nil
	}

	ctnts, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(ctnts), &conf)
	if err != nil {
		return err
	}

	for k := range conf.Files {
		if err := yieldStitchedFile(conf.Files[k], k); err != nil {
			return err
		}
	}

	return nil
}

// YieldStitchedFile stitches each file from the iterated directory
func yieldStitchedFile(fileConf FileConf, filename string) error {

	cp := fileCompletePath(fileConf.Yield, fileConf.Extension, filename)
	if _, err := os.Stat(cp); err == nil {
		if err := os.Remove(cp); err != nil {
			return err
		}
	}

	files, err := ioutil.ReadDir("./" + fileConf.Directory)
	if err != nil {
		return err
	}

	fmt.Println(cp)

	err = ioutil.WriteFile(cp, stitchedFileHeader(), 0644)
	if err != nil {
		return err
	}

	for f := range filterFiles(fileConf.Exclude, files, fileConf.Extension) {
		fp := fileCompletePath(fileConf.Directory, "", f)
		ctnts, err := ioutil.ReadFile(fp)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(cp, fileContent(f, string(ctnts)), 0644)
		if err != nil {
			return err
		}

	}

	return nil
}

// FileCompletePath returns complete path of the file
func fileCompletePath(path string, ext string, filename string) string {
	return fmt.Sprintf("./%s/%s%s", path, filename, ext)
}

// StitchedFileHeader returns stitched file header
func stitchedFileHeader() []byte {
	return []byte("/* GENERATED FILE DO NOT EDIT */ \n")
}

// Formats file content
func fileContent(filename string, content string) []byte {
	return []byte(fmt.Sprintf("-- %s\n%s", filename, content))
}

// filters files
func filterFiles(exclude []string, files []os.FileInfo, ext string) map[string]int {
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
