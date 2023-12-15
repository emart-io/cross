package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/NYTimes/gziphandler"
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
	mux.Handle("/", fileServerWithExt(http.FS(dist)))
	log.Infoln("listen:" + port)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", port), certFile, keyFile, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Handle gRPC-web
		if request.Header.Get("Content-Type") == "application/grpc-web+proto" {
			request.Header.Set("Content-Type", "application/grpc")
		}
		if request.ProtoMajor == 2 && strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(writer, request)
			return
		}
		mux.ServeHTTP(writer, request)
	})))
}

// Handle 404 to "/index.html" and gzip compression
func fileServerWithExt(root http.FileSystem) http.Handler {
	return gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := root.Open(r.URL.Path)
		if err != nil && os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		if err == nil {
			f.Close()
		}
		http.FileServer(root).ServeHTTP(w, r)
	}))
}
