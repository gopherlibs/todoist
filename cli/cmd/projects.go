package cmd

import (
	"fmt"

	"github.com/gopherlibs/todoist/api"

	"github.com/spf13/cobra"
)

// projectsCmd represents the standings command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {

		c := api.New()

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
