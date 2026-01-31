package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gopherlibs/todoist/api"

	"github.com/spf13/cobra"
)

// deleteCmd represents the standings command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
