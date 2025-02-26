package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc-todo/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatal("Failed to connect:", err)
    }
    defer conn.Close()

    client := proto.NewTodoServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Add a new task
    addTask(ctx, client, "Learn gRPC")

    // Get all tasks
    getTasks(ctx, client)

    // Complete a task
    completeTask(ctx, client, 1)

    // Get all tasks again to see the updated status
    getTasks(ctx, client)

    // Delete a task
    deleteTask(ctx, client, 1)

    // Get all tasks again to see the updated list
    getTasks(ctx, client)
}

func addTask(ctx context.Context, client proto.TodoServiceClient, title string) {
    res, err := client.AddTask(ctx, &proto.TaskRequest{Title: title})
    if err != nil {
        log.Fatal("Failed to add task:", err)
    }
    fmt.Println("New Task:", res.Task)
}

func getTasks(ctx context.Context, client proto.TodoServiceClient) {
    res, err := client.GetTasks(ctx, &emptypb.Empty{})
    if err != nil {
        log.Fatal("Failed to get tasks:", err)
    }
    fmt.Println("Tasks:")
    for _, task := range res.Tasks {
        fmt.Printf("- %d: %s (completed: %v)\n", task.Id, task.Title, task.Completed)
    }
}

func completeTask(ctx context.Context, client proto.TodoServiceClient, id int32) {
    res, err := client.CompleteTask(ctx, &proto.Task{Id: id})
    if err != nil {
        log.Fatal("Failed to complete task:", err)
    }
    fmt.Println("Completed Task:", res.Task)
}

func deleteTask(ctx context.Context, client proto.TodoServiceClient, id int32) {
    _, err := client.DeleteTask(ctx, &proto.Task{Id: id})
    if err != nil {
        log.Fatal("Failed to delete task:", err)
    }
    fmt.Println("Deleted Task with ID:", id)
}
