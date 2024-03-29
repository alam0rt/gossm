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
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/spf13/cobra"
)

var Path string
var Decrypt bool
var Recurse bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(args)
		c := cmd.Flag("recurse").Value
		fmt.Print(c)
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		svc := ssm.New(sess)

		p, err := svc.GetParametersByPath(&ssm.GetParametersByPathInput{
			Path:           &args[0],
			Recursive:      cmd.Flag("recurse").Value.Type(),
			WithDecryption: cmd.Flag("decrypt").Value.String(),
		})

		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(p)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().String("path", "p", "path to a parameter")
	getCmd.Flags().BoolP("decrypt", "d", false, "try to decrypt a secure string")
	getCmd.Flags().BoolP("recurse", "r", false, "be recursive")
}
