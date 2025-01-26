package graphql

import (
	"net/http"
	"order-system/initializers"
	"order-system/models"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var orderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"customerID": &graphql.Field{Type: graphql.Int},
		"driverID":   &graphql.Field{Type: graphql.Int},
		"productID":  &graphql.Field{Type: graphql.Int},
		"quantity":   &graphql.Field{Type: graphql.Int},
		"status":     &graphql.Field{Type: graphql.String},
		"location":   &graphql.Field{Type: graphql.String},
		"totalPrice": &graphql.Field{Type: graphql.Int},
		"orderDate":  &graphql.Field{Type: graphql.String},
	},
})

// Define the query type
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"order": &graphql.Field{
			Type: orderType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				var order models.Order
				if err := initializers.DB.First(&order, id).Error; err != nil {
					return nil, err
				}
				return order, nil
			},
		},
		"orders": &graphql.Field{
			Type: graphql.NewList(orderType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var orders []models.Order
				if err := initializers.DB.Find(&orders).Error; err != nil {
					return nil, err
				}
				return orders, nil
			},
		},
	},
})

// Export the schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

// GraphQL routes
func SetupGraphQLRoutes(r *gin.Engine) {
	r.GET("/graphql", func(c *gin.Context) {
		var params struct {
			Query string `json:"query"`
		}
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Execute the GraphQL query
		result := graphql.Do(graphql.Params{
			Schema:        Schema, // Use the exported Schema
			RequestString: params.Query,
		})

		// Return the result or errors
		if len(result.Errors) > 0 {
			c.JSON(http.StatusBadRequest, result)
			return
		}

		c.JSON(http.StatusOK, result)
	})
}

func SetupGraphQLRoutes2(r *gin.Engine) {
	r.POST("/graphql", func(c *gin.Context) {
		var params struct {
			Query string `json:"query"`
		}
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Execute the GraphQL query
		result := graphql.Do(graphql.Params{
			Schema:        Schema, // Use the exported Schema
			RequestString: params.Query,
		})

		// Return the result or errors
		if len(result.Errors) > 0 {
			c.JSON(http.StatusBadRequest, result)
			return
		}

		c.JSON(http.StatusOK, result)
	})
}
