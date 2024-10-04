/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/johnmckenna/provider-dynatrace/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal dynatrace credentials as JSON"
	// provider config keys
	keyAccountID              = "account_id"
	keyAutomationClientID     = "automation_client_id"
	keyAutomationClientSecret = "automation_client_secret"
	keyAutomationEnvURL       = "automation_env_url"
	keyAutomationTokenURL     = "automation_token_url"
	keyClientID               = "client_id"
	keyClientSecret           = "client_secret"
	keyDtClusterAPIToken      = "dt_cluster_api_token"
	keyDtClusterURL           = "dt_cluster_url"
	keyDtAPIToken             = "dt_api_token"
	keyDtEnvURL               = "dt_env_url"
	keyIAMAccountID           = "iam_account_id"
	keyIAMClientID            = "iam_client_id"
	keyIAMClientSecret        = "iam_client_secret"
	keyIAMEndpointURL         = "iam_endpoint_url"
	keyIAMTokenURL            = "iam_token_url"
)

// TerraformSetupBuilder builds a Terraform setup function which returns Terraform provider setup configuration.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}

		pc, err := getProviderConfig(ctx, client, configRef.Name)
		if err != nil {
			return ps, err
		}

		if err := trackProviderConfigUsage(ctx, client, mg); err != nil {
			return ps, err
		}

		creds, err := extractCredentials(ctx, client, pc)
		if err != nil {
			return ps, err
		}

		ps.Configuration = setProviderConfiguration(creds)
		return ps, nil
	}
}

func getProviderConfig(ctx context.Context, client client.Client, name string) (*v1beta1.ProviderConfig, error) {
	pc := &v1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}
	return pc, nil
}

func trackProviderConfigUsage(ctx context.Context, client client.Client, mg resource.Managed) error {
	t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return errors.Wrap(err, errTrackUsage)
	}
	return nil
}

func extractCredentials(ctx context.Context, client client.Client, pc *v1beta1.ProviderConfig) (map[string]string, error) {
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errExtractCredentials)
	}
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}
	return creds, nil
}

func setProviderConfiguration(creds map[string]string) map[string]any {
	config := map[string]any{}
	keys := []string{
		keyAccountID, keyAutomationClientID, keyAutomationClientSecret, keyAutomationEnvURL, keyAutomationTokenURL,
		keyClientID, keyClientSecret, keyDtClusterAPIToken, keyDtClusterURL, keyDtAPIToken, keyDtEnvURL,
		keyIAMAccountID, keyIAMClientID, keyIAMClientSecret, keyIAMEndpointURL, keyIAMTokenURL,
	}
	for _, key := range keys {
		if v, ok := creds[key]; ok {
			config[key] = v
		}
	}
	return config
}
