// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("heartbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
