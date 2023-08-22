/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package cmd

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/justjcurtis/unq/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unq",
	Short: "unq is a cli tool for creating and analyzing unique sets of strings",
	Long: `unq is a cli tool for creating and analyzing unique sets of strings

It can be used to create a unique set of strings from a file or stdin
or to compare two sets of strings to find the unique strings between them.
unq can be used in many ways to analyze and manage sets of strings.`,
	Run: func(cmd *cobra.Command, args []string) {
		showStats, _ := cmd.Flags().GetBool("stats")
		showCount, _ := cmd.Flags().GetBool("count")
		showPercent, _ := cmd.Flags().GetBool("percent")
		lines := utils.GetStdIn()
		length := len(lines)
		delimiter, err := utils.GetDelimiter(lines)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if length == 1 {
			lines = strings.Split(lines[0], delimiter)
			length = len(lines)
		} else {
			for i, line := range lines {
				if strings.HasSuffix(line, delimiter) {
					lines[i] = strings.TrimSuffix(line, delimiter)
				}
			}
			delimiter += "\n"
		}
		unique, m := utils.GetUnique(lines)
		if showStats {
			fmt.Println("Total Lines:", len(lines))
			fmt.Println("Total Unique Lines:", len(unique))
			fmt.Println("Total Duplicate Lines:", len(lines)-len(unique))
		}
		if showCount || showPercent {
			for k, count := range m {
				percentage := math.Round(float64(count)/float64(length)*10000) / 100
				if showCount && showPercent {
					fmt.Printf("%v: %v %v%%\n", k, count, percentage)
				} else if showCount {
					fmt.Printf("%v: %v\n", k, count)
				} else {
					fmt.Printf("%v: %v%%\n", k, percentage)
				}
			}
		} else {
			fmt.Println(strings.Join(unique, delimiter))
		}
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
	rootCmd.Flags().BoolP("stats", "s", false, "output stats about the unique set")
	rootCmd.Flags().BoolP("count", "c", false, "output count of entires in the unique set")
	rootCmd.Flags().BoolP("percent", "p", false, "output percentage of entires in the unique set")
}
