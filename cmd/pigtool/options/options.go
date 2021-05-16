package options

import (
	"errors"
	"flag"
)

type Options struct {
	Server bool
	Client bool
	ModuleName string
	ClientModuleName string
	IdlPath string
	OutDir string
}

func (opt *Options) Parse()error  {
	flag.BoolVar(&opt.Client,"c",false,"")
	flag.BoolVar(&opt.Server,"s",false,"")
	flag.StringVar(&opt.IdlPath,"i","","")
	flag.StringVar(&opt.ClientModuleName,"m","","")

	flag.Parse()
	if !opt.Server && !opt.Client {
		return errors.New("need -c or -s")
	}
	if opt.IdlPath == "" {
		return errors.New("must specify idl path")
	}
	if opt.Client && opt.ModuleName == "" {
		return errors.New("must specify idl module name")
	}
	return nil
}
