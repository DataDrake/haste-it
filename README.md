# haste-it

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/haste-it)](https://goreportcard.com/report/github.com/DataDrake/haste-it) [![license](https://img.shields.io/github/license/DataDrake/haste-it.svg)]() 

Simple little hastebin CLI tool written in Go

## Usage

Haste a file:

```
haste-it /path/to/file
```


Haste from a pipe:

```
echo "Hello World" | haste-it
```

## Build & Install

```
make
```

```
sudo make install [DESTDIR=/path/to/staging]
```

## License
 
Copyright 2017 Bryan T. Meyers <bmeyers@datadrake.com>
 
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 
http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.