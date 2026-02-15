package cmd

import (
	"fmt"
	"os"

	"github.com/gopherlibs/todoist/api"

	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		c := api.New(os.Getenv("TODOIST_TOKEN"))

		tasks, err := c.Tasks(args[0])
		if err != nil {
			fmt.Printf("Running 'Tasks' failed. Err: %s\n", err.Error())
			return err
		}

		for _, t := range tasks.Results {
			fmt.Printf("- %s - %s\n", t.Content, t.ID)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
}
