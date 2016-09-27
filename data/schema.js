const typeDefinitions = `

scalar Date

type GasData {
  id: ID!
  lastUpdated: Date!
  lastUpdatedHuman: String
  e5: Float
  e10: Float
  superPlus: Float
  diesel: Float
  autogas: Float
}

type Query {
  allGasData(last: Int): [GasData!]
  gasData(id: ID!): GasData!
}

type Mutation {
  addGasData(lastUpdated: Date!, e5: Float, e10: Float, superPlus: Float, diesel: Float, autogas: Float): GasData!
}

schema {
  query: Query
  mutation: Mutation
}
`;

export default [typeDefinitions];
