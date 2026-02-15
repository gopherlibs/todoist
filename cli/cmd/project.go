package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gopherlibs/todoist/api"

	"github.com/spf13/cobra"
)

// projectCmd retrieves info on a project
var projectCmd = &cobra.Command{
	Use:   "project <project-id>",
	Short: "View a project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if args[0] == "" {
			return errors.New("error - arg")
		}

		c := api.New(os.Getenv("TODOIST_TOKEN"))

		p, err := c.Project(args[0])
		if err != nil {
			fmt.Printf("Running 'Project' failed. Err: %s\n", err.Error())
			return err
		}

		fmt.Printf("The project name is: %s\n", p.Name)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
