package main


import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "GRPC/unary_grpc1/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct{
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser)(*pb.User,error){
	log.Printf("Receved:%v",in.GetName())

	var user_id int32 = int32(rand.Intn(1000))

	return &pb.User{Name: in.GetName(),Age: in.GetAge(),Id: user_id},nil
}

func main(){
	lis,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("faild to listen:%v",err)
	}

	s:=grpc.NewServer()

	pb.RegisterUserManagementServer(s,&UserManagementServer{})
	log.Printf("server listening %v",lis.Addr())
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("faild to serve :%v",err)
	}
}