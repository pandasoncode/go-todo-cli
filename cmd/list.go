package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pandasoncode/go-todo-cli/models"
	"github.com/spf13/cobra"
)

var category string
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if category != "" {
			for _, task := range getTasksByCategory(category) {
				fmt.Println(task)
			}
		} else {
			for _, task := range getTasks() {
				fmt.Println(task)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&category, "category", "c", "", "The category of the tasks")
}

func getTasks() []models.Task {
	var tasks []models.Task

	filePath := "todo.txt"
	var task models.Task

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		json.Unmarshal([]byte(linea), &task)
		tasks = append(tasks, task)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}

	return tasks
}

func getTasksByCategory(category string) []models.Task {
	var tasks []models.Task

	for _, task := range getTasks() {
		if task.Category == category {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
