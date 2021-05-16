package utils

import (
	"github.com/mike-zeng/pigkit/cmd/pigtool/options"
	"github.com/mike-zeng/pigkit/cmd/pigtool/parser"
	"strings"
)

func GetRootPath(opt options.Options) string  {
	if opt.OutDir== "" {
		opt.OutDir = "./"
	}
	return opt.OutDir + opt.ModuleName
}

func GetProtoFileName(idlPath string) string{
	split := strings.Split(idlPath,".")
	temp := split[len(split)-2]
	return strings.Replace(temp, "/", "", -1)
}

func GetPkgNameOfPb(opt options.Options,parser *parser.ProtoParser) string  {
	rootPath := GetRootPath(opt)
	return strings.Replace(rootPath, "./", "",1) + "/gen/"+parser.GetPackageName()
}