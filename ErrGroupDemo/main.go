package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		str := []byte("hello")
		_, _ = w.Write(str)
		fmt.Println(string(str))
	})

	// http请求，触发server退出
	serverOut := make(chan int)
	httpHandler.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		log.Println("will shut down server...")
		serverOut <- 1
	})

	server := http.Server{
		Handler: httpHandler,
		Addr:    ":8080",
	}

	// g1 - http server
	// g1 通过 g2 引导退出
	g.Go(func() error {
		return server.ListenAndServe()
	})

	// g2 - 引导g1退出
	// g2 退出的前提，满足其一：1、http shutdown接口触发退出；2、g3 退出，ctx返回信号
	g.Go(func() error {
		select {
			case <-ctx.Done():
				log.Println("[g2] - in - errgroup exit...")
			case <-serverOut:
				log.Println("[g2] - in - server will out...")
		}

		timeoutCtx, _ := context.WithTimeout(context.Background(), 3 * time.Second)

		log.Println("[g1] shut down server...")
		err := server.Shutdown(timeoutCtx)
		log.Printf("[g2] - out - %v \n", err)
		return err
	})

	// g3
	// g3 退出的前提，满足其一：1、捕获到 os 退出信号将会退出，然后促使 g2 退出；2、g2 使 g1 退出，ctx返回信号
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
			case <-ctx.Done():
				log.Println("[g3] errgroup exit...")
				return ctx.Err()
			case sig := <-quit:
				log.Println("[g3] os exit...")
				return fmt.Errorf("get os signal: %v", sig)
		}
	})

	// 退出路径
	// 1、[g1] http请求 -> g1发送信号 -> [g2]继续 -> [g1]退出 -> g3[退出] -> [g2]退出 -> wait退出
	// 2、[g3] os中断 -> [g3]退出 -> [g2]继续 -> [g1]退出  -> [g2]退出 -> wait退出

	err := g.Wait()

	fmt.Printf("errgroup exiting: %+v\n", err)
}
