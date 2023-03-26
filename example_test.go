// example_test.go
//
// Copyright (c) 2023 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php

package civitai

import (
	"context"
	"log"

	"github.com/go-openapi/swag"
	"github.com/jkawamoto/go-civitai/client"
	"github.com/jkawamoto/go-civitai/client/operations"
)

func Example() {
	ctx := context.Background()

	// List creators.
	creators, err := client.Default.Operations.GetCreators(
		operations.NewGetCreatorsParamsWithContext(ctx).WithLimit(swag.Int64(3)))
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range creators.Payload.Items {
		log.Println(v.Username, v.ModelCount)
	}

	// List models.
	models, err := client.Default.Operations.GetModels(
		operations.NewGetModelsParamsWithContext(ctx).WithLimit(swag.Int64(3)))
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range models.Payload.Items {
		log.Println(v.Name, v.ID, v.Tags)
	}

	// Retrieve a model.
	m, err := client.Default.Operations.GetModel(
		operations.NewGetModelParamsWithContext(ctx).WithModelID(6424))
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range m.Payload.ModelVersions {
		log.Println(v.Name, v.DownloadURL, v.ID, v.Files[0].ID)
	}

	// Retrieve a model version.
	ver, err := client.Default.Operations.GetModelVersion(
		operations.NewGetModelVersionParamsWithContext(ctx).WithModelVersionID(8958))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ver.Payload.Name, ver.Payload.CreatedAt)

	// Retrieve a model version by a hash.
	hash := "64018b0e58e2495dbdc6b5ddfd97b39528af531c97ab4073ff13b45858a200a2"
	ver2, err := client.Default.Operations.GetModelVersionByHash(
		operations.NewGetModelVersionByHashParamsWithContext(ctx).WithHash(hash))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ver2.Payload.Name, ver2.Payload.UpdatedAt)

	// List tags.
	tags, err := client.Default.Operations.GetTags(
		operations.NewGetTagsParamsWithContext(ctx).WithLimit(swag.Int64(3)))
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range tags.Payload.Items {
		log.Println(v.Name, v.ModelCount)
	}
}
