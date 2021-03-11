/*
 *  mcAsm - a minecraft Assembler
 *     Copyright (C) 2021 OmegaRogue
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"mcAsm/internal/app"
)

var AsmFilePath string
var OutPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mcAsm",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		asmData, _ := ioutil.ReadFile(AsmFilePath) // ignoring errors for brevity
		if OutPath == "" {
			OutPath = strings.TrimSuffix(AsmFilePath, filepath.Ext(AsmFilePath)) + ".hack"
		}

		var p app.Parser
		p.Init(asmData)
		hackFile := p.Parse()

		var b bytes.Buffer
		for _, i := range hackFile.Instructions {
			b.WriteString(i.BinaryString())
		}
		err := ioutil.WriteFile(OutPath, b.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("error on write output: %w", err)
		}
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVarP(&AsmFilePath, "file", "f", "", "source file")
	rootCmd.MarkFlagRequired("file")
	rootCmd.Flags().StringVarP(&OutPath, "out", "o", "", "output file")
}
