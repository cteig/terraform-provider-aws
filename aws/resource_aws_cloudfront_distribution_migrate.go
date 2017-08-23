package aws

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func resourceAwsCloudFrontDistributionMigrateState(
	v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		log.Println("[INFO] Found AWS Cloudfront Distribution State v0; migrating to v1")
		return migrateCloudFrontDistributionStateV0toV1(is)
	default:
		return is, fmt.Errorf("Unexpected schema version: %d", v)
	}
}

func migrateCloudFrontDistributionStateV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Println("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}

	log.Printf("[DEBUG] Attributes before migration: %#v", is.Attributes)

	prefix := "cache_behavior"
	entity := resourceAwsCloudFrontDistribution()

	// Old schema for reading
	prior_schema := map[string]*schema.Schema{
		"cache_behavior": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"allowed_methods": {
						Type:     schema.TypeSet,
						Required: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"cached_methods": {
						Type:     schema.TypeSet,
						Required: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"compress": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"default_ttl": {
						Type:     schema.TypeInt,
						Required: true,
					},
					"forwarded_values": {
						Type:     schema.TypeSet,
						Required: true,
						Set:      forwardedValuesHash,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"cookies": {
									Type:     schema.TypeSet,
									Required: true,
									Set:      cookiePreferenceHash,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"forward": {
												Type:     schema.TypeString,
												Required: true,
											},
											"whitelisted_names": {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
										},
									},
								},
								"headers": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"query_string": {
									Type:     schema.TypeBool,
									Required: true,
								},
								"query_string_cache_keys": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
					"lambda_function_association": {
						Type:     schema.TypeSet,
						Optional: true,
						MaxItems: 4,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"event_type": {
									Type:     schema.TypeString,
									Required: true,
								},
								"lambda_arn": {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
						Set: lambdaFunctionAssociationHash,
					},
					"max_ttl": {
						Type:     schema.TypeInt,
						Required: true,
					},
					"min_ttl": {
						Type:     schema.TypeInt,
						Required: true,
					},
					"path_pattern": {
						Type:     schema.TypeString,
						Required: true,
					},
					"smooth_streaming": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"target_origin_id": {
						Type:     schema.TypeString,
						Required: true,
					},
					"trusted_signers": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"viewer_protocol_policy": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}

	// Read old set
	reader := &schema.MapFieldReader{
		Schema: prior_schema,
		Map:    schema.BasicMapReader(is.Attributes),
	}
	result, err := reader.ReadField([]string{prefix})
	if err != nil {
		return nil, err
	}

	oldSet, ok := result.Value.(*schema.Set)
	if !ok {
		return nil, fmt.Errorf("Got unexpected value from state: %#v", result.Value)
	}

	// Convert to list
	newList := oldSet.List()

	// Delete old set
	for k := range is.Attributes {
		if strings.HasPrefix(k, fmt.Sprintf("%s.", prefix)) {
			delete(is.Attributes, k)
		}
	}

	// Write new list
	writer := schema.MapFieldWriter{
		Schema: entity.Schema,
	}
	if err := writer.WriteField([]string{prefix}, newList); err != nil {
		return is, err
	}
	for k, v := range writer.Map() {
		is.Attributes[k] = v
	}

	log.Printf("[DEBUG] Attributes after migration: %#v", is.Attributes)
	return is, nil
}
