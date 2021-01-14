// Package main implements a client for Greeter service.
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultAddress = "localhost:4300"
	defaultName    = "world"
	caCert         = "cert/ca-cert.pem"
	clientCert     = "cert/client-cert.pem"
	clientKey      = "cert/client-key.pem"
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		return nil, err
	}

	// Load certificate of the CA who signed server's certificate
	certPool := x509.NewCertPool()
	pemServerCA, err := ioutil.ReadFile(caCert)
	if err != nil {
		return nil, err
	}
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	serverAddress := flag.String("address", defaultAddress, "the server address")
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// Set up a connection to the server.
	log.Printf("Dialing on %v ...", *serverAddress)
	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("Could not dial : %s, Error : %v", *serverAddress, err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn) // Create new client

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Print(err.Error())
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
