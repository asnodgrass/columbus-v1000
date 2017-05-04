// Copyright © 2017 Adam Snodgrass <asnodgrass@sarchasm.us>
//
// This file is part of columbus-v1000.
//
// columbus-v1000 is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// columbus-v1000 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with columbus-v1000. If not, see <http://www.gnu.org/licenses/>.
//

package cmd

import (
	"os"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/asnodgrass/columbus-v1000/v1000"
)

// csvCmd represents the csv command
var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Converts to CSV format",
	Long: `Converts a Columbus V1000 GPS file to CSV format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inFile == "" {
			fmt.Println("error: input file required")
			return
		}

		file, err := os.Open(inFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		var out *os.File

		if outFile != "" {
			out, err = os.Open(outFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer out.Close()
		} else {
			out = os.Stdout
		}

		ok, err := v1000.CheckHeader(file)
		if err != nil || !ok {
			fmt.Println(err)
			return
		}

		printHeader(out)
		for {
			rec, err := v1000.ParseRecord(file)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			printRow(&rec, out)
		}
	},
}

func init() {
	RootCmd.AddCommand(csvCmd)
	csvCmd.Flags().StringVarP(&inFile, "in-file", "i", "", "input file (required)")
	csvCmd.Flags().StringVarP(&outFile, "out-file", "o", "", "output file")
}

func printHeader(out *os.File) {
	fields := []string{
		"INDEX",
		"TAG",
		"DATE",
		"TIME",
		"LATITUDE N/S",
		"LONGITUDE E/W",
		"HEIGHT",
		"SPEED",
		"HEADING",
		"PRES",
		"TEMP",
	}
	hdr := strings.Join(fields, ",")
	fmt.Fprintln(out, hdr)
}

func printRow(rec *v1000.Record, out *os.File) {
	fields := []string{
		fmt.Sprintf("%d", rec.Index),
		rec.Type,
		formatDateCSV(rec.Time),
		formatTimeCSV(rec.Time),
		fmt.Sprintf("%f", rec.Latitude),
		fmt.Sprintf("%f", rec.Longitude),
		fmt.Sprintf("%d", rec.Altitude),
		fmt.Sprintf("%.1f", rec.Speed),
		fmt.Sprintf("%d", rec.Heading),
		fmt.Sprintf("%.1f", rec.Pressure),
		fmt.Sprintf("%d", rec.Temperature),
	}
	row := strings.Join(fields, ",")
	fmt.Fprintln(out, row)
}

func formatDateCSV(date v1000.Date) string {
	return fmt.Sprintf("%02d%02d%02d", date.Year - 2000, date.Month, date.Day)
}

func formatTimeCSV(date v1000.Date) string {
	return fmt.Sprintf("%02d%02d%02d", date.Hour, date.Minute, date.Second)
}
