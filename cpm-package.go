/*
A command line tool to package files for copying to Grant Searle's CP/M.
See http://www.searle.wales/ for details (Windows packager).

Andy Anderson 24th April 2020
*/

package main

import (
	"fmt"
	"flag"
	"os"
	"path"
	"path/filepath"
	"io/ioutil"
	"encoding/hex"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
The checksum is two pairs of HEX values:
The first one is the low-byte of the length of the file being uploaded. Because CP/M files are normally saved in blocks of 128, this is normally "80" or "00", but doesn't need to be.
The second one is the low-byte of the SUM of each byte being uploaded.
*/
func checksum(data []byte) string {
	var sum = 0;

	for _, b := range data {
		sum += int(b)
	}

	sSum := fmt.Sprintf("%02X%02X", len(data) & 0xff, sum & 0xff)

	return sSum
}

func main() {
	outputPtr := flag.String("o", "", "output file")
	userPtr := flag.String("u", "U0", "user")
	receivePtr := flag.String("r", "A:DOWNLOAD", "cpm/m download executable")

	flag.Parse()

	var err error
	var fOutput = os.Stdout

	// If specified, open the output file. Otherwise use stdout.
	if *outputPtr != "" {
		fOutput, err = os.Create(*outputPtr)
		check(err) 
		defer fOutput.Close()
	}

	// Read each specified input file completely, and write into the output file.
	for _, filespec := range flag.Args() {

		matches, err := filepath.Glob(filespec)
		check(err)

		for _, f := range matches {
			var data []byte

			data, err = ioutil.ReadFile(f)
			check(err)

			_, fname := path.Split(f)

			/*
			First line:
				A:DOWNLOAD <filename> where <filename> is the CPM file that is required. Must be a space between DOWNLOAD and the filename.
			Second line:
				U0 This is the "user" number that can see the destination file. If you are not familiar with CP/M users, then always use U0
			Third line:
				A continual stream of HEX values (must be 2 chars each) for each byte in the file, followed by a > character, followed by the checksum
			*/
			fmt.Fprintf(fOutput, "%s %s\r\n", *receivePtr, fname)
			fmt.Fprintf(fOutput, "%s\r\n", *userPtr)
			fmt.Fprintf(fOutput, ":%s>%s\r\n", strings.ToUpper(hex.EncodeToString(data)), checksum(data))
		}
	}
}
