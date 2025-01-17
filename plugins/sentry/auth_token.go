package sentry

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func AuthToken() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.AuthToken,
		DocsURL:       sdk.URL("https://docs.sentry.io/api/auth/"),
		ManagementURL: sdk.URL("https://sentry.io/settings/account/api/auth-tokens/"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Token,
				MarkdownDescription: "Token used to authenticate to Sentry.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 64,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                fieldname.Organization,
				MarkdownDescription: "The default organization used for this auth token.",
				Optional:            true,
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
			TrySentryclircFile(),
		),
	}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"SENTRY_AUTH_TOKEN": fieldname.Token,
	"SENTRY_ORG":        fieldname.Organization,
}

func TrySentryclircFile() sdk.Importer {
	// TODO: Try importing from ~/.sentryclirc file
	return importer.NoOp()
}
