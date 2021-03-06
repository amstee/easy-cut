package vars

import "github.com/amstee/easy-cut/src/es"

const RatingMapping = `
	{
		"settings": {
			"number_of_shards": 2,
			"number_of_replicas": 1
		},
		"mappings": {
			"rating": {
				"properties": {
					"stars": {
						"type": "integer"
					},
					"user_id": {
						"type": "keyword"
					},
					"target_id": {
						"type": "keyword"
					},
					"target_type": {
						"type": "keyword"
					},
					"comment": {
						"type": "keyword"
					},
					"created": {
						"type": "date"
					},
					"updated": {
						"type": "date"
					}
				}
			}
		}
	}
`


func Register() error {
	return es.RegisterIndex("rating", RatingMapping)
}