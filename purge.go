package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	configFileName = "purge.yml"
)

var (
	tokens    chan struct{}
	cfg       *Config
	startTime = time.Now().Round(time.Second)
	//go build -ldflags "-X main.buildtime '2015-12-22' -X main.version 'v1.0'"
	version   = "debug build"
	buildtime = "n/a"
)

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

func main() {
	var err error
	cfg, err = reloadConfig(configFileName)
	if err != nil {
		log.Fatalf("Can't open %s: %s", configFileName, err)
	}
	fmt.Printf("%+v\n", cfg)

	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	//defer os.RemoveAll(tmpDir)
	//os.Chdir(tmpDir)
	os.Chdir("/")

	//subDirToSkip := "skip"

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			//fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			//return filepath.SkipDir
			//}
			fmt.Printf("visited dir: %q with modtime %s\n", filepath.Base(path),info.ModTime())
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
}
