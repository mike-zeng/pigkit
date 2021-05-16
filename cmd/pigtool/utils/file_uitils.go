package utils

import (
	"bytes"
	"fmt"
	"github.com/mike-zeng/pigkit/cmd/pigtool/env"
	"github.com/mike-zeng/pigkit/cmd/pigtool/options"
	"github.com/mike-zeng/pigkit/cmd/pigtool/parser"
	"os"
	"os/exec"
)

type WriterWrap struct {
	file *os.File
	err error
}

func (receiver *WriterWrap) Write(content string)  {
	if receiver.err != nil {
		return
	}
	_, err := receiver.file.Write([]byte(content))
	if err != nil {
		receiver.err = err
	}
}

func (receiver WriterWrap) WriteWhiteLine()  {
	receiver.Write(fmt.Sprintf("\n"))
}
func (receiver *WriterWrap) Error()error {
	return receiver.err
}


func GenServerCode(opt options.Options) (string,error)  {
	outPath := GetRootPath(opt) + "/gen/"
	fmt.Println(outPath)
	pigRpcOutParam := fmt.Sprintf("--pigrpc_out=gen_obj=server,plugin:%s",outPath)
	cmd := exec.Command("protoc", opt.IdlPath,pigRpcOutParam)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "",err
	}
	return out.String(),nil
}

func CreateModFile(opt options.Options) error{
	rootPath := GetRootPath(opt)
	modFile, err := os.Create(rootPath + "/go.mod")
	if os.IsExist(err) {
		return err
	}

	writerWrap := WriterWrap{
		file: modFile,
	}

	writerWrap.Write(fmt.Sprintf("module %s",opt.ModuleName))
	writerWrap.WriteWhiteLine()
	writerWrap.WriteWhiteLine()
	writerWrap.Write(fmt.Sprintf("go %s\n",env.GetGoVersionNum()))
	writerWrap.Write(fmt.Sprintf(`
require (
	github.com/golang/protobuf v1.5.2
	github.com/mike-zeng/pigkit/rpc %s
)`,env.Version))
	if env.IsDebug() {
		writerWrap.WriteWhiteLine()
		writerWrap.Write("replace github.com/mike-zeng/pigkit/rpc v1.0.0 => ../../../../rpc")
		writerWrap.WriteWhiteLine()
	}
	return writerWrap.err
}

func CreateMainFile(opt options.Options,parser *parser.ProtoParser) error{
	rootPath := GetRootPath(opt)
	mainFile, err := os.Create(rootPath + "/main.go")
	serviceName := parser.GetServiceName()
	if os.IsExist(err) {
		return err
	}
	writerWrap := WriterWrap{
		file: mainFile,
	}
	pkgName := GetPkgNameOfPb(opt,parser)
	writerWrap.Write(fmt.Sprintf(`
package main

import (
	"github.com/mike-zeng/pigkit/rpc/server"
	desc "%s"
)

func main() {`,pkgName))

	writerWrap.Write(fmt.Sprintf(`
	pigServer := server.NewPigServer(desc.%sServiceDesc, %sHandler{}, "%s")
	options := &server.Options{
		Port: 8080,
	}`,serviceName,serviceName,"./pig.yaml"))
	writerWrap.Write(fmt.Sprintf(`
	pigServer.Serve(options)
`))
	writerWrap.Write(fmt.Sprintf("}"))
	return err
}



func CreateHandlerFile(opt options.Options,parser *parser.ProtoParser)error {
	rootPath := GetRootPath(opt)
	handlerFile, err := os.Create(rootPath + "/handler.go")
	if os.IsExist(err) {
		return err
	}
	writerWrap := WriterWrap{
		file: handlerFile,
	}
	// write package
	pkgName := GetPkgNameOfPb(opt,parser)
	writerWrap.Write(fmt.Sprintf(`
package main

import (
	"context"
	desc "%s"
)
`,pkgName))
	// write struct
	writerWrap.Write(fmt.Sprintf(`
type %sHandler struct{
	// handler struct
}`,parser.GetServiceName()))
	writerWrap.WriteWhiteLine()

	// parser methods
	methods := parser.GetMethodList()
	for _,method := range methods {
		writerWrap.Write(fmt.Sprintf(`
// %s pig tool auto gen
func (handler *%sHandler)%s(ctx context.Context, req *desc.%s)(*desc.%s, error) {
	// todo impl handler method
	return nil, nil
}`,method.GetName(),parser.GetServiceName(),method.GetName(),
method.GetInputType().GetName(),method.GetOutputType().GetName()))
	}
	return err
}

func GenClientProxyCode(opt options.Options) (string,error)  {
	outPath := GetRootPath(opt) + "/gen/client/"
	fmt.Println(outPath)
	pigRpcOutParam := fmt.Sprintf("--pigrpc_out=gen_obj=client,plugin:%s",outPath)
	cmd := exec.Command("protoc", opt.IdlPath,pigRpcOutParam)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "",err
	}
	return out.String(),nil
}

func CreateConfigFile(opt options.Options)error {
	rootPath := GetRootPath(opt)
	handlerFile, err := os.Create(rootPath + "/pig.yaml")
	if os.IsExist(err) {
		return err
	}
	writerWrap := WriterWrap{
		file: handlerFile,
	}
	writerWrap.Write(fmt.Sprintf(`server:
  port: 8080
  timeout: 2000
  etcd:
    timeout: 2000
    hosts:
      - 127.0.0.1:2379
`))
	return nil
}