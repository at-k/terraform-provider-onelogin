package usercustomattributesschema

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/users"
)

// Schema returns a key/value map of the various fields that make up a OneLogin User.
func Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"custom_attributes": &schema.Schema{
			Type:     schema.TypeMap,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

// Inflate takes a map representation of a User and returns a User object
func Inflate(s map[string]interface{}) (users.User, error) {
	out := users.User{}

	if user_id, notNil := s["user_id"].(string); notNil {
		id_i, _ := strconv.Atoi(user_id)
		out.ID = oltypes.Int32(int32(id_i))
	}

	if custom_attributes, notNil := s["custom_attributes"].(map[string]interface{}); notNil {
		out.CustomAttributes = custom_attributes
	}

	return out, nil
}
