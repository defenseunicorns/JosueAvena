package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-task/task/v3"
	"github.com/go-task/task/v3/taskfile"
	"github.com/spf13/cobra"
)

var defaultTask = "default"

var taskfilePath string

var confirm bool

var rootCmd = &cobra.Command{
	Use:   "main COMMAND",
	Short: "My CLI program",
}
var toolsCmd = &cobra.Command{
	Use:   "tools COMMAND",
	Short: "Tools related commands",
}

var taskCmd = &cobra.Command{
	Use:   "task TASK",
	Short: "Run tasks using go-task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = append(args, defaultTask)
		}
		runTask(args)
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)
	toolsCmd.AddCommand(taskCmd)
	// Add flag for taskfile path
	taskCmd.Flags().StringVarP(&taskfilePath, "file", "f", "Taskfile.yml", "Path to Taskfile")
	taskCmd.Flags().BoolVarP(&confirm, "confirm", "c", false, "Assume yes to prompts")
}

func runTask(args []string) {
	dir := taskfilePath // Directory where the Taskfile.yml exists

	e := task.Executor{
		Dir:       dir,
		Stdout:    log.Writer(),
		Stderr:    log.Writer(),
		Color:     true,
		AssumeYes: confirm,
	}

	if err := e.Setup(); err != nil {
		fmt.Fprintln(os.Stderr, "Executor setup failed:", err)
		os.Exit(1)
	}
	// Create a new context for the task run
	ctx := context.Background()

	// Convert string slice args to taskfile.Call slice
	var calls []taskfile.Call
	for _, a := range args {
		calls = append(calls, taskfile.Call{Task: a})
	}

	err := e.Run(ctx, calls...)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Task failed:", err)
		os.Exit(1)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
