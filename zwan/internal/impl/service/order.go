package service

import (
	"context"

	pb "github.com/emart.io/cross/zwan/service/go"
	"github.com/google/uuid"
	"github.com/jmzwcn/db"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	orderTable = "orders"
)

type OrdersImpl struct {
	pb.UnimplementedOrdersServer
}

func (s *OrdersImpl) Add(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	in.Id = uuid.NewString()
	in.Created = timestamppb.Now()
	if err := db.Upsert(orderTable, in.Id, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *OrdersImpl) Get(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	order := pb.Order{}
	if err := db.GetById(orderTable, in.Id, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrdersImpl) Update(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	return in, nil
}

func (s *OrdersImpl) List(in *pb.Order, stream pb.Orders_ListServer) error {
	orders := []*pb.Order{}
	if err := db.List(orderTable, &orders, " order by data->'$.created' desc"); err != nil {
		return err
	}

	for _, v := range orders {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrdersImpl) Delete(ctx context.Context, in *pb.Order) (*emptypb.Empty, error) {
	if err := db.Delete(orderTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
