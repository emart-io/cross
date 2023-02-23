package main

import (
	"net"

	impl "github.com/emart.io/cross/zwan/internal/impl/service"
	pb "github.com/emart.io/cross/zwan/service/go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	//go serveFile()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAttendantsServer(s, &impl.AttendantsImpl{})
	//pb.RegisterUsersServer(s, &impl.UsersImpl{})
	//pb.RegisterOrdersServer(s, &impl.OrdersImpl{})
	//pb.RegisterCouponsServer(s, &impl.CouponImpl{})
	//pb.RegisterAccountsServer(s, &impl.AccountImpl{})
	//pb.RegisterCommentsServer(s, &impl.CommentImpl{})
	//pb.RegisterMessagesServer(s, &impl.MessageImpl{})
	//pb.RegisterAddressesServer(s, &impl.AddressImpl{})
	//pb.RegisterCommoditiesServer(s, &impl.CommoditiesImpl{})

	log.Infoln("listen:" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
