// Copyright © 2018 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"github.com/banzaicloud/banzai-cli/.gen/pipeline"
	log "github.com/sirupsen/logrus"
)

// LogAPIError logs API request errors.
func LogAPIError(action string, err error, request interface{}) {
	if err, ok := err.(pipeline.GenericOpenAPIError); ok {
		log.Errorf("failed to %s: %v (err %[2]T, request=%+v, response=%s)", action, err, request, err.Body())
	} else {
		log.Errorf("failed to %s: %v", action, err)
	}
}
