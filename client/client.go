package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

	fmt.Println("Select task priority:")
	fmt.Println("0. Low")
	fmt.Println("1. Mid")
	fmt.Println("2. High")

	var priority int32
	fmt.Print("Enter priority (0, 1, or 2): ")
	_, err := fmt.Scan(&priority)
	if err != nil {
		log.Fatalf("Failed to read priority: %v", err)
	}

	// Validar la entrada de prioridad
	if priority < 0 || priority > 2 {
		log.Fatalf("Invalid priority. Please enter a valid priority (0, 1, or 2).")
	}

	task.Priority = protoService.PRIORITY(priority)

	fmt.Print("Enter task tags (comma-separated, leave empty for none): ")
	scanner.Scan()
	tagsInput := scanner.Text()

	// Dividir las tags por comas y agregarlas al slice de tags
	task.Tags = make([]string, 0)
	if tagsInput != "" {
		tags := strings.Split(tagsInput, ",")
		for _, tag := range tags {
			task.Tags = append(task.Tags, strings.TrimSpace(tag))
		}
	}

	// Crear la solicitud y llamar al servidor
	req := &protoService.TaskRequest{
		Task: task,
	}

	res, err := c.CreateTask(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling createTask RPC: %v", err)
	}

	// Imprimir detalles de la tarea recién creada
	if res.Task != nil {
		fmt.Printf("Task ID: %d\n", res.Task.Id)
		fmt.Printf("Task Message: %s\n", res.Task.Message)
		fmt.Printf("Task Limit Date: %s\n", res.Task.LimitDate)
		fmt.Printf("Task Team Name: %s\n", res.Task.TeamName)
		fmt.Printf("Task Priority: %s\n", res.Task.Priority.String())

		// Imprimir tags si existen
		if len(res.Task.Tags) > 0 {
			fmt.Printf("Task Tags: %s\n", strings.Join(res.Task.Tags, ", "))
		}
	} else {
		fmt.Println("No task details returned.")
	}
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

	formattedTasks := formatTasks(res.Tasks)
	fmt.Println(formattedTasks)
}

func formatTasks(tasks []*protoService.Task) string {
	if len(tasks) == 0 {
		return "No tasks found."
	}

	result := "Tasks:\n"
	for _, task := range tasks {
		result += fmt.Sprintf("     Task ID: %d\n", task.Id)
		result += fmt.Sprintf("    Message: %s\n", task.Message)
		result += fmt.Sprintf("    Limit Date: %s\n", task.LimitDate)
		result += fmt.Sprintf("    Team Name: %s\n", task.TeamName)
		result += fmt.Sprintf("    Priority: %s\n", task.Priority.String())
		if len(task.Tags) > 0 {
			result += fmt.Sprintf("    Tags: %s\n", strings.Join(task.Tags, ", "))
		}
		result += "\n"
	}

	return result
}
