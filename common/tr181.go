/**
* Copyright 2021 Comcast Cable Communications Management, LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
* SPDX-License-Identifier: Apache-2.0
*/
package common

const (
	TR181Str  = 0
	TR181Int  = 1
	TR181Uint = 2
	TR181Bool = 3
	TR181Blob = 12
)

type TR181Entry struct {
	Name     string `json:"name" msgpack:"name"`
	Value    string `json:"value" msgpack:"value"`
	DataType int    `json:"dataType" msgpack:"dataType,omitempty"`
}

type TR181Output struct {
	Parameters []TR181Entry `json:"parameters" msgpack:"parameters"`
}
