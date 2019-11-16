package main

import (
	"net/http"

	"github.com/friendsofgo/graphiql"
)

func main() {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	//http.Handle("/graphql", gqlHandler())
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)

}
