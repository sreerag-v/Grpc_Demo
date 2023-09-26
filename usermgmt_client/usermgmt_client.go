package main

import (
	"context"
	"log"
	"time"

	pb "GRPC/unary_grpc1/usermgmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("faild tot connect:%v", err)
	}

	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cance := context.WithTimeout(context.Background(), time.Second)

	defer cance()

	var new_user = make(map[string]int32)

	new_user["Sreerag"] = 21
	new_user["vaishak"] = 21

	for name, age := range new_user {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})

		if err != nil {
			log.Fatalf("could not create user:%v", err)
		}

		log.Printf(`UserDetails:
		Name:%s
		Age:%d
		Id:%d`, r.GetName(), r.GetAge(), r.GetId())
	}

}
