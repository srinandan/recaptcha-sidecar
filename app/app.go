// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/srinandan/recaptcha-sidecar/recaptcha"
)

//log levels, default is error
var (
	//Info is used for debug logs
	Info *log.Logger
	//Error is used to log errors
	Error *log.Logger
)

//Ctx for client connection
var Ctx context.Context

//Address to start server
const Address = "0.0.0.0:8080"

//InitLog function initializes the logger objects
func initLog() {
	var infoHandle = ioutil.Discard

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	if debug {
		infoHandle = os.Stdout
	}

	errorHandle := os.Stdout

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//Initialize logging, context, sec mgr and kms
func Initialize() {
	//init logging
	initLog()

	if err := recaptcha.Init(); err != nil {
		Error.Fatalln("error starting app: ", err)
	}
}
