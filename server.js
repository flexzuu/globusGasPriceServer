import express from 'express';

import typeDefs from './data/schema';
import mocks from './data/mocks';
import resolvers from './data/resolvers';
import connectors from './data/connectors';

// NEW or changed imports:
import { apolloExpress, graphiqlExpress } from 'apollo-server';
import { makeExecutableSchema, addMockFunctionsToSchema } from 'graphql-tools';
import bodyParser from 'body-parser';

const GRAPHQL_PORT = 8080;

const graphQLServer = express();

const schema = makeExecutableSchema({
  typeDefs,
  resolvers,
  logger: { log: (e) => console.log(e) },
});

addMockFunctionsToSchema({
  schema,
  mocks,
  preserveResolvers: true,
});

// `context` must be an object and can't be undefined when using connectors
graphQLServer.use('/graphql', bodyParser.json(), apolloExpress({
  schema,
  context: {
    connectors,
  }, //at least(!) an empty object
}));

graphQLServer.use('/graphiql', graphiqlExpress({
  endpointURL: '/graphql',
}));

graphQLServer.listen(GRAPHQL_PORT, () => console.log(
  `GraphQL Server is now running on http://localhost:${GRAPHQL_PORT}/graphql`
));
