package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pandasoncode/go-todo-cli/models"
	"github.com/spf13/cobra"
)

var title string
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getTaskByTitle(title))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&title, "title", "t", "", "The title of the task")
}

func getTaskByTitle(title string) *models.Task {
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
		if task.Title == title {
			return &task
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}

	return nil
}
