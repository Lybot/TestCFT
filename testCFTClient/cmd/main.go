package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	pb "lybot/testCFTCLient/grpcService"
	"os"

	"github.com/BurntSushi/toml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

type Config struct {
	Server_address string `toml:"Server_addr"`
	Server_port    string `toml:"Server_port"`
	SDOHNI         string `toml:"SDOHNI"`
}

func NewConfig() *Config {
	return &Config{}
}

var (
	server_address   string
	server_port      string
	command          string
	command_argument string
	file_save_path   string
)

func init() {
	flag.StringVar(&server_address, "addr", "", "server address")
	flag.StringVar(&server_port, "p", "", "server port")
	flag.StringVar(&command, "c", "get_files", "command to execute (get_files, get_file (need file_save_path), delete_file, update_file, put_file")
	flag.StringVar(&command_argument, "path", "test.txt", "path to file to work")
	flag.StringVar(&file_save_path, "s_path", "test.txt", "path to save file from server (only need in get_file command)")
}
func main() {
	flag.Parse()
	if server_address == "" || server_port == "" {
		config := NewConfig()
		_, err := toml.DecodeFile("./config.toml", config)
		if err != nil {
			fmt.Println("Error config")
			return
		}
		server_address = config.Server_address
		server_port = config.Server_port
	}
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(server_address+":"+server_port, opts)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCftServiceClient(conn)
	switch command {
	case "get_files":
		get_files(client)
	case "get_file":
		get_file(client, command_argument, file_save_path)
	case "delete_file":
		delete_file(client, command_argument)
	case "update_file":
		update_file(client, command_argument)
	case "put_file":
		put_file(client, command_argument)
	default:
		fmt.Println("Incorrent command")
	}
}
func get_files(client pb.CftServiceClient) {
	request := &pb.Empty{}
	response, err := client.GetFiles(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	for _, value := range response.Files {
		println("Filename: " + value.Name + ", Hash: " + hex.EncodeToString(value.Data))
	}
}
func get_file(client pb.CftServiceClient, file_name string, save_path string) {
	request := &pb.FileWithData{
		Name: file_name,
	}
	response, err := client.GetFile(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response)
	if response.Error == "" {
		os.WriteFile(save_path, response.Files[0].Data, os.ModeAppend)
		println(save_path + " was successfully downloaded")
	}
}
func delete_file(client pb.CftServiceClient, file_name string) {
	request := &pb.FileWithData{
		Name: file_name,
	}
	response, err := client.DeleteFile(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	if response.Error == "" {
		println(file_name + " was sucessfully removed")
	}
}
func update_file(client pb.CftServiceClient, file_name string) {
	file, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println("File doesn't exist")
		return
	}
	request := &pb.FileWithData{
		Name: file_name,
		Data: file,
	}
	response, err := client.UpdateFile(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	if response.Error == "" {
		println(file_name + "was successfully updated")
	}
}
func put_file(client pb.CftServiceClient, file_name string) {
	file, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println("File doesn't exist")
		return
	}
	request := &pb.FileWithData{
		Name: file_name,
		Data: file,
	}
	response, err := client.CreateFile(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	if response.Error == "" {
		println(file_name + " successfully created")
	}
}
