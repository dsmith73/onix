/*
   Onix Config Manager - Terraform Provider
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
	. "github.com/gatblau/oxc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ItemTypeAttributeResource() *schema.Resource {
	return &schema.Resource{
		Create: createItemTypeAttribute,
		Read:   readItemTypeAttr,
		Update: updateItemTypeAttribute,
		Delete: deleteItemTypeAttribute,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item_type_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"def_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func createItemTypeAttribute(data *schema.ResourceData, meta interface{}) error {
	// get the Ox client
	c := meta.(Config).Client

	// read the resource data into an Item Type Attribute
	itemTypeAttr := newItemTypeAttr(data)

	// put the Item Type Attr to the Web API
	err := err(c.PutItemTypeAttr(itemTypeAttr))
	if err != nil {
		return err
	}

	// set Item Type Attribute Id key
	data.SetId(itemTypeAttr.Key)

	return readItemTypeAttr(data, meta)
}

func updateItemTypeAttribute(data *schema.ResourceData, meta interface{}) error {
	return createItemTypeAttribute(data, meta)
}

func deleteItemTypeAttribute(data *schema.ResourceData, meta interface{}) error {
	// get the Ox client
	c := meta.(Config).Client

	// read the resource data into an Item Type Attribute
	itemTypeAttr := newItemTypeAttr(data)

	// delete the itemTypeAttr
	return err(c.DeleteItemTypeAttr(itemTypeAttr))
}

func newItemTypeAttr(data *schema.ResourceData) *ItemTypeAttribute {
	return &ItemTypeAttribute{
		Key:         data.Get("key").(string),
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		Type:        data.Get("type").(string),
		DefValue:    data.Get("def_value").(string),
		Required:    data.Get("required").(bool),
		Regex:       data.Get("regex").(string),
		ItemTypeKey: data.Get("item_type_key").(string),
		Version:     getVersion(data),
	}
}

// populate the Item Type Attribute with the data in the terraform resource
func populateItemTypeAttr(data *schema.ResourceData, typeAttr *ItemTypeAttribute) {
	data.SetId(typeAttr.Key)
	data.Set("key", typeAttr.Key)
	data.Set("description", typeAttr.Description)
	data.Set("type", typeAttr.Type)
	data.Set("def_value", typeAttr.DefValue)
	data.Set("required", typeAttr.Required)
	data.Set("regex", typeAttr.Regex)
	data.Set("item_type_key", typeAttr.ItemTypeKey)
	data.Set("created", typeAttr.Created)
	data.Set("updated", typeAttr.Updated)
	data.Set("changed_by", typeAttr.ChangedBy)
}
