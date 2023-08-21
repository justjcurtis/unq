/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/justjcurtis/unq/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unq",
	Short: "a cli tool for creating and analyzing unique sets of strings",
	Long: `It can be used to create a unique set of strings from a file or stdin
or to compare two sets of strings to find the unique strings between them.
unq can be used in many ways to analyze and manage sets of strings.`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := utils.GetStdIn()
		delimiter, err := utils.GetDelimiter(lines)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		println("\"" + delimiter + "\"")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.unq.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
