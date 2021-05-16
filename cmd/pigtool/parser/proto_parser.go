package parser

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"log"
)

type ProtoParser struct {
	desc *desc.FileDescriptor
}

func NewProtoParser(idlPath string) *ProtoParser {
	Parser := protoparse.Parser{}
	//加载并解析 proto文件,得到一组 FileDescriptor
	descList, err := Parser.ParseFiles(idlPath)
	if err != nil {
		log.Fatalln(err)
	}
	if len(descList) == 0 {
		log.Fatalln("parser proto file error")
	}
	return &ProtoParser{
		desc: descList[0],
	}
}

func (parser *ProtoParser) GetServiceName()string {
	services := parser.desc.GetServices()
	if len(services) == 0 {
		log.Fatalln("proto file do not have service desc")
	}
	serDescriptor := services[0]
	return serDescriptor.GetName()
}

func (parser *ProtoParser) GetPackageName()string {
	goPackage := parser.desc.GetFileOptions().GetGoPackage()
	if goPackage == "" {
		panic("must assign go_package option in proto file")
	}
	return goPackage
}

func (parser ProtoParser) GetMethodList() []*desc.MethodDescriptor{
	svrDescriptor := parser.desc.GetServices()[0]
	return svrDescriptor.GetMethods()
}