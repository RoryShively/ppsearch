/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/RoryShively/ppsearch/engine"
	"github.com/spf13/cobra"
)

// constructCmd represents the construct command
var constructCmd = &cobra.Command{
	Use:   "construct",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a url to construct index")
		}
		if len(args) > 1 {
			return errors.New("too many urls entered into search")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Get DB
		db := engine.GetDB()
		// db.Clear()

		// Parse page
		url := args[0]
		if !strings.HasPrefix(url, "http") {
			url = fmt.Sprintf("http://%v", url)
		}
		idx := engine.NewIndexer()
		idx.Start(url)

		// Save data
		db.SaveIndex(idx)

	},
}

func init() {
	rootCmd.AddCommand(constructCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// constructCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// constructCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
