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

package v1000

import (
  "testing"
)

func Test_parseV1000Date(t *testing.T) {
  expected := Date{
    Year: 2017,
    Month: 4,
    Day: 1,
    Hour: 12,
    Minute: 34,
    Second: 56,
  }
  data := uint32(84068536)
  t.Logf("Checking whether parseV1000Date() correctly parses %d", data)
  out := parseV1000Date(data)
  if out.Year != expected.Year {
    t.Errorf("Expected Year %d, got %d", expected.Year, out.Year)
  }
  if out.Month != expected.Month {
    t.Errorf("Expected Month %d, got %d", expected.Month, out.Month)
  }
  if out.Day != expected.Day {
    t.Errorf("Expected Day %d, got %d", expected.Day, out.Day)
  }
  if out.Hour != expected.Hour {
    t.Errorf("Expected Hour %d, got %d", expected.Hour, out.Hour)
  }
  if out.Minute != expected.Minute {
    t.Errorf("Expected Minute %d, got %d", expected.Minute, out.Minute)
  }
  if out.Second != expected.Second {
    t.Errorf("Expected Second %d, got %d", expected.Second, out.Second)
  }
}

func Test_hasBit(t *testing.T) {
  b := byte(42)
  t.Logf("Checking whether hasBit() can determine each bit value in %d", b)
  bits := [8]bool{false, true, false, true, false, true, false, false}
  for idx, expected := range bits {
    if out := hasBit(b, uint(idx)); out != expected {
      t.Errorf("Expected %v at position %d, but got %v", expected, idx, out)
    }
  }
}

func Test_shift(t *testing.T) {
  value := uint32(42)
  shifter := uint32(0xf)
  expected := uint32(10)
  t.Logf("Checking whether shift() properly bitshifts %d by %d", value, shifter)
  if out := shift(value, shifter); out != expected {
    t.Errorf("Expected %d, but got %d", expected, out)
  }
}
