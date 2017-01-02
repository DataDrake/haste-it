//
// Copyright © 2017 Bryan T. Meyers <bmeyers@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implie$
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"encoding/json"
    "flag"
	"net/http"
	"os"
)

func haste(f *os.File) {
	c := &http.Client{}
	r, err := c.Post("https://hastebin.com/documents", "application/json", f)
	if err != nil {
		println("Failed to send, reason: " + err.Error())
		os.Exit(1)
	}
	d := json.NewDecoder(r.Body)
	var v struct {
		Key string
        Message string
	}
	err = d.Decode(&v)
	if err != nil {
		println("Failed to parse response")
		os.Exit(1)
	}
    if len(v.Key) == 0 {
        if len(v.Message) > 0 {
            println("Haste failed, reason: " + v.Message)
        } else {
            println("Haste failed for unknown reason")
        }
        os.Exit(1)
    }
	println("https://hastebin.com/" + v.Key)
}

func hasteFile(fp string) {
	f, err := os.Open(fp)
	if err != nil {
		println("Failed to open file: '" + fp + "', reason: " + err.Error())
		os.Exit(1)
	}
	defer f.Close()
	haste(f)
}

func usage() {
	print("haste-it --- Simple little client for hastebin\n\n")
	print("Usage: haste [file]\n\n")
	print("    Read from specified [file] or stdin\n\n")
    flag.PrintDefaults()
}

func main() {
    flag.Usage = usage
    var h1 = flag.Bool("h", false, "Same as --help")
    var h2 = flag.Bool("-help", false, "Print usage")
    flag.Parse()
    if *h1 || *h2 {
        usage()
        os.Exit(0)
    }
	switch len(flag.Args()) {
	case 0:
		haste(os.Stdin)
	case 1:
		hasteFile(flag.Args()[0])
	default:
		usage()
		os.Exit(1)
	}
}
