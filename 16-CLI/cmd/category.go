/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// categoryName, _ := cmd.Flags().GetString("name")
		fmt.Println("category called with name: " + categoryName)
		exists, _ := cmd.Flags().GetString("exists")
		fmt.Println("category called with exists: " + fmt.Sprint(exists))
		id, _ := cmd.Flags().GetString("id")
		fmt.Println("category called with id: " + fmt.Sprint(id))
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado antes do Run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado depois do Run")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("ocorreu um erro")
	},
}

var categoryName string

func init() {
	rootCmd.AddCommand(categoryCmd)
	// Persistent or Global - createCmd and all subCmds
	// categoryCmd.PersistentFlags().StringP("name", "n", "Default", "Name of the category")
	categoryCmd.PersistentFlags().StringVarP(&categoryName, "name", "n", "Default", "Name of the category")

	// Local flag, only for createCmd
	categoryCmd.PersistentFlags().StringP("type", "t", "Default", "Name of the category")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if a category exists")
	categoryCmd.PersistentFlags().Int16P("id", "i", 0, "ID of a category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
