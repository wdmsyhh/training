package main

import (
	"gRPC/message"
	"google.golang.org/grpc"
	"net"
	"context"
	"math"
	"google.golang.org/grpc/metadata"
	"fmt"
)

type CalculateServiceImpl struct {
}

func (cs *CalculateServiceImpl) CheckTriangle(ctx context.Context, request *message.CalculateRequest) (*message.CheckResponse, error) {

	var response *message.CheckResponse

	if ( request.FirstSide + request.SecondSide > request.ThirdSide && request.FirstSide + request.ThirdSide >  request.SecondSide && request.SecondSide + request.ThirdSide > request.FirstSide) {
		response = &message.CheckResponse{CheckResult: true}
	} else {
		response = &message.CheckResponse{CheckResult: false}
	}
	return response, nil
}

func (cs *CalculateServiceImpl) CalculateCircumferenceOrArea(ctx context.Context, request *message.CalculateRequest) (*message.CalculateResponse, error) {

	//获取 metadata 数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	metadataSlice, _ := md["key"]
	fmt.Println(metadataSlice[0])

	var response *message.CalculateResponse
	p := (request.FirstSide + request.SecondSide + request.ThirdSide)/2

	//计算周长
	if (metadataSlice[0] == "circumference") {
		calculateResult := request.FirstSide + request.SecondSide + request.ThirdSide
		response = &message.CalculateResponse{CalculateResult: calculateResult, CircumferenceOrArea: metadataSlice[0]}
	}
	//计算面积
	if (metadataSlice[0] == "area") {
		calculateResult := math.Sqrt(float64(p*(p - request.FirstSide)*(p - request.SecondSide)*(p - request.ThirdSide)))
		response = &message.CalculateResponse{CalculateResult: float32(calculateResult), CircumferenceOrArea: metadataSlice[0]}
	}
	return response, nil
}

func main() {
	server := grpc.NewServer()
	message.RegisterCalculateServiceServer(server, new(CalculateServiceImpl))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}