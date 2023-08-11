package main

import (
	"context"

	"github.com/lukehagar/plexgo"
)

func main() {

	ctx := context.TODO()

	configuration := plexgo.NewDefaultConfiguration()

	apiClient := plexgo.NewAPIClient(configuration)
	configuration.HTTPClient.RetryMax = 10

	getResults(ctx, apiClient)

	//getAllPaginatedResults(ctx, apiClient)

}

func getResults(ctx context.Context, apiClient *plexgo.APIClient) {

	apiClient.GetLibrarySections(ctx)

}
