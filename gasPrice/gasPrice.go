package gasPrice

import (
	"context"

	driver "github.com/arangodb/go-driver"
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
type date struct {
	Time graphql.Time
}

type Resolver struct {
	Client *driver.Client
}

func (r *Resolver) GasPrices() ([]*gasPriceResolver, error) {
	c := *r.Client
	ctx := context.Background()
	db, err := c.Database(ctx, "_system")
	if err != nil {
		return nil, err
	}

	query := "FOR d IN gas RETURN d"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	var l []*gasPriceResolver
	for {
		var gasPrice gasPrice
		meta, err := cursor.ReadDocument(ctx, &gasPrice)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		gasPrice.ID = graphql.ID(meta.Key)
		l = append(l, &gasPriceResolver{&gasPrice})
	}
	return l, nil
}

func (r *Resolver) AddGasPrice(args *struct {
	LastUpdated graphql.Time
	E5          float64
	E10         float64
	SuperPlus   float64
	Diesel      float64
	Autogas     float64
}) (*gasPriceResolver, error) {
	c := *r.Client
	ctx := context.Background()
	db, err := c.Database(ctx, "_system")
	if err != nil {
		return nil, err
	}
	col, err := db.Collection(ctx, "gas")
	if err != nil {
		return nil, err
	}
	date := date{
		Time: args.LastUpdated,
	}
	gasPrice := gasPrice{
		LastUpdated: date,
		E5:          args.E5,
		E10:         args.E10,
		SuperPlus:   args.SuperPlus,
		Diesel:      args.Diesel,
		Autogas:     args.Autogas,
	}
	meta, err := col.CreateDocument(ctx, gasPrice)
	if err != nil {
		return nil, err
	}
	gasPrice.ID = graphql.ID(meta.Key)
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
