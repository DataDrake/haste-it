//
// Copyright 2017-2020 Bryan T. Meyers <bmeyers@datadrake.com>
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
	"errors"
	"flag"
	"net/http"
	"os"
)

func haste(f *os.File) error {
	r, err := http.Post("https://hastebin.com/documents", "application/json", f)
	if err != nil {
		return errors.New("failed to send, reason: " + err.Error())
	}
	defer r.Body.Close()
	var v struct {
		Key     string
		Message string
	}
	if err = json.NewDecoder(r.Body).Decode(&v); err != nil {
		return errors.New("failed to parse response, code: " + r.Status)
	}
	if len(v.Key) == 0 {
		if len(v.Message) > 0 {
			return errors.New("haste failed, reason: " + v.Message)
		}
		return errors.New("Haste failed for unknown reason")
	}
	println("https://hastebin.com/" + v.Key)
	return nil
}

func hasteFile(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return errors.New("Failed to open file: '" + fp + "', reason: " + err.Error())
	}
	defer f.Close()
	return haste(f)
}

func usage() {
	print("haste-it --- Simple little client for hastebin\n\n")
	print("Usage: haste [file]\n\n")
	print("    Read from specified [file] or stdin\n\n")
	flag.PrintDefaults()
}

func main() {
	var h1 = flag.Bool("h", false, "Same as --help")
	var h2 = flag.Bool("-help", false, "Print usage")
	flag.Parse()
	if *h1 || *h2 {
		usage()
		os.Exit(0)
	}
	var err error
	switch len(flag.Args()) {
	case 0:
		err = haste(os.Stdin)
	case 1:
		err = hasteFile(flag.Args()[0])
	default:
		usage()
		os.Exit(1)
	}
	if err != nil {
		println(err)
		os.Exit(1)
	}
}
