// package main

// import (
// 	"Template_Echo/pkg/configs"
// 	"Template_Echo/pkg/routes"
// 	"Template_Echo/pkg/utils"
// 	"fmt"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func main() {
// 	app := echo.New()
// 	routes.Routes(app)

// 	configs.Configs()
// 	appPort := configs.AppPort()

// 	logConfig := utils.DefaultLoggerConfig
// 	logConfig.Output = utils.Log()
// 	logConfig.Format = `${remote_ip} ${data_in_out} | ${method}:${uri} | ${status} | ${latency_human} | ${error}`

// 	app.Use(utils.LoggerWithConfig(logConfig))
// 	app.Use(middleware.Recover())

// 	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", appPort)))
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}