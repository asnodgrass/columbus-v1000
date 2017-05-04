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
  "log"
  "time"
	"io"
	"io/ioutil"
  "encoding/xml"

	"github.com/spf13/cobra"
	"github.com/asnodgrass/columbus-v1000/v1000"
)

// gpxCmd represents the gpx command
var gpxCmd = &cobra.Command{
	Use:   "gpx",
	Short: "Converts to GPX format",
	Long: `Converts a Columbus V1000 GPS file to GPX format.`,
	Run: func(cmd *cobra.Command, args []string) {
		var trkpts []trackPoint

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

		ok, err := v1000.CheckHeader(file)
		if err != nil || !ok {
			fmt.Println(err)
			return
		}

		for {
			rec, err := v1000.ParseRecord(file)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			trkpts = append(trkpts, recordToTrackPoint(rec))
		}

		gpxData := generateGPX(trkpts)
		if outFile != "" {
			ioutil.WriteFile(outFile, gpxData, 0644)
		} else {
			fmt.Println(string(gpxData[:]))
		}
	},
}

func init() {
	RootCmd.AddCommand(gpxCmd)
	gpxCmd.Flags().StringVarP(&inFile, "in-file", "i", "", "input file (required)")
	gpxCmd.Flags().StringVarP(&outFile, "out-file", "o", "", "output file")
}

// GPXTrackPoint ...
type trackPoint struct {
  XMLName xml.Name `xml:"trkpt"`
  Latitude float32 `xml:"lat,attr"`
  Longitude float32 `xml:"lon,attr"`
  Time string `xml:"time"`
  Altitude uint32 `xml:"ele"`
  Speed float32 `xml:"speed"`
  Heading uint16 `xml:"course"`
}

type trackSegment struct {
  XMLName xml.Name `xml:"trkseg"`
  TrackPoints []trackPoint `xml:"trkpt"`
}

type track struct {
  XMLName xml.Name `xml:"trk"`
  Name string `xml:"name"`
  Description string `xml:"desc"`
  TrackSegments []trackSegment `xml:"trkseg"`
}

type header struct {
  XMLName xml.Name `xml:"gpx"`
  Version string `xml:"version,attr"`
  Creator string `xml:"creator,attr"`
  Namespace string `xml:"xmlns,attr"`
  Time string `xml:"time"`
	Bounds string `xml:"bounds"`
  Track track `xml:"trk"`
}

func recordToTrackPoint(rec v1000.Record) (trackPoint) {
  tp := trackPoint{
    Latitude: rec.Latitude,
    Longitude: rec.Longitude,
    Altitude: rec.Altitude,
    Speed: rec.Speed,
    Heading: rec.Heading,
    Time: formatDateRFC3339(rec.Time, timeZone),
  }
  return tp
}

func generateGPX(trkpts []trackPoint) []byte {
  trksegs := make([]trackSegment, 1)
  trksegs[0].TrackPoints = trkpts
  trk := track{TrackSegments: trksegs}
  trk.Name = "please fix me"
  trk.Description = "fix this too"
  gpx := header{Track: trk}
  gpx.Version = "1.0"
  gpx.Namespace = "http://www.topografix.com/GPX/1/0"
  gpx.Creator = "columbus-v1000"
  gpx.Time = time.Now().UTC().Format(time.RFC3339)
  body, err := xml.MarshalIndent(gpx, "", "  ")
  if err != nil {
    log.Fatal(err)
  }
	out := []byte(xml.Header)
	out = append(out, body...)
  return out
}

func formatDateRFC3339(date v1000.Date, zone string) (string) {
  yr  := int(date.Year)
  mon := time.Month(date.Month)
  day := int(date.Day)
  hr  := int(date.Hour)
  min := int(date.Minute)
  sec := int(date.Second)
  tz, err := time.LoadLocation(zone)
  if err != nil {
    fmt.Println(err)
    return ""
  }
  t := time.Date(yr, mon, day, hr, min, sec, 0, tz)
  return t.Format(time.RFC3339)
}
