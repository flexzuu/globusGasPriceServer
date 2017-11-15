package gasPrice

import (
	"database/sql"
	"time"

	graphql "github.com/neelance/graphql-go"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	type Mutation {
		AddGasPrice(lastUpdated: Time!, e5: Float!, e10: Float!, superPlus: Float!, diesel: Float!, autogas: Float!): GasPrice!
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
		GasPrices: [GasPrice]!
	}
	
	type GasPrice {
		id: ID!
		lastUpdated: Date!
		e5: Float!
		e10: Float!
		superPlus: Float!
		diesel: Float!
		autogas: Float!
	}
	type Date {
		time: Time!
		human: String!
	}
	scalar Time
`

type gasPrice struct {
	ID          graphql.ID
	LastUpdated date
	E5          float64
	E10         float64
	SuperPlus   float64
	Diesel      float64
	Autogas     float64
}

func (gp *gasPrice) setID(id int64) {
	intID := int32(id)
	gp.ID.UnmarshalGraphQL(intID)
}
func (gp *gasPrice) setLastUpdated(lastUpdated time.Time) {
	gp.LastUpdated.Time.UnmarshalGraphQL(lastUpdated)
}

type date struct {
	Time graphql.Time
}

type Resolver struct {
	Database *sql.DB
}

func (r *Resolver) GasPrices() ([]*gasPriceResolver, error) {

	query := `SELECT "id", "lastUpdated", "e5", "e10", "superPlus", "diesel", "autogas" FROM "public"."gasPrices"  
			  ORDER BY "id" ASC`
	rows, err := r.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var gasPriceResolvers []*gasPriceResolver
	for rows.Next() {
		var gasPrice gasPrice
		var id int64
		var lastUpdated time.Time
		if err := rows.Scan(&id, &lastUpdated, &gasPrice.E5, &gasPrice.E10, &gasPrice.SuperPlus, &gasPrice.Diesel, &gasPrice.Autogas); err != nil {
			return nil, err
		}
		gasPrice.setID(id)
		gasPrice.setLastUpdated(lastUpdated)
		gasPriceResolvers = append(gasPriceResolvers, &gasPriceResolver{&gasPrice})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return gasPriceResolvers, nil
}

func (r *Resolver) AddGasPrice(args *struct {
	LastUpdated graphql.Time
	E5          float64
	E10         float64
	SuperPlus   float64
	Diesel      float64
	Autogas     float64
}) (*gasPriceResolver, error) {
	var id int64
	var lastUpdated time.Time
	var gasPrice gasPrice
	query := `
	INSERT INTO "gasPrices"("lastUpdated", "e5", "e10", "superPlus", "diesel", "autogas") 
	VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING "id", "lastUpdated", "e5", "e10", "superPlus", "diesel", "autogas";
	`
	err := r.Database.
		QueryRow(query, args.LastUpdated.Time, args.E5, args.E10, args.SuperPlus, args.Diesel, args.Autogas).
		Scan(&id, &lastUpdated, &gasPrice.E5, &gasPrice.E10, &gasPrice.SuperPlus, &gasPrice.Diesel, &gasPrice.Autogas)
	if err != nil {
		return nil, err
	}

	gasPrice.setID(id)
	gasPrice.setLastUpdated(lastUpdated)
	return &gasPriceResolver{&gasPrice}, nil
}

type gasPriceResolver struct {
	gasPrice *gasPrice
}

func (r *gasPriceResolver) ID() graphql.ID {
	return r.gasPrice.ID
}

func (r *gasPriceResolver) LastUpdated() *dateResolver {
	return &dateResolver{&r.gasPrice.LastUpdated}
}

type dateResolver struct {
	*date
}

func (r *dateResolver) Time() graphql.Time {
	return r.date.Time
}

func (r *dateResolver) Human() string {
	return "foobar"
}

func (r *gasPriceResolver) E5() float64 {
	return r.gasPrice.E5
}

func (r *gasPriceResolver) E10() float64 {
	return r.gasPrice.E10
}
func (r *gasPriceResolver) SuperPlus() float64 {
	return r.gasPrice.SuperPlus
}

func (r *gasPriceResolver) Diesel() float64 {
	return r.gasPrice.Diesel
}
func (r *gasPriceResolver) Autogas() float64 {
	return r.gasPrice.Autogas
}
