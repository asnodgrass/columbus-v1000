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
  "os"
  "encoding/binary"
)

// Record is the contents of a record in a V1000 GPS file
type Record struct {
  Index uint32
  Type string
  Time Date
  Longitude float64
  South bool
  Latitude float64
  West bool
  Altitude uint32
  Speed float64
  Heading uint16
  Pressure float64
  Temperature uint16
}

// Date is a structure to hold a date/time stamp
type Date struct {
  Year uint32
  Month uint32
  Day uint32
  Hour uint32
  Minute uint32
  Second uint32
}

// CheckHeader ...
func CheckHeader(file *os.File) (bool, error) {
  hdr, err := parseShort(file)
  if err != nil {
    return false, err
  }
  return hdr == 1799, nil
}

// ParseRecord ...
func ParseRecord(file *os.File) (Record, error) {
  var rec Record

  err := parseIndex(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parseByte3(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parseTime(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parseCoords(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parseAltitude(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parseSpeed(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parseHeading(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parsePressure(file, &rec)
  if err != nil {
    return rec, err
  }
  err = parseTemperature(file, &rec)
  if err != nil {
    return rec, err
  }
  return rec, nil
}

func parseShort(file *os.File) (uint16, error) {
  var value uint16
  err := binary.Read(file, binary.BigEndian, &value)
  if err != nil {
    return 0, err
  }
  return value, nil
}

func parseLong(file *os.File) (uint32, error) {
  var value uint32
  err := binary.Read(file, binary.BigEndian, &value)
  if err != nil {
    return 0, err
  }
  return value, nil
}

func parseIndex(file *os.File, rec *Record) (error) {
  data := make([]byte, 3)
  _, err := file.Read(data)
  if err != nil {
    return err
  }
  rec.Index = uint32(data[2]) | uint32(data[1])<<8 | uint32(data[0])<<16
  return nil
}

// FIXME: rename this
func parseByte3(file *os.File, rec *Record) (error) {
  var byte3 byte
  err := binary.Read(file, binary.BigEndian, &byte3)
  if err != nil {
    return err
  }
  if hasBit(byte3, 0) && !hasBit(byte3, 1) {
    rec.Type = "P"
  } else {
    rec.Type = "T"
  }
  rec.South = hasBit(byte3, 2)
  rec.West = hasBit(byte3, 3)
  return nil
}

func parseTime(file *os.File, rec *Record) (error) {
  value, err := parseLong(file)
  if err != nil {
    return err
  }
  rec.Time = parseV1000Date(value)
  return nil
}

func parseCoords(file *os.File, rec *Record) (error) {
  value, err := parseLong(file)
  if err != nil {
    return err
  }
  rec.Latitude = float64(value) / 1000000.0
  if rec.South {
    rec.Latitude = -rec.Latitude
  }
  value, err = parseLong(file)
  if err != nil {
    return err
  }
  rec.Longitude = float64(value) / 1000000.0
  if rec.West {
    rec.Longitude = -rec.Longitude
  }
  return nil
}

func parseAltitude(file *os.File, rec *Record) error {
  value, err := parseLong(file)
  if err != nil {
    return err
  }
  rec.Altitude = value / 10
  return nil
}

func parseSpeed(file *os.File, rec *Record) error {
  value, err := parseShort(file)
  if err != nil {
    return err
  }
  rec.Speed = float64(value) / 10.0
  return nil
}

func parseHeading(file *os.File, rec *Record) error {
  value, err := parseShort(file)
  if err != nil {
    return err
  }
  rec.Heading = value
  return nil
}

func parsePressure(file *os.File, rec *Record) error {
  value, err := parseShort(file)
  if err != nil {
    return err
  }
  rec.Pressure = float64(value) / 10.0
  return nil
}

func parseTemperature(file *os.File, rec *Record) error {
  value, err := parseShort(file)
  if err != nil {
    return err
  }
  rec.Temperature = value / 10
  return nil
}

func parseV1000Date(value uint32) (Date) {
  var date Date
  date.Year   = (shift(value, uint32(0x3f << 26)) >> 26) + 2016
  date.Month  = (shift(value, uint32(0xf  << 22)) >> 22) //- 1
  date.Day    =  shift(value, uint32(0x1f << 17)) >> 17
  date.Hour   =  shift(value, uint32(0x1f << 12)) >> 12
  date.Minute =  shift(value, uint32(0x3f << 6))  >> 6
  date.Second =  shift(value, uint32(0x3f))
  return date
}

func hasBit(b byte, pos uint) bool {
  return (b >> pos & 0x1) == 0x1
}

func shift(value uint32, shifter uint32) (uint32) {
  shifted := value & shifter
  return shifted
}
