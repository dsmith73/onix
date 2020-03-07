/*
   Onix Config Manager - Copyright (c) 2018-2019 by www.gatblau.org

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
	"github.com/hashicorp/terraform/helper/schema"
)

/*
   LINK TYPE RESOURCE
*/

func LinkTypeResource() *schema.Resource {
	return &schema.Resource{
		Create: createOrUpdateLinkType,
		Read:   readLinkType,
		Update: createOrUpdateLinkType,
		Delete: deleteLinkType,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"meta_schema": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"model_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"encrypt_meta": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"encrypt_txt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func createOrUpdateLinkType(data *schema.ResourceData, m interface{}) error {
	return put(data, m, linkTypePayload(data), "%s/linktype/%s", "key", "")
}

func deleteLinkType(data *schema.ResourceData, m interface{}) error {
	return delete(m, linkTypePayload(data), "%s/linktype/%s", "key", "")
}

func linkTypePayload(data *schema.ResourceData) Payload {
	key := data.Get("key").(string)
	name := data.Get("name").(string)
	description := data.Get("description").(string)
	modelKey := data.Get("model_key").(string)
	return &LinkType{
		Key:         key,
		Name:        name,
		Description: description,
		Model:       modelKey,
		MetaSchema:  data.Get("meta_schema").(map[string]interface{}),
		EncryptMeta: data.Get("encrypt_meta").(bool),
		EncryptTxt:  data.Get("encrypt_txt").(bool),
		Managed:     data.Get("managed").(bool),
		Tag:         data.Get("tag").([]interface{}),
	}
}