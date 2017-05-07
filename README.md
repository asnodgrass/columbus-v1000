# columbus-v1000

[![Build Status](https://travis-ci.org/asnodgrass/columbus-v1000.svg?branch=master)](https://travis-ci.org/asnodgrass/columbus-v1000)

columbus-v1000 is a command line tool authored in Go to convert the binary
`.gps` files produced by a Columbus V1000 into CSV or GPX formats.

## Background

The V1000 is a GPS logging device with approximately 16MiB of internal storage,
and can store logs in one of three formats:

* Standard GPX format
  * does not include any temperature or barometric pressure data
  * most verbose, maximizes device storage consumption
* Comma separated values (CSV)
  * includes each datum produced by the device
  * moderate storage consumption
* Proprietary binary format
  * includes each datum produced by the device
  * 28 byte records, minimizes device storage consumption

GPX is the most widely used format, but also the most verbose; the binary format
is the best choice to get the most out of the limited device storage.

Use this tool to convert binary format files into your choice of CSV or GPX.
[gpsbabel][] can be used to convert from GPX into many other possible formats.

## Installation

Assuming Go has been installed locally:

    go get github.com/asnodgrass/columbus-v1000
    go install github.com/asnodgrass/columbus-v1000

## Usage

The syntax and arguments to convert to either GPX or CSV are nearly identical:

    Converts a Columbus V1000 GPS file to CSV format.

    Usage:
      columbus-v1000 csv [flags]

    Flags:
      -i, --in-file string    input file (required)
      -o, --out-file string   output file

    Global Flags:
      -z, --timezone string   Timezone for input file (default: UTC)

For GPX conversion, use `gpx` rather than `csv`.

The `--out-file` flag can be omitted, in which case the result will be sent to
stdout.

## Contributing

There are likely many things that can be improved here. Pull requests are
welcomed!

## License

[GNU General Public License v3](LICENSE)

[gpsbabel]: <https://www.gpsbabel.org/>
