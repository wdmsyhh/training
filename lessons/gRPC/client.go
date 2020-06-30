package main

import (
	"google.golang.org/grpc"
	"gRPC/message"
	"context"
	"fmt"
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

	if checkRsponse != nil {
		fmt.Println(checkRsponse.CheckResult)
		//判断是三角形再进行计算
		if checkRsponse.CheckResult == true {
			calculateResponse, _ := client.CalculateCircumferenceOrArea(ctx, calculateRequest)
			if calculateResponse != nil {
				fmt.Println(calculateResponse.CircumferenceOrArea, calculateResponse.CalculateResult)
			}
		}
	}
}