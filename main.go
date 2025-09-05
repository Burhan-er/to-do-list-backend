package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"to-do-list/db"

	"github.com/urfave/cli/v3"
)

const tasksFilePath = "tasks.json"

var database *db.Db

func main() {
	homeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(homeDir, ".todo.json")
	database = db.NewDatabase(dbPath)

	app := &cli.Command{
		Name:  "ToDo-List",
		Usage: "A simple command-line todo list application.",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a new todo task.",
				Action:  add,
			}, {
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "lists all todo tasks.",
				Action:  list,
			}, {
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "when todo task is done.",
				Action:  done,
			}, {
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "delete a todo task.",
				Action:  remove,
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	
}

func add(ctx context.Context, cmd *cli.Command) error {
	taskText := cmd.Args().First()
	if taskText == "" {
		return fmt.Errorf("task cannot be empty")
	}
	return database.Add(taskText)
}
func list(ctx context.Context, cmd *cli.Command) error {
	tasks, err := database.GetTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks to display. Use 'add' to add a new task.")
		return nil
	}

	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, i+1, task.Text)
	}
	return nil
}

func done(ctx context.Context, cmd *cli.Command) error {
	arg := cmd.Args().First()
	if arg == "" {
		return fmt.Errorf("task index is not specified")
	}

	index, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("invalid index: %s", arg)
	}

	if err := database.Complete(index - 1); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as done.\n", index)
	return nil
}

func remove(ctx context.Context, cmd *cli.Command) error {
	arg := cmd.Args().First()
	if arg == "" {
		return fmt.Errorf("task index is not specified")
	}

	index, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("invalid index: %s", arg)
	}

	if err := database.Delete(index - 1); err != nil {
		return err
	}

	fmt.Printf("Task %d removed.\n", index)
	return nil
}
