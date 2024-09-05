package minio

import (
	"context"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMinioLdapIntergration() *schema.Resource {
	return &schema.Resource{
		Description:   "Intergration minio server with ldap service for access management.",
		CreateContext: minioCreateMinioLdapIntergration,
		ReadContext:   minioReadMinioLdapIntergration,
		UpdateContext: minioPutMinioLdapIntergration,
		DeleteContext: minioDeleteMinioLdapIntergration,
		Importer:      &schema.ResourceImporter{
			// StateContext: minioImportLDAPGroupPolicyAttachment,
		},
		Schema: map[string]*schema.Schema{
			"server_addr": {
				Type:        schema.TypeString,
				Description: "Ldap server address",
				Required:    true,
				// ValidateFunc: validateIAMNamePolicy,
			},
			"lookup_bind_dn": {
				Type:        schema.TypeString,
				Description: "DN (Distinguished Name) for LDAP read-only service account used to perform DN and group lookups.",
				Required:    true,
				// ValidateFunc: validateMinioIamGroupName,
			},
			"lookup_bind_password": {
				Type:        schema.TypeString,
				Description: "Password for LDAP read-only service account used to perform DN and group lookups.",
				Required:    true,
				// ValidateFunc: validateMinioIamGroupName,
			},
			"user_dn_search_base_dn": {
				Type:        schema.TypeString,
				Description: "Base DN under which to perform user search.",
				Required:    true,
				// ValidateFunc: validateMinioIamGroupName,
			},
			"user_dn_search_filter": {
				Type:        schema.TypeString,
				Description: "LDAP user search filter.",
				Required:    true,
				// ValidateFunc: validateMinioIamGroupName,
			},
			"group_search_base_dn": {
				Type:        schema.TypeString,
				Description: "Base DN under which to perform group search.",
				Optional:    true,
				// ValidateFunc: validateMinioIamGroupName,
			},
			"group_search_filter": {
				Type:        schema.TypeString,
				Description: "LDAP group search filter.",
				Optional:    true,
				// ValidateFunc: validateMinioIamGroupName,
			},
			"server_insecure": {
				Type:        schema.TypeString,
				Description: "Ldap server address",
				Optional:    true,
				Default:     "on",
				// ValidateFunc: validateIAMNamePolicy,
			},
		},
	}
}

func minioCreateMinioLdapIntergration(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := uuid.New().String()
	d.SetId("minio_ldap_config" + id)

	return minioReadMinioLdapIntergration(ctx, d, meta)
}

func minioPutMinioLdapIntergration(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("meta value: %+v", meta)

	// ldapConfig := LdapCheckConfig(d, meta)

	// log.Printf("LDAP Configuration: %+v", ldapConfig)

	return nil
}

func minioReadMinioLdapIntergration(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	ldapConfig := LdapCheckConfig(d, meta)
	log.Printf("LDAP Configuration: %+v", ldapConfig)

	typ := reflect.TypeOf(ldapConfig)
	val := reflect.ValueOf(ldapConfig)

	for i := 0; i < val.NumField(); i++ {

		if err := d.Set(typ.Field(i).Name, val.Field(i).Interface()); err != nil {
			return NewResourceError("error reading Ldap config %s: %s", d.Id(), err)
		}
	}

	return nil
}

func minioDeleteMinioLdapIntergration(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

// {key: "server_insecure", value: "on"}
// {key: "server_addr", value: "10.10.3.64:30089"}
// {key: "lookup_bind_dn", value: "cn=admin,dc=example,dc=org"}
// {key: "lookup_bind_password", value: "Not@SecurePassw0rd"}
// {key: "user_dn_search_base_dn", value: "ou=users,dc=example,dc=org"}
// {key: "user_dn_search_filter", value: "(&(objectClass=inetOrgPerson)(uid=%s))"}
// {key: "group_search_base_dn", value: "ou=test-group,dc=example,dc=org"}
// {key: "group_search_filter", value: "(&(objectClass=groupOfNames)(member=%d))"}
// http://10.10.2.54:9001/api/v1/configs/identity_ldap
