package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"template/internal/myapp"
	"template/pkg/util"
	proto "template/proto/myapp"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ServerPort string
	CertName string
	CertPemPath string
	CertKeyPath string
	EndPoint string
)

func init() {
	ServerPort = "50052"
	CertName = "www.eline.com"
	CertPemPath = "./certs/server.pem"
	CertKeyPath = "./certs/server.key"
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	EndPoint = ":" + ServerPort
	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	tlsConfig := util.GetTLSConfig(CertPemPath, CertKeyPath)
	srv := createInternalServer(conn, tlsConfig)

	myapp.InitApp()

	log.Printf("gRPC and https listen on: %s\n", ServerPort)

	g.Go(func() error {
		return srv.Serve(tls.NewListener(conn, tlsConfig))
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		timeoutCtx, _ := context.WithTimeout(context.Background(), 3 * time.Second)

		select {
		case <-ctx.Done():
			log.Println("[g3] errgroup exit...")
			_ = srv.Shutdown(timeoutCtx)
			return ctx.Err()
		case sig := <-quit:
			log.Println("[g3] os exit...")
			_ = srv.Shutdown(timeoutCtx)
			return fmt.Errorf("get os signal: %v", sig)
		}
	})

	err = g.Wait()
	fmt.Printf("errgroup exiting: %+v\n", err)
}

func createInternalServer(conn net.Listener, tlsConfig *tls.Config) (*http.Server) {
	var opts []grpc.ServerOption

	// grpc server
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}

	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)

	// register grpc pb
	proto.RegisterHelloWorldServer(grpcServer, NewUserService())

	// gw server
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertName)
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb
	if err := proto.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}