// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
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

package types

import "github.com/glant509/go-substrate-rpc-client/v4/scale"

// Modelled after packages/types/src/Metadata/v10/toV11.ts
type MetadataV11 struct {
	MetadataV10
	Extrinsic ExtrinsicV11
}

// Modelled after packages/types/src/Metadata/v10/toV11.ts
type ExtrinsicV11 struct {
	Version          uint8
	SignedExtensions []string
}

func (m *MetadataV11) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.MetadataV10)
	if err != nil {
		return err
	}
	return decoder.Decode(&m.Extrinsic)
}

func (m MetadataV11) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.MetadataV10)
	if err != nil {
		return err
	}
	return encoder.Encode(m.Extrinsic)
}

func (e *ExtrinsicV11) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&e.Version)
	if err != nil {
		return err
	}

	return decoder.Decode(&e.SignedExtensions)
}

func (e ExtrinsicV11) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(e.Version)
	if err != nil {
		return err
	}

	return encoder.Encode(e.SignedExtensions)
}
