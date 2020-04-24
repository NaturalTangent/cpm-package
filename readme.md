# cpm-package

A cross platform, command line tool to generate package files compliant with Grant Searle's CP/M computers (including derivatives such as RC2014).

__Only tested on linux so far!__

## Building

(There's a debian linux x64 binary in the binaries directory).

Make sure you have the go language tools installed (https://golang.org/doc/install).

Clone this repo (or download the single .go file if you want);

Build using;


```shell
go build cpm-package.go 
```

## Usage

This tool allows files to be packaged in preparation for transferring to a CP/M machine. For details of how to transfer this package file, see http://www.searle.wales.

_Note: I had to add a 10ms delay between characters when sending the file - using cutecom on linux._


### Command Line Parameters

```shell
cpm-package [-o outputfile] [-u user] [-r receiving-exe] binary1 binary2 ... binaryX
```

Where;
* _outputfile_  is the resulting package file. If not specified _stdout_ is used.
* _user_ is the cpm user. If not specified _U0_ is used.
* _receiving-exe_ is the cp/m executable use to process the data on the cp/m machine. If not specified _A:DOWNLOAD_ is used.

The list of files to be packaged must be specified last - wildcards are supported.


```shell
./cpm-package -o output.txt BBCBASIC.COM *.BBC
```
