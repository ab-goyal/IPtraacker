/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:   "mul",
	Short: "multiplication of numbers",
	Long:  `multiplication of numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		mul := 1
		for _, args := range args {
			num, err := strconv.Atoi(args)

			if err != nil {
				fmt.Println(err)
			}
			mul = mul * num
		}
		fmt.Println("Resultnant multiplication is", mul)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// multiplyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// multiplyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
