# dcmscan

A quick and simple DICOM collection scanner.

## Usage

    dcmscan [DIRECTORY]

dcmscan will scan all files in the given directory. If no directory is
specified the current directory will be used. Once all files are scanned the
unique series numbers and descriptions will be printed.

## Todo

* Handle multiple sessions
* JSON output
* Format output as a table
* Read more headers
* Verbose mode

## Thanks

* [suyashkumar](https://github.com/suyashkumar/dicom) for the go DICOM package
* [bmatcuk](https://github.com/bmatcuk/doublestar) for the doublestar package
