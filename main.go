package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var objectType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Object",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func main() {
	type Object struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	objects := map[string]Object{
		"42": {
			ID:   "42",
			Name: "John Doe",
		},
	}

	fields := graphql.Fields{
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Retrieve the name of a random fictional animal",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Marsupilami 🐒", nil
			},
		},
		"objectByID": &graphql.Field{
			Type: objectType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, isOK := p.Args["id"].(string)
				if isOK {
					return objects[idQuery], nil
				}
				return nil, nil
			},
			Description: "Get an object by its ID",
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", graphqlHandler)
	err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("failed to listen and serve, error: %v", err)
	}
}
