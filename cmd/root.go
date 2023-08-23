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
	"slices"
)

var rootCmd = &cobra.Command{
	Use:   "unq",
	Short: "unq is a cli tool for creating and analyzing unique sets of strings",
	Long: `unq is a cli tool for creating and analyzing unique sets of strings

It can be used to create a unique set of strings from a file or stdin
or to compare two sets of strings to find the unique strings between them.
unq can be used in many ways to analyze and manage sets of strings.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get flags
		showStats, _ := cmd.Flags().GetBool("stats")
		showCount, _ := cmd.Flags().GetBool("count")
		showPercent, _ := cmd.Flags().GetBool("percent")
		order, _ := cmd.Flags().GetBool("order")
		orderAz, _ := cmd.Flags().GetBool("order-az")
		reverse, _ := cmd.Flags().GetBool("reverse")
		trim, _ := cmd.Flags().GetBool("trim")
		trim = !trim
		delimiterIn, _ := cmd.Flags().GetString("delemiter-in")
		delimiterOut, _ := cmd.Flags().GetString("delemiter-out")

		// get input
		input := utils.GetStdIn(trim)

		// get delimiter
		if delimiterIn == "" {
			smartDelimiter, err := utils.GetDelimiter(input)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			delimiterIn = smartDelimiter
		} else {
			delimiterIn = strings.ReplaceAll(delimiterIn, "\\n", "\n")
			delimiterIn = strings.ReplaceAll(delimiterIn, "\\t", "\t")
		}
		if delimiterOut == "" {
			delimiterOut = delimiterIn
		} else {
			delimiterOut = strings.ReplaceAll(delimiterOut, "\\n", "\n")
			delimiterOut = strings.ReplaceAll(delimiterOut, "\\t", "\t")
		}

		// split lines
		entries := strings.Split(input, delimiterIn)
		length := len(entries)

		// get unique set
		unique, m := utils.GetUnique(entries)

		// sort unique set
		if order || orderAz {
			if orderAz {
				slices.Sort(unique)
			} else {
				slices.SortFunc(unique, func(a string, b string) int {
					return m[a] - m[b]
				})
			}
		}

		// reverse unique set
		if reverse {
			slices.Reverse(unique)
		}

		// output stats
		if showStats {
			fmt.Println("Entries:", len(entries))
			fmt.Println("Unique entries:", len(unique))
			fmt.Println("Duplicate entries:", len(entries)-len(unique))
		}

		// output unique set
		if showCount || showPercent {
			for _, k := range unique {
				count := m[k]
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
			fmt.Println(strings.Join(unique, delimiterOut))
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
	rootCmd.Flags().BoolP("order", "o", false, "order the unique set by count")
	rootCmd.Flags().BoolP("order-az", "O", false, "order the unique set alphabetically")
	rootCmd.Flags().BoolP("reverse", "r", false, "reverse the order of the unique set")
	rootCmd.Flags().BoolP("trim", "t", false, "trim whitespace from entries (true by default)")
	rootCmd.Flags().StringP("delemiter-in", "d", "", "delimiter to use when splitting input")
	rootCmd.Flags().StringP("delemiter-out", "D", "", "delimiter to use when outputting unique set")
}
