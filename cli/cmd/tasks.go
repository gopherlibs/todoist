package cmd

import (
	"fmt"

	"github.com/gopherlibs/todoist/api"

	"github.com/spf13/cobra"
)

// tasksCmd represents the standings command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		c := api.New()

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
