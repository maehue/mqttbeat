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

package elasticsearch

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
