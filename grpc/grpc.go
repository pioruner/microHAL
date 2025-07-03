package grpc

import (
	"context"
	"microHAL/devices"
	"sync"
)

import (
	pb "microHAL/device-service/proto"
)

type GRPCServer struct {
	pb.UnimplementedDeviceServiceServer
	Devices map[string]devices.Device
	Mu      sync.Mutex
}

func NewServer() *GRPCServer {
	return &GRPCServer{
		Devices: make(map[string]devices.Device),
	}
}

func (s *GRPCServer) Connect(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	dev := &devices.TCPDevice{Address: req.Id} // req.Id содержит TCP-адрес
	err := dev.Connect()
	if err != nil {
		return &pb.DeviceResponse{Success: false, Error: err.Error()}, nil
	}

	s.Devices[req.Id] = dev
	return &pb.DeviceResponse{Success: true}, nil
}

func (s *GRPCServer) Read(ctx context.Context, req *pb.DeviceRequest) (*pb.ReadResponse, error) {
	s.Mu.Lock()
	dev, ok := s.Devices[req.Id]
	s.Mu.Unlock()
	if !ok {
		return &pb.ReadResponse{Error: "device not found"}, nil
	}
	data, err := dev.Read()
	if err != nil {
		return &pb.ReadResponse{Error: err.Error()}, nil
	}
	return &pb.ReadResponse{Data: data}, nil
}
