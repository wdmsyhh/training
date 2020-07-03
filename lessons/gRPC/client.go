package main

import (
	"context"
	"fmt"
	"gRPC/message"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	client := message.NewCalculateServiceClient(conn)

	//创建 metadata 值为 circumference 计算周长，值为 area 计算面积
	md := metadata.Pairs("key", "circumference")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	calculateRequest := &message.CalculateRequest{FirstSide: 2.0, SecondSide: 3.0, ThirdSide: 4.0}

	checkRsponse, err := client.CheckTriangle(context.Background(), calculateRequest)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(checkRsponse.CheckResult)
	//判断是三角形再进行计算
	if checkRsponse.CheckResult == false {
		return
	}
	calculateResponse, err := client.CalculateCircumferenceOrArea(ctx, calculateRequest)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(calculateResponse.CircumferenceOrArea, calculateResponse.CalculateResult)
}
