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

func Test_formatDateCSV(t *testing.T) {
  expected := "170401"
  t.Logf("Checking for correct CSV date format.. (expected: %s)", expected)
  date := v1000.Date{
    Year: 2017,
    Month: 4,
    Day: 1,
    Hour: 12,
    Minute: 34,
    Second: 56,
  }
  if out := formatDateCSV(date); out != expected {
    t.Errorf("Expected '%s', but got '%s' instead", expected, out)
  }
}

func Test_formatTimeCSV(t *testing.T) {
  expected := "123456"
  t.Logf("Checking for correct CSV time format.. (expected: %s)", expected)
  date := v1000.Date{
    Year: 2017,
    Month: 4,
    Day: 1,
    Hour: 12,
    Minute: 34,
    Second: 56,
  }

  if out := formatTimeCSV(date); out != expected {
    t.Errorf("Expected '%s', but got '%s' instead", expected, out)
  }
}

func Test_formatLatLon(t *testing.T) {
  expected := "24.123456N"
  t.Logf("Checking for correctly formatted Lat/Lon.. (expected: %s)", expected)
  latLon := 24.123456
  northSouth := true

  if out := formatLatLon(latLon, northSouth); out != expected {
    t.Errorf("Expected '%s', but got '%s' instead", expected, out)
  }
}
