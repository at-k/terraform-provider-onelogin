package onelogin

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/onelogin/onelogin-go-sdk/pkg/client"
	"github.com/onelogin/terraform-provider-onelogin/ol_schema/user_custom_attributes"
)

// Users returns a resource with the CRUD methods and Terraform Schema defined
func UserCustomAttributes() *schema.Resource {
	return &schema.Resource{
		Create:   userCustomAttributesCreate,
		Read:     userCustomAttributesRead,
		Update:   userCustomAttributesUpdate,
		Delete:   userCustomAttributesDelete,
		Importer: &schema.ResourceImporter{},
		Schema:   usercustomattributesschema.Schema(),
	}
}

func userCustomAttributesCreate(d *schema.ResourceData, m interface{}) error {
	user, _ := usercustomattributesschema.Inflate(map[string]interface{}{
		"user_id":           d.Get("user_id"),
		"custom_attributes": d.Get("custom_attributes"),
	})
	client := m.(*client.APIClient)
	err := client.Services.UsersV2.Update(&user)
	if err != nil {
		log.Println("[ERROR] There was a problem creating the user!", err)
		return err
	}
	log.Printf("[CREATED] Created user with %d", *(user.ID))

	d.SetId(fmt.Sprintf("%d", *(user.ID)))
	return userCustomAttributesRead(d, m)
}

func userCustomAttributesUpdate(d *schema.ResourceData, m interface{}) error {
	user, _ := usercustomattributesschema.Inflate(map[string]interface{}{
		"user_id":           d.Get("user_id"),
		"custom_attributes": d.Get("custom_attributes"),
	})
	client := m.(*client.APIClient)
	err := client.Services.UsersV2.Update(&user)
	if err != nil {
		log.Println("[ERROR] There was a problem updating the user!", err)
		return err
	}
	log.Printf("[UPDATED] Updated user with %d", *(user.ID))

	d.SetId(fmt.Sprintf("%d", *(user.ID)))
	return userCustomAttributesRead(d, m)
}

func userCustomAttributesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.APIClient)
	uid, _ := strconv.Atoi(d.Id())
	user, err := client.Services.UsersV2.GetOne(int32(uid))
	if err != nil {
		log.Printf("[ERROR] There was a problem reading the user!")
		log.Println(err)
		return err
	}
	if user == nil {
		d.SetId("")
		return nil
	}
	log.Printf("[READ] Reading user with %d", *(user.ID))

	d.Set("custom_attributes", user.CustomAttributes)

	return nil
}

func userCustomAttributesDelete(d *schema.ResourceData, m interface{}) error {
	// user, _ := usercustomattributesschema.Inflate(map[string]interface{}{
	// 	"user_id":           d.Get("user_id"),
	// 	"custom_attributes": nil,
	// })
	uid, _ := strconv.Atoi(d.Id())
	// client := m.(*client.APIClient)
	// err := client.Services.UsersV2.Update(&user)
	//if err != nil {
	//	log.Printf("[ERROR] There was a problem deleting the user!")
	//	log.Println(err)
	//} else {
	log.Printf("[DELETED] Deleted user with %d", uid)
	d.SetId("")
	// }

	return nil
}
