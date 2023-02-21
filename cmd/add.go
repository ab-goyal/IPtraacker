/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition of numbers",
	Long:  `Addition of numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		sum := 0
		for _, args := range args {
			num, err := strconv.Atoi(args)

			if err != nil {
				fmt.Println(err)
			}
			sum = sum + num
		}
		fmt.Println("Resultnant addition is", sum)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
