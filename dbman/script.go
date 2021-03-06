/*
   Onix Config Manager - Dbman- Onix Database Manager
   Copyright (c) 2018-2020 by www.gatblau.org

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software distributed under
   the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
   either express or implied.
   See the License for the specific language governing permissions and limitations under the License.

   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/
package main

import (
	"errors"
	"fmt"
)

type Script struct {
	client    *Client
	index     Index
	manifests []Manifest
	cfg       *Config
}

func NewScript(cfg *Config) (*Script, error) {
	script := new(Script)
	script.cfg = cfg
	script.client = new(Client)
	return script, nil
}

// fetches the release index
func (s *Script) fetchIndex() (*Index, error) {
	if s.cfg == nil {
		return nil, errors.New("configuration object not set")
	}
	response, err := s.client.get(fmt.Sprintf("%s/index.json", s.cfg.SchemaURI))
	if err != nil {
		return nil, err
	}
	i := &Index{}
	i, err = i.decode(response)

	defer func() {
		if ferr := response.Body.Close(); ferr != nil {
			err = ferr
		}
	}()
	return i, err
}
