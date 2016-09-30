import { Kind } from 'graphql/language';

const resolvers = {
  Query: {
    allGasData(obj, {last = 0}, {connectors}) {
      return connectors.gasData.get(last);
    },
    gasData(obj, {id}, {connectors}) {
      return connectors.gasData.getOne(id);
    },
    lastGasData(obj, _ , {connectors}) {
      return connectors.gasData.getLast();
    },
  },
  GasData: {
    id(obj){
      return obj['_id'];
    },
    lastUpdatedHuman(obj){
      return new Date(obj.lastUpdated);
    }
  },
  Mutation: {
    addGasData(obj, args, {connectors}){
      return connectors.gasData.add(args);
    }
  },
  DateHuman: {
    day(obj){
      return obj.getDate();
    },
    month(obj){
      return obj.getMonth()+1;
    },
    year(obj){
      return obj.getFullYear();
    },
    hours(obj){
      return obj.getHours();
    },
    minutes(obj){
      return obj.getMinutes();
    },
    seconds(obj){
      return obj.getSeconds();
    },
    offset(obj){
      return obj.getTimezoneOffset();
    },
    readable(obj){
      return obj.toString();
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
