package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cristianino/gRPC-unary-pattern-reminder-service.practice/protoService"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Fail to connect %v", err)
	}

	defer cc.Close()

	c := protoService.NewRemiderServiceClient(cc)

	// Presenta la interfaz al usuario
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Create a task")
		fmt.Println("2. Get tasks")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}

		switch choice {
		case 1:
			createTask(c)
		case 2:
			getTasks(c)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

func createTask(c protoService.RemiderServiceClient) {
	fmt.Println("Creating a task")

	// Solicitar datos al usuario
	task := &protoService.Task{}
	fmt.Print("Enter task message: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	task.Message = scanner.Text()

	fmt.Print("Enter task limit date: ")
	scanner.Scan()
	task.LimitDate = scanner.Text()

	fmt.Print("Enter task team name: ")
	scanner.Scan()
	task.TeamName = scanner.Text()

	task.Priority = protoService.PRIORITY_HIGH

	// Crear la solicitud y llamar al servidor
	req := &protoService.TaskRequest{
		Task: task,
	}

	res, err := c.CreateTask(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling createTask RPC: %v", err)
	}

	log.Printf("Response createTask: %v", res.Message)
}

func getTasks(c protoService.RemiderServiceClient) {
	fmt.Println("Getting tasks")

	// Solicitar datos al usuario
	fmt.Print("Enter a message (leave empty for all tasks): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	// Crear la solicitud y llamar al servidor
	req := &protoService.GetTasksRequest{
		Message: message,
	}

	res, err := c.GetTasks(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling getTasks RPC: %v", err)
	}

	log.Printf("Response getTasks: %v", res.Tasks)
}
