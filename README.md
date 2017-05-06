# columbus-v1000

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
[gpxbabel][] can be used to convert from GPX into many other possible formats.

## Installation

Docs coming soon

## Contributing

Pull requests are welcomed.

## License

[GNU General Public License v3](LICENSE)
