package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var (
	server   *http.Server
	listener net.Listener
	graceful = flag.Bool("graceful", false, "listen on fd open 3(internal use only)")
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	w.Write([]byte("hello world !!!"))
}

func main() {
	flag.Parse()

	http.HandleFunc("/hello", handler)
	server = &http.Server{
		Addr: ":9000",
	}

	var err error
	if *graceful {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		log.Print("main: Listening on a new file descriptor.")
		listener, err = net.Listen("tcp", server.Addr)
	}

	if err != nil {
		log.Fatalf("listener error: %v\n", err)
	}

	go func() {
		// server.Shutdown() stops Serve() immediately, this server.Serve() should not be in main goroutine
		err = server.Serve(listener)
		log.Printf("server.Serve err: %v\n", err)
	}()
	signalHandler()
	log.Printf("signal end")
}

func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		log.Printf("signal: %v\n", sig)

		// timeout context for shutdown
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			// stop
			log.Printf("stop")
			signal.Stop(ch)
			server.Shutdown(ctx)
			log.Printf("graceful shutdown\n")
			return
		case syscall.SIGUSR2:
			// reload
			log.Printf("reload")
			err := reload()
			if err != nil {
				log.Fatalf("graceful restart error: %v\n", err)
			}
			server.Shutdown(ctx)
			log.Printf("graceful reload")
			return
		}
	}
}

func reload() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//put socket fd at the first enty
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}
