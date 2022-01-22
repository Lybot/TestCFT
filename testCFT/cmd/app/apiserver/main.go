package main

import (
	// "log"
	"bytes"
	"context"
	"flag"
	"log"
	pb "lybot/testCFT/grpcService"
	"lybot/testCFT/internal/app/apiserver"
	"net"
	"os"

	"github.com/BurntSushi/toml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	config_path string
)

func init() {
	flag.StringVar(&config_path, "config_path", "config.toml", "path to config file")
}
func main() {
	// print("kekw")
	// server := apiserver.New()
	// if err := server.Start(); err!=nil{
	// 	log.Fatal(err)
	// }
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(config_path, config)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", config.Bind_addr)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCftServiceServer(grpcServer, &server{})
	println("Server started at port" + config.Bind_addr)
	grpcServer.Serve(listener)
}

type server struct {
	pb.UnimplementedCftServiceServer
}

func (s *server) CreateFile(c context.Context, request *pb.FileWithData) (response *pb.Response, err error) {
	response = &pb.Response{
		Error: "",
		Files: nil,
	}
	if _, err := os.Stat("../../../tmp/" + request.Name); err == nil {
		response.Error = "File exists"
		return response, nil
	}
	os.WriteFile("../../../tmp/"+request.Name, request.Data, os.ModeAppend)
	response.Error = ""
	println(request.Name + "file created")
	return response, nil
}
func (s *server) DeleteFile(c context.Context, request *pb.FileWithData) (response *pb.Response, err error) {
	response = &pb.Response{
		Error: "",
		Files: nil,
	}
	if _, err := os.Stat("../../../tmp/" + request.Name); err != nil {
		response.Error = "File doesn't exist"
	}
	os.Remove("../../../tmp/" + request.Name)
	println(request.Name + "file created")
	return response, nil
}
func (s *server) UpdateFile(c context.Context, request *pb.FileWithData) (response *pb.Response, err error) {
	response = &pb.Response{
		Error: "",
		Files: nil,
	}
	if _, err := os.Stat("../../../tmp/" + request.Name); err != nil {
		response.Error = "File doesn't exist"
		println(request.Name + " " + response.Error)
		return response, nil
	}
	hash, erro := apiserver.GetHash("../../../tmp/" + request.Name)
	requestHash, _ := apiserver.GetHashBytes(request.Data)
	if erro != nil || bytes.Equal(hash, requestHash) {
		response.Error = "Same file"
		return response, nil
	}
	os.Remove("../../../tmp/" + request.Name)
	os.WriteFile("../../../tmp/"+request.Name, request.Data, os.ModeAppend)
	response.Error = ""
	println(request.Name + "file updated")
	return response, nil
}
func (s *server) GetFile(c context.Context, request *pb.FileWithData) (response *pb.Response, err error) {
	file, err := os.ReadFile("../../../tmp/" + request.Name)
	response = &pb.Response{
		Error: "",
		Files: nil,
	}
	if err != nil {
		response.Error = err.Error()
		return response, nil
	}
	response = &pb.Response{
		Error: "",
		Files: []*pb.FileWithData{
			{
				Data: file,
				Name: request.Name,
			},
		},
	}
	println(request.Name + " was sent")
	return response, nil
}
func (s *server) GetFiles(c context.Context, request *pb.Empty) (response *pb.Response, err error) {
	dirPath := "../../../tmp/"
	files, err := os.ReadDir(dirPath)
	response = &pb.Response{
		Error: "",
		Files: nil,
	}
	if err != nil {
		response.Error = err.Error()
		return response, nil
	}

	result := []*pb.FileWithData{}
	for _, file := range files {
		hash, hashErr := apiserver.GetHash(dirPath + file.Name())
		if hashErr != nil {
			response.Error = hashErr.Error()
			return response, nil
		}
		temp := pb.FileWithData{
			Name: file.Name(),
			Data: hash,
		}
		result = append(result, &temp)
	}
	response = &pb.Response{
		Error: "",
		Files: result,
	}
	println("Отдал список файлов")
	return response, nil
}
