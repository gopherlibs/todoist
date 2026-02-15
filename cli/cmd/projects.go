package cmd

import (
	"fmt"
	"os"

	"github.com/gopherlibs/todoist/api"

	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {

		c := api.New(os.Getenv("TODOIST_TOKEN"))

		projects, err := c.Projects()
		if err != nil {
			fmt.Printf("Running 'Projects' failed. Err: %s\n", err.Error())
			return err
		}

		for _, p := range projects.Results {
			fmt.Printf("%s - %s\n", p.Name, p.ID)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
