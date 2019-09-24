package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strings"
)

type run struct {
	Port  string
	RegistryAddress  string
	Namespace  string
}

func runStartAction(c *cli.Context) error {
	base, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if b.Name == "" {
		return errors.New("need name option")
	}

	runDir := buildDir(base, "bin", 5)
	logDir := buildDir(base, "logs", 5)

	logName := fmt.Sprintf("%s/run.log", logDir)
	logg, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Print("打开或者创建运行日志日志文件失败")
	}

	defer logg.Close()

	for _, v := range obj {
		runName := fmt.Sprintf("%s/%s_%s", runDir, b.Name, v)
		pidName := fmt.Sprintf("%s/%s_%s.pid", runDir, b.Name, v)
		pidf, err := os.OpenFile(pidName, os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command(runName)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GIN_MODE=release")
		cmd.Dir = runDir
		cmd.Stdout = logg
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			panic(err)
		}

		pid := fmt.Sprintf("%d", cmd.Process.Pid)
		data := []byte(pid)
		data = append(data, '\n')
		pidf.Truncate(0)
		pidf.Write(data)
		fmt.Printf("start pid: %s", pid)
		pidf.Close()
	}

	fmt.Printf("pandora: %s\n", Version)
	fmt.Printf("%s start success.\n", b.Name)
	return nil
}

func runStopAction(c *cli.Context) error {
	base, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if b.Name == "" {
		return errors.New("need name option")
	}

	runDir := buildDir(base, "bin", 5)
	logDir := buildDir(base, "logs", 5)

	logName := fmt.Sprintf("%s/run.log", logDir)
	logg, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Print("打开或者创建运行日志日志文件失败")
	}

	defer logg.Close()

	for _, v := range obj {
		pidName := fmt.Sprintf("%s/%s_%s.pid", runDir, b.Name, v)
		pidf, err := os.OpenFile(pidName, os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			panic(err)
		}

		buf := bufio.NewReader(pidf)
		pid, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		pid = strings.TrimSpace(pid)
		args := []string{pid}
		cmd := exec.Command("kill", args...)
		cmd.Dir = runDir
		cmd.Stdout = logg
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			panic(err)
		}
		pidf.Close()
		if err := os.Remove(pidName); err != nil {
			panic(err)
		}
		fmt.Printf("stop pid: %s", pid)
	}

	fmt.Printf("pandora: %s\n", Version)
	fmt.Printf("%s stop success.\n", b.Name)
	return nil
}

func runGatewayAction(c *cli.Context) error {
	base, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if &r.Port == nil {
		panic("need port option")
	}
	if &r.RegistryAddress == nil {
		panic("need registry address option")
	}
	if &r.Namespace == nil {
		panic("need namespace option")
	}
	if len(c.Args()) < 1 {
		panic("invalid arguments")
	}

	runDir := buildDir(base, "bin", 5)
	logDir := buildDir(base, "logs", 5)

	logName := fmt.Sprintf("%s/run.log", logDir)
	logg, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Print("打开或者创建运行日志日志文件失败")
	}

	pidName := fmt.Sprintf("%s/gateway.pid", runDir)
	pidf, err := os.OpenFile(pidName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	}

	defer logg.Close()

	typeName := c.Args()[0]
	switch typeName {
	case "start":
		registryAddress := fmt.Sprintf("--registry_address=%s", r.RegistryAddress)
		namespace := fmt.Sprintf("--namespace=%s", r.Namespace)
		address := fmt.Sprintf("--address=0.0.0.0:%s", r.Port)
		args := []string{"--registry=consul", registryAddress, "api", namespace, address, "--handler=http"}
		fmt.Print("=========")
		fmt.Print(args)
		cmd := exec.Command("micro", args...)
		cmd.Dir = runDir
		cmd.Stdout = logg
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			panic(err)
		}
		pid := fmt.Sprintf("%d", cmd.Process.Pid)
		data := []byte(pid)
		data = append(data, '\n')
		pidf.Truncate(0)
		pidf.Write(data)
		fmt.Printf("gateway start success pid: %s\n", pid)
		break
	case "stop":
		buf := bufio.NewReader(pidf)
		pid, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		pid = strings.TrimSpace(pid)
		args := []string{"-15", pid}
		cmd := exec.Command("kill", args...)
		cmd.Dir = runDir
		cmd.Stdout = logg
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			panic(err)
		}
		pidf.Close()
		if err := os.Remove(pidName); err != nil {
			panic(err)
		}
		fmt.Printf("gateway stop success.\n")
		break
	default:
		return errors.New("arguments must be start/stop")
	}

	fmt.Printf("pandora: %s\n", Version)
	return nil
}
