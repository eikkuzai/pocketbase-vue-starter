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
			"id": "4t0byk7a8t8ecow",
			"created": "2023-04-02 13:03:03.440Z",
			"updated": "2023-04-02 13:03:03.440Z",
			"name": "bookings",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "pmr3gi3n",
					"name": "startTime",
					"type": "date",
					"required": true,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "g45rrfez",
					"name": "endTime",
					"type": "date",
					"required": true,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "vqbhfjyt",
					"name": "customerGuid",
					"type": "relation",
					"required": true,
					"unique": true,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "5tsarfzy",
					"name": "serviceGuid",
					"type": "relation",
					"required": true,
					"unique": true,
					"options": {
						"collectionId": "h4fmcqhn3a6msl7",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
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

		collection, err := dao.FindCollectionByNameOrId("4t0byk7a8t8ecow")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
