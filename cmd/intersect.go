/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/justjcurtis/unq/utils"
	"github.com/spf13/cobra"
)

// intersectCmd represents the intersect command
var intersectCmd = &cobra.Command{
	Use:   "intersect [file1] [file2?]",
	Short: "unq is a cli tool for creating and analyzing unique sets of strings",
	Long: `unq is a cli tool for creating and analyzing unique sets of strings

It can be used to create a unique set of strings from a file or stdin
or to compare two sets of strings to find the unique strings between them.
unq can be used in many ways to analyze and manage sets of strings.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get flags
		showCount, _ := cmd.Flags().GetBool("count")
		delimiterIn, _ := cmd.Flags().GetString("delemiter-in")
		delimiterOut, _ := cmd.Flags().GetString("delemiter-out")
		invert, _ := cmd.Flags().GetBool("invert")
		order, _ := cmd.Flags().GetBool("order")
		orderAz, _ := cmd.Flags().GetBool("order-az")
		showPercent, _ := cmd.Flags().GetBool("percent")
		reverse, _ := cmd.Flags().GetBool("reverse")
		showStats, _ := cmd.Flags().GetBool("stats")
		trim, _ := cmd.Flags().GetBool("trim")
		trim = !trim

		// get files from args
		if len(args) < 1 {
			fmt.Println("Please provide at least one file")
			os.Exit(1)
		}
		file1 := args[0]

		file2Contents := ""
		// get input from stdin
		input := utils.GetStdIn(trim)
		if len(input) > 0 {
			file2Contents = input
		} else {
			if len(args) < 2 {
				fmt.Println("Please provide a second file or input from stdin")
				os.Exit(1)
			}
			file2 := args[1]
			// get input from file2
			contents, err := utils.GetFileContents(file2, trim)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			file2Contents = contents
		}

		// get input from file1
		file1Contents, err := utils.GetFileContents(file1, trim)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		delimiter1 := ""
		delimiter2 := ""
		// get delimiter
		if delimiterIn == "" {
			smartDelimiter1, err := utils.GetDelimiter(file1Contents)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			smartDelimiter2, err := utils.GetDelimiter(file2Contents)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			delimiter1 = smartDelimiter1
			delimiter2 = smartDelimiter2
			delimiterIn = smartDelimiter1
		} else {
			delimiterIn = strings.ReplaceAll(delimiterIn, "\\n", "\n")
			delimiterIn = strings.ReplaceAll(delimiterIn, "\\t", "\t")
			delimiter1 = delimiterIn
			delimiter2 = delimiterIn
		}
		if delimiterOut == "" {
			delimiterOut = delimiterIn
		} else {
			delimiterOut = strings.ReplaceAll(delimiterOut, "\\n", "\n")
			delimiterOut = strings.ReplaceAll(delimiterOut, "\\t", "\t")
		}

		// split file1Contents
		entries1 := strings.Split(file1Contents, delimiter1)
		entries2 := strings.Split(file2Contents, delimiter2)

		// get unique sets
		unique1, m1 := utils.GetUnique(entries1)
		unique2, m2 := utils.GetUnique(entries2)

		discarded := []string{}
		d := make(map[string]int)
		// get intersect set
		m := make(map[string]int)
		unique := []string{}
		for _, k := range unique1 {
			if _, ok := m2[k]; ok {
				m[k] = m1[k] + m2[k]
				unique = append(unique, k)
			} else {
				discarded = append(discarded, k)
				if _, ok := d[k]; ok {
					d[k]++
				} else {
					d[k] = 1
				}
			}
		}
		for _, k := range unique2 {
			if _, ok := m1[k]; !ok {
				discarded = append(discarded, k)
				if _, ok := d[k]; ok {
					d[k]++
				} else {
					d[k] = 1
				}
			}
		}

		// handle inversion
		if invert {
			temp := unique
			unique = discarded
			discarded = temp
			t := m
			m = d
			d = t
		}

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

		length := len(entries1) + len(entries2)
		uniqueLength := len(unique)
		// output stats
		if showStats {
			fmt.Println("Entries:", length)
			fmt.Println("Unique entries:", uniqueLength)
			fmt.Println("Duplicate entries:", length-uniqueLength)
			fmt.Println("Discarded entries:", len(discarded))
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

func init() {
	rootCmd.AddCommand(intersectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// intersectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	intersectCmd.Flags().BoolP("count", "c", false, "output count of entires in the unique set")
	intersectCmd.Flags().StringP("delemiter-in", "d", "", "delimiter to use when splitting input")
	intersectCmd.Flags().StringP("delemiter-out", "D", "", "delimiter to use when outputting unique set")
	intersectCmd.Flags().BoolP("invert", "i", false, "invert the intersection to get the difference")
	intersectCmd.Flags().BoolP("order", "o", false, "order the unique set by count")
	intersectCmd.Flags().BoolP("order-az", "O", false, "order the unique set alphabetically")
	intersectCmd.Flags().BoolP("percent", "p", false, "output percentage of entires in the unique set")
	intersectCmd.Flags().BoolP("reverse", "r", false, "reverse the order of the unique set")
	intersectCmd.Flags().BoolP("stats", "s", false, "output stats about the unique set")
	intersectCmd.Flags().BoolP("trim", "t", false, "trim whitespace from entries (true by default)")
}
