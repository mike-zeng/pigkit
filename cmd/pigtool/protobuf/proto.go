package protobuf

import "os/exec"

const  compiler = "protoc"

func Gen(idlFile string)  {
	exec.Command(compiler,
	"--proto_path="+idlFile,
	)
}