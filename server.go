package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	driver "github.com/arangodb/go-driver"
	db "github.com/arangodb/go-driver/http"

	"github.com/flexzuu/gas-price-server/gasPrice"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

var schema *graphql.Schema

func init() {
	conn, err := db.NewConnection(db.ConnectionConfig{
		Endpoints: []string{os.Getenv("DB_CONNECTION")},
	})
	if err != nil {
		// Handle error
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", os.Getenv("DB_PASSWORD")),
	})
	if err != nil {
		// Handle error
	}
	schema = graphql.MustParseSchema(gasPrice.Schema, &gasPrice.Resolver{&client})
}

func main() {
	fmt.Print("Server started - localhost:8080")
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/graphql", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
