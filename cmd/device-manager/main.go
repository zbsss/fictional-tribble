package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/zbsss/device-manager/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port          = flag.Int("port", 50051, "The server port")
	tokenLifetime = flag.Int("token", 250, "Lifetime of token in milliseconds")
	windowSize    = flag.Int("windowSize", 10000, "Window size in milliseconds")
)

type server struct {
	pb.UnimplementedDeviceManagerServer
}

var podMemoryQuota = map[string]float64{
	"pod-1": 1.0,
}

var podMemoryQuotaUsage = map[string]uint64{
	"pod-1": 0,
}

var deviceTotalMemory = map[string]uint64{
	"1": 1536,
}

var devicesMemoryUsage = map[string]uint64{}

func (s *server) GetToken(ctx context.Context, in *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	log.Printf("Received: GetToken for device %s", in.Device)

	time.Sleep(5 * time.Second)

	log.Printf("Sending token")

	expiresAt := time.Now().Add(time.Duration(*tokenLifetime) * time.Millisecond).UnixNano()

	return &pb.GetTokenReply{ExpiresAt: expiresAt}, nil
}

func (s *server) ReturnToken(ctx context.Context, in *pb.ReturnTokenRequest) (*pb.ReturnTokenReply, error) {
	log.Printf("Received: ReturnToken")

	return &pb.ReturnTokenReply{}, nil
}

func (s *server) GetMemoryQuota(ctx context.Context, in *pb.GetMemoryQuotaRequest) (*pb.GetMemoryQuotaReply, error) {
	log.Printf("Received: GetMemoryQuota for device %s: %d", in.Device, in.Memory)

	podId := "pod-1"

	if deviceTotalMemory[in.Device] < devicesMemoryUsage[in.Device]+in.Memory ||
		uint64(podMemoryQuota[podId]*float64(deviceTotalMemory[in.Device])) < podMemoryQuotaUsage[podId]+in.Memory {
		return nil, fmt.Errorf("OOM: memory quota exceeded")
	}

	devicesMemoryUsage[in.Device] += in.Memory
	podMemoryQuotaUsage[podId] += in.Memory

	log.Println(devicesMemoryUsage[in.Device])

	return &pb.GetMemoryQuotaReply{}, nil
}

func (s *server) ReturnMemoryQuota(ctx context.Context, in *pb.ReturnMemoryQuotaRequest) (*pb.ReturnMemoryQuotaReply, error) {
	log.Printf("Received: ReturnMemoryQuota for device %s: %d", in.Device, in.Memory)

	podId := "pod-1"

	devicesMemoryUsage[in.Device] -= in.Memory
	podMemoryQuotaUsage[podId] -= in.Memory

	log.Println(devicesMemoryUsage[in.Device])

	return &pb.ReturnMemoryQuotaReply{}, nil
}

func main() {
	flag.Parse()

	go func() {
		for {
			log.Println("Memory usage per device: ", devicesMemoryUsage)
			log.Println("Memory usage per pod: ", devicesMemoryUsage)

			time.Sleep(10 * time.Second)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterDeviceManagerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
