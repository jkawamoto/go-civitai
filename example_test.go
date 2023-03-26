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

func ExampleGetCreators() {
	ctx := context.Background()

	res, err := client.Default.Operations.GetCreators(
		operations.NewGetCreatorsParamsWithContext(ctx).WithLimit(swag.Int64(3)))
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Payload.Items {
		log.Println(v.Username, v.ModelCount)
	}
}

func ExampleGetModels() {
	ctx := context.Background()

	res, err := client.Default.Operations.GetModels(
		operations.NewGetModelsParamsWithContext(ctx).WithLimit(swag.Int64(3)))
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Payload.Items {
		log.Println(v.Name, v.ID, v.Tags)
	}
}

func ExampleGetModel() {
	ctx := context.Background()

	res, err := client.Default.Operations.GetModel(
		operations.NewGetModelParamsWithContext(ctx).WithModelID(6424))
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Payload.ModelVersions {
		log.Println(v.Name, v.DownloadURL, v.ID, v.Files[0].ID)
	}
}

func ExampleGetModelVersion() {
	ctx := context.Background()

	res, err := client.Default.Operations.GetModelVersion(
		operations.NewGetModelVersionParamsWithContext(ctx).WithModelVersionID(8958))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Payload.Name, res.Payload.CreatedAt)
}

func ExampleGetModelVersionByHash() {
	ctx := context.Background()
	hash := "64018b0e58e2495dbdc6b5ddfd97b39528af531c97ab4073ff13b45858a200a2"

	res, err := client.Default.Operations.GetModelVersionByHash(
		operations.NewGetModelVersionByHashParamsWithContext(ctx).WithHash(hash))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Payload.Name, res.Payload.UpdatedAt)
}

func ExampleGetTags() {
	ctx := context.Background()

	res, err := client.Default.Operations.GetTags(
		operations.NewGetTagsParamsWithContext(ctx).WithLimit(swag.Int64(3)))
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Payload.Items {
		log.Println(v.Name, v.ModelCount)
	}
}
