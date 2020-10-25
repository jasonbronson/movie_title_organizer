package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func CopyFile(src, des string, BUFFERSIZE int64) error {

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	_, err = os.Stat(des)
	if err == nil {
		return fmt.Errorf("File %s already exists", des)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(des)
	if err != nil {
		return err
	}
	defer destination.Close()

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}

	return err

}

func GetDirectoryFiles(directory string) []string {

	list := make([]string, 0)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		//fmt.Println(f.Name())
		list = append(list, f.Name())
	}
	return list
}
