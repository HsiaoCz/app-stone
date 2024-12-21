package main

import (
	"fmt"
	"os"

	"github.com/HsiaoCz/app-stone/go-todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "A simple CLI todo list application",
	}

	var addCmd = &cobra.Command{
		Use:   "add [task]",
		Short: "Add a new todo item",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task := args[0]
			todo.AddTodo(task)
			fmt.Printf("Added todo: %s\n", task)
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all todo items",
		Run: func(cmd *cobra.Command, args []string) {
			todos := todo.ListTodos()
			for _, t := range todos {
				fmt.Println(t)
			}
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete [task]",
		Short: "Delete a todo item",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task := args[0]
			if err := todo.DeleteTodo(task); err != nil {
				fmt.Printf("Error deleting todo: %s\n", err)
			} else {
				fmt.Printf("Deleted todo: %s\n", task)
			}
		},
	}

	rootCmd.AddCommand(addCmd, listCmd, deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
