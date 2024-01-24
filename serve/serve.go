package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/cristianino/gRPC-unary-pattern-reminder-service.practice/protoService"
	"google.golang.org/grpc"
)

type ItaskRepository interface {
	CreateTask(*protoService.Task)
	GetTasks() []*protoService.Task
}

type taskStore struct {
	mu    sync.Mutex
	Tasks []*protoService.Task
}

func (ts *taskStore) CreateTask(task *protoService.Task) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.Tasks = append(ts.Tasks, task)
}

func (ts *taskStore) GetTasks() []*protoService.Task {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	tasksCopy := make([]*protoService.Task, len(ts.Tasks))
	copy(tasksCopy, ts.Tasks)

	return tasksCopy
}

type server struct {
	TaskRepo ItaskRepository
}

func (s *server) CreateTask(ctx context.Context, req *protoService.TaskRequest) (*protoService.TaskResponse, error) {
	fmt.Printf("CreateTask function was called with %v \n", req)
	task := req.GetTask()

	s.TaskRepo.CreateTask(task)
	response := &protoService.TaskResponse{Message: "Task created successfully"}

	return response, nil
}

func (s *server) GetTasks(ctx context.Context, req *protoService.GetTasksRequest) (*protoService.TasksResponse, error) {
	fmt.Printf("GetTasks function was called with %v \n", req)
	tasks := s.TaskRepo.GetTasks()

	response := &protoService.TasksResponse{Tasks: tasks}

	return response, nil
}

func main() {
	fmt.Println("remider service, Go server is runnig")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Fail to listen %v", err)
	}
	taskRepo := &taskStore{Tasks: make([]*protoService.Task, 0)}
	s := grpc.NewServer()

	protoService.RegisterRemiderServiceServer(s, &server{taskRepo})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
