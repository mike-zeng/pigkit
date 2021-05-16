package main

import (
	"fmt"
	"github.com/mike-zeng/pigkit/cmd/pigtool/options"
	"github.com/mike-zeng/pigkit/cmd/pigtool/parser"
	"github.com/mike-zeng/pigkit/cmd/pigtool/utils"
	"log"
)



func main() {
 	opt := options.Options{}
	err := opt.Parse()
	if err != nil {
		log.Fatalln(err)
	}
	if opt.Server {
		// 1. create project dir
		err := utils.CreateRootDirIfNotExist(opt)
		if err != nil {
			log.Fatalln(err)
		}
		protoParser := parser.NewProtoParser(opt.IdlPath)
		//2. create gen dir
		err = utils.CreateGenDirIfNotExist(opt)
		if err != nil {
			log.Fatalln(err)
		}

		// 3. gen server code
		result, err := utils.GenServerCode(opt)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(result)

		// 4. gen main code
		err = utils.CreateMainFile(opt,protoParser)
		if err != nil {
			log.Fatalln(err)
		}

		// 5. gen go.mod
		err = utils.CreateModFile(opt)
		if err != nil {
			log.Fatalln(err)
		}

		// 6. gen handler file
		err = utils.CreateHandlerFile(opt,protoParser)
		if err != nil {
			log.Fatalln(err)
		}

		// 7. gen configFile
		err = utils.CreateConfigFile(opt)
		if err != nil{
			log.Fatalln(err)
		}

	}else {
		// 1. create dir if not exist
		err := utils.CreateClientProxyDirIfNotExist(opt)
		if err != nil {
			log.Fatalln(err)
		}
		// 2. gen client proxy dir
		result, err := utils.GenClientProxyCode(opt)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(result)

	}
}


