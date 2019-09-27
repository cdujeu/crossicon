package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/cdujeu/crossicon/lib"
	"github.com/spf13/cobra"
)

var (
	pngInput     string
	outputPrefix string
	packageName  string
	writeICO     bool
	writeBytes   bool
)

// RootCmd is the parent of all commands defined in this package.
var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Tool for converting PNG to bytes array",
	Run: func(cmd *cobra.Command, args []string) {
		if pngInput == "" || outputPrefix == "" || packageName == "" {
			cmd.Help()
			log.Fatal("Missing Parameters")
		}
		if !writeBytes && !writeICO {
			cmd.Help()
			log.Fatal("Please select either --ico or --bytes or both")
		}

		f, e := os.Open(pngInput)
		if e != nil {
			log.Fatal("Cannot open input file: ", e)
		}
		defer f.Close()
		// UNIX VERSION (PNG)
		if writeBytes {
			unixOutput := outputPrefix + "unix.go"
			os.Truncate(unixOutput, 0)
			out, e := os.OpenFile(unixOutput, os.O_CREATE|os.O_WRONLY, 0755)
			if e != nil {
				log.Fatal("Cannot open unix output file: ", e)
			}
			defer out.Close()
			err := lib.AsBytesArray(out, packageName, "linux darwin", f)
			if err != nil {
				log.Fatal("Cannot PNG write as bytes: ", err)
			} else {
				log.Println("[--bytes] PNG version written as bytes array to " + unixOutput)
			}
		}

		// WINDOWS VERSION (ICO)
		bb := bytes.NewBuffer([]byte{})
		// Rewind file
		f.Seek(0, 0)
		if err := lib.ConvertToIco(f, bb); err != nil {
			log.Fatal("Cannot convert PNG to ICO: ", err)
		}
		if writeICO {
			// Write to icon.ico
			ico := outputPrefix + ".ico"
			os.Truncate(ico, 0)
			if err := ioutil.WriteFile(ico, bb.Bytes(), 0755); err != nil {
				log.Println("[Error] Cannot write ICO file", err.Error())
			} else {
				log.Println("[--ico] ICO file written to " + ico)
			}
		}
		if writeBytes {
			// Write to iconwin.go
			winOutput := outputPrefix + "win.go"
			os.Truncate(winOutput, 0)
			wOut, e := os.OpenFile(winOutput, os.O_CREATE|os.O_WRONLY, 0755)
			if e != nil {
				log.Fatal("Cannot open windows output file: ", e)
			}
			defer wOut.Close()
			err := lib.AsBytesArray(wOut, packageName, "windows", bb)
			if err != nil {
				log.Fatal("Cannot ICO write as bytes: ", err)
			} else {
				log.Println("[--bytes] ICO version written as bytes array to " + winOutput)
			}
		}

	},
}

func init() {
	RootCmd.Flags().BoolVar(&writeICO, "ico", false, "Write ICO to file")
	RootCmd.Flags().BoolVar(&writeBytes, "bytes", false, "Write to bytes files for both unix and windows")

	RootCmd.Flags().StringVarP(&pngInput, "input", "i", "", "Path to PNG input file")
	RootCmd.Flags().StringVarP(&outputPrefix, "output", "o", "", "Prefix for writing output files")
	RootCmd.Flags().StringVarP(&packageName, "package", "p", "", "Package name")
}
