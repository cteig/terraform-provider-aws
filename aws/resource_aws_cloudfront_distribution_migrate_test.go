package aws

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/terraform"
)

func TestAwsCloudFrontDistributionMigrateState(t *testing.T) {
	cases := map[string]struct {
		StateVersion int
		ID           string
		Attributes   map[string]string
		Expected     map[string]string
		Meta         interface{}
	}{
		"v0_1": {
			StateVersion: 0,
			ID:           "some_id",
			Attributes: map[string]string{
				"cache_behavior.#":                                                                             "2",
				"cache_behavior.1328956990.allowed_methods.#":                                                  "2",
				"cache_behavior.1328956990.allowed_methods.0":                                                  "HEAD",
				"cache_behavior.1328956990.allowed_methods.1":                                                  "GET",
				"cache_behavior.1328956990.cached_methods.#":                                                   "2",
				"cache_behavior.1328956990.cached_methods.0":                                                   "HEAD",
				"cache_behavior.1328956990.cached_methods.1":                                                   "GET",
				"cache_behavior.1328956990.compress":                                                           "false",
				"cache_behavior.1328956990.default_ttl":                                                        "86400",
				"cache_behavior.1328956990.forwarded_values.#":                                                 "1",
				"cache_behavior.1328956990.forwarded_values.2759845635.cookies.#":                              "1",
				"cache_behavior.1328956990.forwarded_values.2759845635.cookies.2625240281.forward":             "none",
				"cache_behavior.1328956990.forwarded_values.2759845635.cookies.2625240281.whitelisted_names.#": "0",
				"cache_behavior.1328956990.forwarded_values.2759845635.headers.#":                              "0",
				"cache_behavior.1328956990.forwarded_values.2759845635.query_string":                           "false",
				"cache_behavior.1328956990.forwarded_values.2759845635.query_string_cache_keys.#":              "0",
				"cache_behavior.1328956990.lambda_function_association.#":                                      "0",
				"cache_behavior.1328956990.max_ttl":                                                            "31536000",
				"cache_behavior.1328956990.min_ttl":                                                            "0",
				"cache_behavior.1328956990.path_pattern":                                                       "/robots.txt",
				"cache_behavior.1328956990.smooth_streaming":                                                   "false",
				"cache_behavior.1328956990.target_origin_id":                                                   "foo",
				"cache_behavior.1328956990.trusted_signers.#":                                                  "0",
				"cache_behavior.1328956990.viewer_protocol_policy":                                             "allow-all",
				"cache_behavior.3468461710.allowed_methods.#":                                                  "2",
				"cache_behavior.3468461710.allowed_methods.0":                                                  "HEAD",
				"cache_behavior.3468461710.allowed_methods.1":                                                  "GET",
				"cache_behavior.3468461710.cached_methods.#":                                                   "2",
				"cache_behavior.3468461710.cached_methods.0":                                                   "HEAD",
				"cache_behavior.3468461710.cached_methods.1":                                                   "GET",
				"cache_behavior.3468461710.compress":                                                           "false",
				"cache_behavior.3468461710.default_ttl":                                                        "86400",
				"cache_behavior.3468461710.forwarded_values.#":                                                 "1",
				"cache_behavior.3468461710.forwarded_values.2759845635.cookies.#":                              "1",
				"cache_behavior.3468461710.forwarded_values.2759845635.cookies.2625240281.forward":             "none",
				"cache_behavior.3468461710.forwarded_values.2759845635.cookies.2625240281.whitelisted_names.#": "0",
				"cache_behavior.3468461710.forwarded_values.2759845635.headers.#":                              "0",
				"cache_behavior.3468461710.forwarded_values.2759845635.query_string":                           "false",
				"cache_behavior.3468461710.forwarded_values.2759845635.query_string_cache_keys.#":              "0",
				"cache_behavior.3468461710.lambda_function_association.#":                                      "0",
				"cache_behavior.3468461710.max_ttl":                                                            "31536000",
				"cache_behavior.3468461710.min_ttl":                                                            "0",
				"cache_behavior.3468461710.path_pattern":                                                       "/favicon.ico",
				"cache_behavior.3468461710.smooth_streaming":                                                   "false",
				"cache_behavior.3468461710.target_origin_id":                                                   "foo",
				"cache_behavior.3468461710.trusted_signers.#":                                                  "0",
				"cache_behavior.3468461710.viewer_protocol_policy":                                             "allow-all",
			},
			Expected: map[string]string{
				"cache_behavior.#":                                                                    "2",
				"cache_behavior.0.allowed_methods.#":                                                  "2",
				"cache_behavior.0.allowed_methods.1445840968":                                         "HEAD",
				"cache_behavior.0.allowed_methods.1040875975":                                         "GET",
				"cache_behavior.0.cached_methods.#":                                                   "2",
				"cache_behavior.0.cached_methods.1445840968":                                          "HEAD",
				"cache_behavior.0.cached_methods.1040875975":                                          "GET",
				"cache_behavior.0.compress":                                                           "false",
				"cache_behavior.0.default_ttl":                                                        "86400",
				"cache_behavior.0.forwarded_values.#":                                                 "1",
				"cache_behavior.0.forwarded_values.2759845635.cookies.#":                              "1",
				"cache_behavior.0.forwarded_values.2759845635.cookies.2625240281.forward":             "none",
				"cache_behavior.0.forwarded_values.2759845635.cookies.2625240281.whitelisted_names.#": "0",
				"cache_behavior.0.forwarded_values.2759845635.headers.#":                              "0",
				"cache_behavior.0.forwarded_values.2759845635.query_string":                           "false",
				"cache_behavior.0.forwarded_values.2759845635.query_string_cache_keys.#":              "0",
				"cache_behavior.0.lambda_function_association.#":                                      "0",
				"cache_behavior.0.max_ttl":                                                            "31536000",
				"cache_behavior.0.min_ttl":                                                            "0",
				"cache_behavior.0.path_pattern":                                                       "/favicon.ico",
				"cache_behavior.0.smooth_streaming":                                                   "false",
				"cache_behavior.0.target_origin_id":                                                   "foo",
				"cache_behavior.0.trusted_signers.#":                                                  "0",
				"cache_behavior.0.viewer_protocol_policy":                                             "allow-all",
				"cache_behavior.1.allowed_methods.#":                                                  "2",
				"cache_behavior.1.allowed_methods.1445840968":                                         "HEAD",
				"cache_behavior.1.allowed_methods.1040875975":                                         "GET",
				"cache_behavior.1.cached_methods.#":                                                   "2",
				"cache_behavior.1.cached_methods.1445840968":                                          "HEAD",
				"cache_behavior.1.cached_methods.1040875975":                                          "GET",
				"cache_behavior.1.compress":                                                           "false",
				"cache_behavior.1.default_ttl":                                                        "86400",
				"cache_behavior.1.forwarded_values.#":                                                 "1",
				"cache_behavior.1.forwarded_values.2759845635.cookies.#":                              "1",
				"cache_behavior.1.forwarded_values.2759845635.cookies.2625240281.forward":             "none",
				"cache_behavior.1.forwarded_values.2759845635.cookies.2625240281.whitelisted_names.#": "0",
				"cache_behavior.1.forwarded_values.2759845635.headers.#":                              "0",
				"cache_behavior.1.forwarded_values.2759845635.query_string":                           "false",
				"cache_behavior.1.forwarded_values.2759845635.query_string_cache_keys.#":              "0",
				"cache_behavior.1.lambda_function_association.#":                                      "0",
				"cache_behavior.1.max_ttl":                                                            "31536000",
				"cache_behavior.1.min_ttl":                                                            "0",
				"cache_behavior.1.path_pattern":                                                       "/robots.txt",
				"cache_behavior.1.smooth_streaming":                                                   "false",
				"cache_behavior.1.target_origin_id":                                                   "foo",
				"cache_behavior.1.trusted_signers.#":                                                  "0",
				"cache_behavior.1.viewer_protocol_policy":                                             "allow-all",
			},
		},
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:         tc.ID,
			Attributes: tc.Attributes,
		}
		is, err := resourceAwsCloudFrontDistributionMigrateState(
			tc.StateVersion, is, tc.Meta)

		if err != nil {
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}

		if !reflect.DeepEqual(is.Attributes, tc.Expected) {
			t.Fatalf("bad Cloudfront Distribution Migrate: %#v\n\n expected: %#v", is.Attributes, tc.Expected)
		}
	}
}
