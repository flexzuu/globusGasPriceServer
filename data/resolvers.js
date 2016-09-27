import { Kind } from 'graphql/language';

const resolvers = {
  Query: {
    allGasData(obj, {last = 0}, {connectors}) {
      return connectors.gasData.get(last);
    },
    gasData(obj, {id}, {connectors}) {
      return connectors.gasData.getOne(id);
    }
  },
  GasData: {
    id(obj){
      return obj['_id'];
    },
    lastUpdatedHuman(obj){
      return new Date(obj.lastUpdated).toString();
    }
  },
  Mutation: {
    addGasData(obj, args, {connectors}){
      return connectors.gasData.add(args);
    }
  },
  Date: {
    __parseValue(value) {
      return value; // value from the client
    },
    __serialize(value) {
      return value; // value sent to the client
    },
    __parseLiteral(ast) {
      if (ast.kind === Kind.INT) {
        return parseInt(ast.value, 10); // ast value is always in string format
      }
      return null;
    },
  }
};

export default resolvers;
