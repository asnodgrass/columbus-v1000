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
  "testing"

  "github.com/asnodgrass/columbus-v1000/v1000"
)

func Test_recordToTrackPoint(t *testing.T) {
  t.Log("Checking for valid trackPoint generation..")
  data := v1000.Record{
    Index: 0,
    Type: "T",
    Time: v1000.Date{
      Year: 2017,
      Month: 4,
      Day: 1,
      Hour: 12,
      Minute: 34,
      Second: 56,
    },
    Latitude: -34.987654,
    South: true,
    Longitude: 99.123456,
    West: false,
    Altitude: 10,
    Speed: 1,
    Heading: 180,
    Pressure: 1000,
    Temperature: 20,
  }
  expected := trackPoint{
    Latitude: -34.987654,
    Longitude: 99.123456,
    Altitude: 10,
    Speed: 1,
    Heading: 180,
    Time: "2017-04-01T12:34:56Z",
  }

  out := recordToTrackPoint(data)
  if out.Latitude != expected.Latitude {
    t.Errorf("Expected Latitude %.6f, got %.6f instead", expected.Latitude, out.Latitude)
  }
  if out.Longitude != expected.Longitude {
    t.Errorf("Expected Longitude %.6f, got %.6f instead", expected.Longitude, out.Longitude)
  }
  if out.Altitude != expected.Altitude {
    t.Errorf("Expected Altitude %d, got %d instead", expected.Altitude, out.Altitude)
  }
  if out.Speed != expected.Speed {
    t.Errorf("Expected Speed %.6f, got %.6f instead", expected.Speed, out.Speed)
  }
  if out.Heading != expected.Heading {
    t.Errorf("Expected Heading %d, got %d instead", expected.Heading, out.Heading)
  }
  if out.Time != expected.Time {
    t.Errorf("Expected Time '%s', got '%d' instead", expected.Time, out.Time)
  }
}

// FIXME: deal with dynamic content
func Test_generateGPX(t *testing.T) {
}

func Test_formatDateRFC3339(t *testing.T) {
  expected := "2017-04-01T12:34:56Z"
  t.Logf("Checking for valid RFC3339 output.. (expected: %s)", expected)
  data := v1000.Date{
    Year: 2017,
    Month: 4,
    Day: 1,
    Hour: 12,
    Minute: 34,
    Second: 56,
  }
  if out := formatDateRFC3339(data, "UTC"); out != expected {
    t.Errorf("Expected '%s', but got '%s'", expected, out)
  }
}

func Test_minimumLatitude(t *testing.T) {
  expected := latLong(-45)
  t.Logf("Checking for valid minimum latitude.. (expected: %f)", expected)
  data := make([]trackPoint, 2)
  data[0].Latitude = 45
  data[1].Latitude = -45
  if out := minimumLatitude(data); out != expected {
    t.Errorf("Expected %f, but got %f", expected, out)
  }
}

func Test_minimumLongitude(t *testing.T) {
  expected := latLong(-90)
  t.Logf("Checking for valid minimum longitude.. (expected: %f)", expected)
  data := make([]trackPoint, 2)
  data[0].Longitude = 90
  data[1].Longitude = -90
  if out := minimumLongitude(data); out != expected {
    t.Errorf("Expected %f, but got %f", expected, out)
  }
}

func Test_maximumLatitude(t *testing.T) {
  expected := latLong(45)
  t.Logf("Checking for valid maximum latitude.. (expected: %f)", expected)
  data := make([]trackPoint, 2)
  data[0].Latitude = 45
  data[1].Latitude = -45
  if out := maximumLatitude(data); out != expected {
    t.Errorf("Expected %f, but got %f", expected, out)
  }
}

func Test_maximumLongitude(t *testing.T) {
  expected := latLong(90)
  t.Logf("Checking for valid maximum longitude.. (expected: %f)", expected)
  data := make([]trackPoint, 2)
  data[0].Longitude = 90
  data[1].Longitude = -90
  if out := maximumLongitude(data); out != expected {
    t.Errorf("Expected %f, but got %f", expected, out)
  }
}
