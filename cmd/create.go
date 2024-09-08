package cmd

import (
	"fmt"
	"os"

	"github.com/pandasoncode/go-todo-cli/models"
	"github.com/spf13/cobra"
)

var task models.Task
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createTask(task)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&task.Title, "title", "t", "", "The title of the task")
	createCmd.Flags().StringVarP(&task.Description, "description", "d", "", "The description of the task")
	createCmd.Flags().StringVarP(&task.Category, "category", "c", "", "The category of the task")
}

func createTask(task models.Task) {
	filePath := "todo.txt"

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	newTask := task.String() + "\n"

	if _, err := file.WriteString(newTask); err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Task created successfully.")
}
