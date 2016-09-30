const typeDefinitions = `

scalar Date

type DateHuman {
  day: Int!
  dayOfWeek: String!
  month: Int!
  year: Int!

  hours: Int!
  minutes: Int!
  seconds: Int!

  offset: Int!
  readable: String!
}

type GasData {
  id: ID!
  lastUpdated: Date!
  lastUpdatedHuman: DateHuman!
  e5: Float
  e10: Float
  superPlus: Float
  diesel: Float
  autogas: Float
}

type Query {
  allGasData(last: Int): [GasData!]
  gasData(id: ID!): GasData
  lastGasData: GasData
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
