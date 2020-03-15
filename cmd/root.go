// Copyright © 2019 Erin Shepherd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	//_ "github.com/erincandescent/nuvoprog/target/all"
	_ "../target/all"
)

var cfgFile string
var verbose bool
var targetName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nuvoprog",
	Short: "Nuvoton device programmer ** experimental MS51FB version **",
	Long: `A tool for programming Nuvoton devices, particularly
	focusing on their modern 8051 family ** experimental MS51FB version **
	
Notes: 
Possible targets are Novoton N76E003 and MS51FB9AE. 
The programmer utility has been tested with Nu-Link-Me on a NT-MS51FB 8051 NuTiny dev board.

The Program memory is limited to 12KB for both the N76E003 and MS51FB9AE processors because the image split command does not 
parse the chip configuration to determine the split between the program flash and load flash memory, 
So it defaults to the worst case of 4KB of load flash. (The N76E003 might be able to be upped to 14KB, but it was set to 
12KB by original author).

If you are going to rebuild this program:
The include paths are setup as relative, so this program source code should be copied to your system 
(git clone of zip file - https://github.com/mountaintom/nuvoprog.git - at this time the latest version may be in one of the branches) and 
compiled (go build) in-place on your computer. The nuvoprog command should be run from there or manually moved to where you want it. 

You can compile the nuvoprog utility as another name (such as nuvoprog-test) by changing the main directory name (such as nuvoprog to nuvoprog-test) then run go build.

Examples:
	Download flash data from chip:
	./nuvoprog read ./flash-read.ihx --target MS51FB9AE 

	Split downloaded flash data into Program, Load ROM and chip configuration files:
	./nuvoprog image split -i ./flash-read.ihx --target MS51FB9AE  -a program-flash-data.ihx -l loader-flas-data.ihx -c chip-configuration.json 
	Note: This is how to get an example chip-config json file to work with.

	Program flash in chip:
	./nuvoprog program --target ms51fb9ae -a ./program-to-flash.ihx -c @chip-configuration.json
        Note: The files may be combined with "image merge" and the resulting ihx file programmed with the "-i" flag. 
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !verbose {
			log.SetOutput(ioutil.Discard)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "make verbose (enable debug logging)")
	rootCmd.PersistentFlags().StringVarP(&targetName, "target", "t", "", "target device")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
