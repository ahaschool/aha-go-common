package main

const (
	_tplGitgnore = `# idea ignore
bin/
.idea/
*.ipr
*.iml
*.iws
logs/*.log
!bin/run.sh
!bin/gateway.sh
`
	_tplGomod = `module {{.ModuleName}}

go 1.12

require (
	github.com/ahaschool/aha-go-common v0.2.3
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/joho/godotenv v0.0.0-20190204044109-5c0e6c6ab1a0
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/micro/go-micro v1.8.0
	github.com/nats-io/nats-server/v2 v2.0.4 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
	golang.org/x/sys v0.0.0-20190801041406-cbf593c0f2f3 // indirect
	google.golang.org/genproto v0.0.0-20190716160619-c506a9f90610 // indirect
	google.golang.org/grpc v1.22.1 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)`
	_tplReadme = `micro --registry=consul --registry_address=127.0.0.1:8500 api --namespace=cn.com.ahaschool.api --address=0.0.0.0:8081  --handler=http

protoc --proto_path=${GOPATH}/src:. --go_out=. --micro_out=. src/srv/*/proto/*.proto
`
	_tplApiCmd = `package main

import (
	"github.com/ahaschool/aha-go-common/log"

	"{{.ModuleName}}/conf"
	httpApi "{{.ModuleName}}/src/api"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
)

func main()  {

	// 初始化配置
	conf.Init()
	log.Init(conf.Conf.LOG)

	// 创建API服务
	api := web.NewService(
		web.Name(conf.Conf.ExampleSrv.ApiName),
		web.Registry(consul.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{
				conf.Conf.ExampleSrv.ConsulUrl,
			}
		})),
	)

	// 初始化服务
	if err := api.Init(); err != nil {
		log.Error("%s", err)
	}

	// 注入执行方法
	router := httpApi.InitRouter()
	api.Handle("/", router)

	// 启动服务
	if err := api.Run(); err != nil {
		log.Error("%s", err)
	}
}`
	_tplSrvCmd = `package main

import (
	"github.com/ahaschool/aha-go-common/cache/redis"
	"github.com/ahaschool/aha-go-common/database/sql"
	"github.com/ahaschool/aha-go-common/log"
	
	"{{.ModuleName}}/conf"
	exampleHd "{{.ModuleName}}/src/srv/example/handler"
	example "{{.ModuleName}}/src/srv/example/proto"

	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"time"
)

func main() {

	// 初始化配置
	conf.Init()
	log.Init(conf.Conf.LOG)
	redis.Init(conf.Conf.REDIS)

	// 创建RPC服务
	srv := micro.NewService(
		//指定服务注册中心，默认 consul
		micro.Registry(consul.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{
				conf.Conf.ExampleSrv.ConsulUrl,
			}
		})),
		//服务名称
		micro.Name(conf.Conf.ExampleSrv.ServerName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(logWrapper),
	)

	// 初始化服务
	srv.Init()
	db := sql.NewMySQL(conf.Conf.SQL)
	exampleHandler := exampleHd.NewExampleHandler(db)

	// 注入执行方法
	if err := example.RegisterExampleServiceHandler(srv.Server(), exampleHandler); err != nil {
		log.Error("%s", err)
	}

	// 启动服务
	if err := srv.Run(); err != nil {
		log.Error("%s", err)
	}
}

// logWrapper is a handler wrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		if err != nil {
			content := fmt.Sprintf("Rsp: %s; Method: %s; Body: %s; Header: %s;", err.Error(), req.Method(), req.Body(), req.Header())
			log.Error(content)
		}
		return err
	}
}`
	_tplConf = `package conf

import (
	"github.com/ahaschool/aha-go-common/cache/redis"
	"github.com/ahaschool/aha-go-common/database/sql"
	"github.com/ahaschool/aha-go-common/log"
	"github.com/ahaschool/aha-go-common/net/netutil/breaker"
	xtime "github.com/ahaschool/aha-go-common/time"

	"fmt"
	"github.com/joho/godotenv"
	"strconv"
	"time"
)

var (
	Conf = &Config{}
)

type Config struct {
	SQL         *sql.Config
	LOG         *log.Config
	RUN         *RUN
	AhaSev      *AhaSev
	REDIS       *redis.Config
	ExampleSrv *ExampleSrv
}

type ExampleSrv struct {
	ServerName string
	ApiName    string
	ConsulUrl  string
}

type RUN struct {
	Addr string
}

type AhaSev struct {
	XToken    string
	OrderHost string
	UserHost  string
}

type Xenv struct {
	UtmSource   string
	UtmMedium   string
	UtmCampaign string
}

func Init() {
	myEnv, _ := godotenv.Read()
	Conf.ExampleSrv = new(ExampleSrv)
	Conf.SQL = new(sql.Config)
	Conf.LOG = new(log.Config)
	Conf.RUN = new(RUN)
	Conf.AhaSev = new(AhaSev)
	Conf.REDIS = new(redis.Config)
	Conf.RUN.Addr = myEnv["RUN_ADDR"]
	Conf.AhaSev.XToken = myEnv["AHA_X_TOKEN"]
	Conf.AhaSev.UserHost = myEnv["AHA_USER_HOST"]
	Conf.AhaSev.OrderHost = myEnv["AHA_ORDER_HOST"]
	DSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", myEnv["DB_USERNAME"], myEnv["DB_PASSWORD"], myEnv["DB_HOST"], myEnv["DB_PORT"], myEnv["DB_DATABASE"])
	bc := &breaker.Config{
		Window:  xtime.Duration(10 * time.Second),
		Sleep:   xtime.Duration(10 * time.Second),
		Bucket:  10, //TODO 待生成环境调整
		Ratio:   0.5,
		Request: 100,
	}
	c := &sql.Config{
		Addr:         "test", //追踪 后续使用
		DSN:          DSN,
		Active:       10,
		Idle:         5,
		IdleTimeout:  xtime.Duration(time.Minute),
		QueryTimeout: xtime.Duration(time.Minute),
		ExecTimeout:  xtime.Duration(time.Minute),
		TranTimeout:  xtime.Duration(time.Minute),
		Breaker:      bc,
	}
	Conf.SQL = c
	Conf.LOG.Dir = myEnv["LOG_PATH"]
	Conf.REDIS.Name = myEnv["REDIS_NAME"]
	Conf.REDIS.Proto = myEnv["REDIS_PROTO"]
	Conf.REDIS.Addr = myEnv["REDIS_ADDR"]
	Conf.REDIS.Auth = myEnv["REDIS_AUTH"]
	if db, err := strconv.Atoi(myEnv["REDIS_DB"]); err != nil {
		content := fmt.Sprintf("redis conf err: %s", err.Error())
		fmt.Println(content)
	} else {
		Conf.REDIS.DB = db
	}
	Conf.ExampleSrv.ConsulUrl = myEnv["CONSUL_URL"]
	Conf.ExampleSrv.ServerName = myEnv["SERVER_NAME"]
	Conf.ExampleSrv.ApiName = myEnv["API_NAME"]
	return
}`
	_tplErrcode = `package conf

import "github.com/ahaschool/aha-go-common/errcode"

// 兑换码错误
var (
	CodeError         = errcode.ServerAdd(-100001, "兑换码错误")
	BatchError        = errcode.ServerAdd(-100002, "批次编号错误")
	BatchNotStart     = errcode.ServerAdd(-100003, "批次还未开始")
	BatchNotLeft      = errcode.ServerAdd(-100004, "兑换码发放完毕")
	CodeExist         = errcode.ServerAdd(-100005, "兑换码已经存在")
	CodeCreateError   = errcode.ServerAdd(-100006, "兑换码创建失败")
	UserIdError       = errcode.ServerAdd(-100007, "用户编号错误")
	CodeFail          = errcode.ServerAdd(-100008, "兑换码不可用")
	CodeRepeat        = errcode.ServerAdd(-100009, "兑换码重复使用")
	CodeTimeError     = errcode.ServerAdd(-1000010, "兑换码已过期或未开始")
	CodeValidateError = errcode.ServerAdd(-1000011, "兑换失败")
	CodeBindError     = errcode.ServerAdd(-1000012, "兑换码绑定失败")
)`
	_tplSrcApiExample = `package api

import (
	"github.com/ahaschool/aha-go-common/errcode"
	"github.com/ahaschool/aha-go-common/http"

	"{{.ModuleName}}/conf"
	example "{{.ModuleName}}/src/srv/example/proto"

	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
)

/**
 * ==================
 *   兑换码的HTTP服务
 * ==================
 */
type ExampleApi struct {
	client example.ExampleService
}

func NewExampleApi() *ExampleApi {
	cl := example.NewExampleService(conf.Conf.ExampleSrv.ServerName, client.DefaultClient)
	return &ExampleApi{client: cl}
}

/**
 * ==================
 *     测试接口
 * ==================
 */
func (a *ExampleApi) GetExampleInfo(ctx *http.Context) {
	exampleName := ctx.Gin.Query("name")
	if exampleName == "" {
		ctx.Error(errcode.ParamsError)
		return
	}

	name, err := a.client.ExampleInfo(context.TODO(), &example.ExampleInfoQuery{Name: exampleName})
	if err != nil {
		ctx.Error(errors.Parse(err.Error()))
		return
	}
	ctx.Response(name)
}`
	_tplSrcApiRouter = `package api

import (
	"github.com/ahaschool/aha-go-common/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() (engine *gin.Engine) {
	engine = gin.New()

	// 兑换码接口路由
	example := NewExampleApi()
	http.NewService(engine, "get", "/operation/example/info", example.GetExampleInfo)
	return
}`
	_tplSrcSrvExampleHandlerExample = `package handler

import (
	"github.com/ahaschool/aha-go-common/database/sql"
	"github.com/ahaschool/aha-go-common/log"

	example "{{.ModuleName}}/src/srv/example/proto"

	"context"
)

type ExampleSrv struct {
	Repo Repository
}

func NewExampleHandler(db *sql.DB) *ExampleSrv {
	repo := &ExampleRepository{Db: db}
	return &ExampleSrv{Repo: repo}
}

// 心跳
func (s *ExampleSrv) Ping(ctx context.Context, req *example.Request, rsp *example.Pong) error {
	rsp.Msg = "pong"
	return nil
}

// ExampleInfo 详情
func (s *ExampleSrv) ExampleInfo(ctx context.Context, req *example.ExampleInfoQuery, rsp *example.ResponseExampleInfo) error {
	data, err := s.Repo.ExampleInfo(ctx, req)
	if err != nil {
		log.Error("Get Name err (%v)", err)
		return nil
	}
	rsp.Name = &example.ExampleData{
		ID:              data.Id,
		Name:            data.Name,
		CreatedAt:       data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return nil
}`
	_tplSrcSrvExampleHandlerRepository = `package handler

import (
	"github.com/ahaschool/aha-go-common/database/sql"

	"{{.ModuleName}}/src/srv/example/model"
	example "{{.ModuleName}}/src/srv/example/proto"

	"context"
)

// Repository 数据库操作
type Repository interface {
	ExampleInfo(context.Context, *example.ExampleInfoQuery) (model.ExampleData, error)
}

// ExchangeRepository 兑换码
type ExampleRepository struct {
	Db *sql.DB
}

// ExampleInfo 获取详情
func (repo *ExampleRepository) ExampleInfo(ctx context.Context, input *example.ExampleInfoQuery) (model.ExampleData, error) {
	data := model.ExampleData{
		Name: "test name",
	}
	return data, nil
}`
	_tplSrcSrvExampleModelExample = `package model

import (
	"time"
)

type ExampleData struct {
	Id              int32
	Name            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}`
	_tplSrcSrvExampleProtoExample = `syntax = "proto3";


package cn.com.ahaschool.srv.operation.example;//包名 倒域名

// 服务接口
service ExampleService {
    // 定义心跳服务接口
    rpc Ping(Request) returns (Pong) {}
    // 定义测试数据接口
    rpc ExampleInfo(ExampleInfoQuery) returns (ResponseExampleInfo) {}
}

message Request {
}
message Response {
}
// 兑换码详情
message ExampleInfoQuery {
    string Name = 1;
}

// 数据
message ExampleData {
    int32   ID = 1;
    string  Name = 3;
    string  CreatedAt = 16;
    string  UpdatedAt = 17;
}

// 响应数据模型
message ResponseExampleInfo {
    ExampleData Name = 1;
}

//返回参数格式
message Pong {
    string msg = 1;
}`
)
