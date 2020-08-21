package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/**
优雅重启:

1、Fork一个子进程，继承父进程的监听socket
2、子进程启动成功之后，接收新的连接
3、父进程停止接收新的连接，等待已有请求处理完成后退出



Q：子进程如何继承父进程的文件句柄？
A：通过os.Cmd对象中的ExtraFiles参数进行传递
注意在linux系统中才支持此特性

Q：如何优雅关闭？
A：使用go1.8版本新增的Shutdown方法进行优雅关闭；

Q：子进程如何接管父进程监听？
A：使用socket继承实现子进程接管父进程的监听socket

Q：如果进行信号处理?
A：Linux下可以通过`kill -信号 进程ID`来发送结束信号。

信号			值		说明
SIGHUP  	1       终端控制进程接收(终端连接断开)
SIGINT  	2		用户发送INTR字符(Ctrl+C)触发
SIGQUIT 	3		用户发送QUIT字符(Ctrl+/)触发
SIGKILL		9		无条件接收程序(不能被捕获、阻塞或忽略)
SIGUSR1		10		用户保留，自定义命令使用
SIGUSR2		12		用户保留，自定义命令使用
SIGPIPE		13		消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
SIGALRM		14		时钟定时信号
SIGTERM		15		结束程序(可以被捕获、阻塞或忽略)

进程接收到信号以后，可以使用os/Signal包进行处理

*/

var (
	child *bool
	wg    sync.WaitGroup
)

func processSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGKILL:
			//捕获不到SIGKILL，此处不会执行打印
			fmt.Printf("receive sigkill\n")
		case syscall.SIGINT:
			fmt.Printf("receive sigint\n")
		case syscall.SIGTERM:
			fmt.Printf("receive sigterm\n")
			return
		case syscall.SIGUSR2:
			// 这里可以实现reload
			fmt.Printf("receive siguser2\n")
			return
		}
	}
}

func startChild(file *os.File) {
	args := []string{"-child"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//put socket FD at the first entry
	cmd.ExtraFiles = []*os.File{file}
	err := cmd.Start()
	if err != nil {
		fmt.Printf("start child failed, err:%v\n", err)
		return
	}
	wg.Done()
}

func init() {
	child = flag.Bool("child", false, "继承于父进程(internal use only)")
	flag.Parse()
}

func readFromParent() {
	//fd = 0: 标准输出
	//fd = 1: 标准输入
	//fd = 2: 标准错误输出
	//fd = 3 ====> ExtraFiles[0]
	//fd = 4 ====> ExtraFiles[1]
	f := os.NewFile(3, "")
	count := 0
	for {
		str := fmt.Sprintf("hello, i'child process, write:%d line\n", count)
		count++
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Printf("write string failed, err:%v\n", err)
			continue
		}
		time.Sleep(time.Second)
	}

}

func parentProcess() {
	file, err := os.OpenFile("./temp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	_, err = file.WriteString("parent write one line\n")
	if err != nil {
		fmt.Printf("parent write failed, err:%v\n", err)
		return
	}
	wg.Add(1)
	go startChild(file)
	wg.Wait()
	fmt.Printf("parent exited\n")
}

func main() {
	if child != nil && *child == true {
		fmt.Printf("继承于父进程的文件句柄\n")
		readFromParent()
		return
	}
	//父进程
	parentProcess()
}
