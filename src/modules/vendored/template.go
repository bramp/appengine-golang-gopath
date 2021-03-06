//Copyright 2015 Google Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package vendored

import (
	"html/template"

	uuid "github.com/nu7hatch/gouuid"
)

var display *template.Template

type data struct {
	UUID         *uuid.UUID
	Letter       string
	ReversedUUID string
}

func initTemplate() {
	t := `
		<html>
			<body>
				<strong>Vendored Module</strong><br/>
				UUID: {{.UUID}}. <br/>
				Letter: {{.Letter}}. <br/>
				Reverse UUID: {{.ReversedUUID}}. <br/>
			</body>
		</html>
	`
	display = template.Must(template.New("html").Parse(t))
}
