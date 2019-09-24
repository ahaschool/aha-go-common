package main

import (
	"bytes"
	"errors"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

// project project config
type project struct {
	Name       string
	Owner      string
	Path       string
	WithGRPC   bool
	Here       bool
	ModuleName string // 支持项目的自定义module名 （go.mod init）
}

const (
	_tplTypeGitgnore = iota
	_tplTypeGomod
	_tplTypeReadme
	_tplTypeApiCmd
	_tplTypeSrvCmd
	_tplTypeConf
	_tplTypeErrcode
	_tplTypeSrcApiExample
	_tplTypeSrcApiRouter
	_tplTypeSrcSrvExampleHandlerExample
	_tplTypeSrcSrvExampleHandlerRepository
	_tplTypeSrcSrvExampleModelExample
	_tplTypeSrcSrvExampleProtoExample
)

var (
	p project
	// files
	files = map[int]string{
		// init doc
		_tplTypeGitgnore:                         "/.gitignore",
		_tplTypeGomod:                            "/go.mod",
		_tplTypeReadme:                           "/README.md",
		_tplTypeApiCmd:                           "/cmd/api.go",
		_tplTypeSrvCmd:                           "/cmd/srv.go",
		_tplTypeConf:                             "/conf/conf.go",
		_tplTypeErrcode:                          "/conf/errcode.go",
		_tplTypeSrcApiExample:                    "/src/api/example.go",
		_tplTypeSrcApiRouter:                     "/src/api/router.go",
		_tplTypeSrcSrvExampleHandlerExample:      "/src/srv/example/handler/example.go",
		_tplTypeSrcSrvExampleHandlerRepository:   "/src/srv/example/handler/repository.go",
		_tplTypeSrcSrvExampleModelExample:        "/src/srv/example/model/example.go",
		_tplTypeSrcSrvExampleProtoExample:        "/src/srv/example/proto/example.proto",
	}

	tpls = map[int]string{
		_tplTypeGitgnore:                         _tplGitgnore,
		_tplTypeGomod:                            _tplGomod,
		_tplTypeReadme:                           _tplReadme,
		_tplTypeApiCmd:                           _tplApiCmd,
		_tplTypeSrvCmd:                           _tplSrvCmd,
		_tplTypeConf:                             _tplConf,
		_tplTypeErrcode:                          _tplErrcode,
		_tplTypeSrcApiExample:                    _tplSrcApiExample,
		_tplTypeSrcApiRouter:                     _tplSrcApiRouter,
		_tplTypeSrcSrvExampleHandlerExample:      _tplSrcSrvExampleHandlerExample,
		_tplTypeSrcSrvExampleHandlerRepository:   _tplSrcSrvExampleHandlerRepository,
		_tplTypeSrcSrvExampleModelExample:        _tplSrcSrvExampleModelExample,
		_tplTypeSrcSrvExampleProtoExample:        _tplSrcSrvExampleProtoExample,
	}
)

func runNew(c *cli.Context) error {
	if len(c.Args()) == 0 {
		return errors.New("required project name")
	}
	p.Name = c.Args()[0]

	if p.ModuleName == "" {
		p.ModuleName = p.Name
	}

	if p.Path != "" {
		p.Path = path.Join(p.Path, p.Name)
	} else {
		pwd, _ := os.Getwd()
		p.Path = path.Join(pwd, p.Name)
	}

	if err := create(); err != nil {
		return err
	}



	return nil
}

func create() (err error) {
	if err = os.MkdirAll(p.Path, 0755); err != nil {
		return
	}
	for t, v := range files {
		i := strings.LastIndex(v, "/")
		if i > 0 {
			dir := v[:i]
			if err = os.MkdirAll(p.Path+dir, 0755); err != nil {
				return
			}
		}
		if err = write(p.Path+v, tpls[t]); err != nil {
			return
		}
	}

	return
}

//func genpb() error {
//	cmd := exec.Command("protoc", "--proto_path=${GOPATH}/src:.", "--go_out=.", "--micro_out=.", p.Path+"/src/srv/example/proto/example.proto")
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	return cmd.Run()
//}

func write(name, tpl string) (err error) {
	data, err := parse(tpl)
	if err != nil {
		return
	}
	return ioutil.WriteFile(name, data, 0644)
}

func parse(s string) ([]byte, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
