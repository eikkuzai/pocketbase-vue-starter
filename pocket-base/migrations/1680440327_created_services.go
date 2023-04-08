package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "h4fmcqhn3a6msl7",
			"created": "2023-04-02 12:58:47.932Z",
			"updated": "2023-04-02 12:58:47.932Z",
			"name": "services",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "yicabzui",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": true,
					"options": {
						"min": 1,
						"max": 50,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "rp3nf5vk",
					"name": "description",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 250,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "bauzym3a",
					"name": "serviceImage",
					"type": "file",
					"required": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"maxSize": 2500000,
						"mimeTypes": [
							"image/png",
							"image/jpeg",
							"image/webp"
						],
						"thumbs": []
					}
				},
				{
					"system": false,
					"id": "ajgntap2",
					"name": "price",
					"type": "number",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null
					}
				},
				{
					"system": false,
					"id": "0xjvlyn6",
					"name": "durationMin",
					"type": "number",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null
					}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("h4fmcqhn3a6msl7")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
