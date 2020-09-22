package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "l2class",
	Short: "l2class is a list of Pro and Cons for each class",
	Long:  "l2 is a list of Pros and Cons for each class. because lineage 2 has many class so if you want to pick some class you have to know pros and cons for each class",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
