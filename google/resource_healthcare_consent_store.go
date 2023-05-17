// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceHealthcareConsentStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcareConsentStoreCreate,
		Read:   resourceHealthcareConsentStoreRead,
		Update: resourceHealthcareConsentStoreUpdate,
		Delete: resourceHealthcareConsentStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceHealthcareConsentStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dataset": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `Identifies the dataset addressed by this request. Must be in the format
'projects/{project}/locations/{location}/datasets/{dataset}'`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name of this ConsentStore, for example:
"consent1"`,
			},
			"default_consent_ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.

A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
			},
			"enable_consent_create_on_update": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-supplied key-value pairs used to organize Consent stores.

Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
conform to the following PCRE regular expression: '[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}'

Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
bytes, and must conform to the following PCRE regular expression: '[\p{Ll}\p{Lo}\p{N}_-]{0,63}'

No more than 64 labels can be associated with a given store.

An object containing a list of "key": value pairs.
Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceHealthcareConsentStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	defaultConsentTtlProp, err := expandHealthcareConsentStoreDefaultConsentTtl(d.Get("default_consent_ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_consent_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultConsentTtlProp)) && (ok || !reflect.DeepEqual(v, defaultConsentTtlProp)) {
		obj["defaultConsentTtl"] = defaultConsentTtlProp
	}
	enableConsentCreateOnUpdateProp, err := expandHealthcareConsentStoreEnableConsentCreateOnUpdate(d.Get("enable_consent_create_on_update"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_consent_create_on_update"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableConsentCreateOnUpdateProp)) && (ok || !reflect.DeepEqual(v, enableConsentCreateOnUpdateProp)) {
		obj["enableConsentCreateOnUpdate"] = enableConsentCreateOnUpdateProp
	}
	labelsProp, err := expandHealthcareConsentStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/consentStores?consentStoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ConsentStore: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ConsentStore: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{dataset}}/consentStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ConsentStore %q: %#v", d.Id(), res)

	return resourceHealthcareConsentStoreRead(d, meta)
}

func resourceHealthcareConsentStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/consentStores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("HealthcareConsentStore %q", d.Id()))
	}

	if err := d.Set("default_consent_ttl", flattenHealthcareConsentStoreDefaultConsentTtl(res["defaultConsentTtl"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConsentStore: %s", err)
	}
	if err := d.Set("enable_consent_create_on_update", flattenHealthcareConsentStoreEnableConsentCreateOnUpdate(res["enableConsentCreateOnUpdate"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConsentStore: %s", err)
	}
	if err := d.Set("labels", flattenHealthcareConsentStoreLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConsentStore: %s", err)
	}

	return nil
}

func resourceHealthcareConsentStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	defaultConsentTtlProp, err := expandHealthcareConsentStoreDefaultConsentTtl(d.Get("default_consent_ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_consent_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultConsentTtlProp)) {
		obj["defaultConsentTtl"] = defaultConsentTtlProp
	}
	enableConsentCreateOnUpdateProp, err := expandHealthcareConsentStoreEnableConsentCreateOnUpdate(d.Get("enable_consent_create_on_update"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_consent_create_on_update"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableConsentCreateOnUpdateProp)) {
		obj["enableConsentCreateOnUpdate"] = enableConsentCreateOnUpdateProp
	}
	labelsProp, err := expandHealthcareConsentStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/consentStores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ConsentStore %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("default_consent_ttl") {
		updateMask = append(updateMask, "defaultConsentTtl")
	}

	if d.HasChange("enable_consent_create_on_update") {
		updateMask = append(updateMask, "enableConsentCreateOnUpdate")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating ConsentStore %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ConsentStore %q: %#v", d.Id(), res)
	}

	return resourceHealthcareConsentStoreRead(d, meta)
}

func resourceHealthcareConsentStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/consentStores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ConsentStore %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ConsentStore")
	}

	log.Printf("[DEBUG] Finished deleting ConsentStore %q: %#v", d.Id(), res)
	return nil
}

func resourceHealthcareConsentStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"(?P<dataset>.+)/consentStores/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{dataset}}/consentStores/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenHealthcareConsentStoreDefaultConsentTtl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareConsentStoreEnableConsentCreateOnUpdate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareConsentStoreLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandHealthcareConsentStoreDefaultConsentTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareConsentStoreEnableConsentCreateOnUpdate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareConsentStoreLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
