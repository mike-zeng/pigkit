package utils

import (
	"github.com/mike-zeng/pigkit/cmd/pigtool/options"
	"os"
)

func CreateRootDirIfNotExist(opt options.Options)error {
	rootPath := GetRootPath(opt)
	stat, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(rootPath, 0766)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		err := os.MkdirAll(rootPath, 0766)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateGenDirIfNotExist(opt options.Options)error {
	genPath := GetRootPath(opt)+"/gen/"
	err := os.MkdirAll(genPath, 0766)
	if err != nil {
		return err
	}
	return nil
}

func CreateClientProxyDirIfNotExist(opt options.Options)error {
	clientPath := GetRootPath(opt) + "/gen/client/"
	stat, err := os.Stat(clientPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(clientPath, 0766)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		err := os.MkdirAll(clientPath, 0766)
		if err != nil {
			return err
		}
	}
	return nil
}

