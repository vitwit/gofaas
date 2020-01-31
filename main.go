package go_faas

import (
	"fmt"
	"github.com/vitwit/go-faas/sdk"
)

func NewClient(creds *sdk.FaasGatewayCredentials) (*sdk.OpenFaasClient, error) {
	if creds.ClusterType != "swarm" && creds.ClusterType != "kubernetes" {
		return &sdk.OpenFaasClient{}, fmt.Errorf("invalid cluster type %v", creds.ClusterType)
	}

	return sdk.InitOpenFaasClient(creds), nil
}
