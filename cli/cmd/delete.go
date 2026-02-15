package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gopherlibs/todoist/api"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <task-id>",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if args[0] == "" {
			return errors.New("error - arg")
		}

		c := api.New(os.Getenv("TODOIST_TOKEN"))

		status, err := c.TaskClose(args[0])
		if err != nil {
			fmt.Printf("Running 'Delete' failed. Err: %s\n", err.Error())
			return err
		}

		fmt.Printf("The returned code was: %d\n", status)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
