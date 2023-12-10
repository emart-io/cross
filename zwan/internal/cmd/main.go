package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	impl "github.com/emart.io/cross/zwan/internal/impl/service"
	pb "github.com/emart.io/cross/zwan/service/go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//go:embed all:dist
var ui embed.FS

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterAttendantsServer(grpcServer, &impl.AttendantsImpl{})
	//pb.RegisterUsersServer(s, &impl.UsersImpl{})
	pb.RegisterOrdersServer(grpcServer, &impl.OrdersImpl{})
	//pb.RegisterCouponsServer(s, &impl.CouponImpl{})
	//pb.RegisterAccountsServer(s, &impl.AccountImpl{})
	//pb.RegisterCommentsServer(s, &impl.CommentImpl{})
	//pb.RegisterMessagesServer(s, &impl.MessageImpl{})
	//pb.RegisterAddressesServer(s, &impl.AddressImpl{})
	//pb.RegisterCommoditiesServer(s, &impl.CommoditiesImpl{})

	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	mux := http.NewServeMux()
	dist, err := fs.Sub(ui, "dist")
	if err != nil {
		log.Fatalf("sub error: %s", err)
		return
	}
	mux.Handle("/", fileServerWith404(http.FS(dist)))
	log.Infoln("listen:" + port)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", port), certFile, keyFile, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// if request.ProtoMajor != 2 {
		// 	mux.ServeHTTP(writer, request)
		// 	return
		// }
		// Handle gRPC-web
		if request.Header.Get("Content-Type") == "application/grpc-web+proto" {
			request.Header.Set("Content-Type", "application/grpc")
		}
		if strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(writer, request)
			return
		}
		mux.ServeHTTP(writer, request)
	})))
}

// Handle 404 to "/index.html"
func fileServerWith404(root http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := root.Open(r.URL.Path)
		if err != nil && os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		if err == nil {
			f.Close()
		}
		http.FileServer(root).ServeHTTP(w, r)
	})
}
