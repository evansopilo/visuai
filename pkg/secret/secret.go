package secret

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type Secret struct {
	client *azsecrets.Client
}

func New(vaultURI string) (*Secret, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}
	client := azsecrets.NewClient(vaultURI, cred, nil)
	return &Secret{client: client}, nil
}

func (s Secret) GetSecret(ctx context.Context, secretName, version string) (*string, error) {
	resp, err := s.client.GetSecret(ctx, secretName, version, nil)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}
