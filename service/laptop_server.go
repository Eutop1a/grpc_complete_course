package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gitlab.com/techschool/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type LaptopServer struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopSever(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID %v", err)
		}
		laptop.Id = id.String()
	}

	// some heavy processing
	time.Sleep(6 * time.Second)

	if errors.Is(ctx.Err(), context.Canceled) {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		log.Print("deadline exceeded")
		return nil, status.Errorf(codes.DeadlineExceeded, "deadline is exceeded")
	}

	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
