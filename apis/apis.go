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

package apis

import (
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/srinandan/recaptcha-sidecar/recaptcha"
	"github.com/srinandan/recaptcha-sidecar/app"

	"net/http"
)

func responseHandler(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		app.Error.Println(err)
	}
}

//HealthHandler handles kubernetes healthchecks
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(map[string]bool{"ok": true}); err != nil {
		app.Error.Println(err)
	}
}

//GetAssessmentHandler retrieves a recaptcha assessment
func GetAssessmentHandler(w http.ResponseWriter, r *http.Request) {
	//read path variables
	vars := mux.Vars(r)
	token := vars["token"]

	s, err := recaptcha.GetAssessment(token)
	if err != nil {
		app.Error.Println(err)
		responseHandler(w, map[string]string{})
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(s, &result)
	if err != nil {
		app.Error.Println(err)
		responseHandler(w, map[string]string{})
		return
	}

	responseHandler(w, result)
}
