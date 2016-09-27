import mongo from 'then-mongo';
import ObjectID from 'bson-objectid';
const db = mongo('mongodb://localhost/main', ['gasData']);
const gasData = db.gasData;

const connectors = {
  gasData: {
    getOne: (id) => gasData.findOne({_id: ObjectID(id)}),
    get: (last) => gasData.find({}).sort({lastUpdated: -1}).limit(last).toArray(),
    getLast: () => gasData.find({}).sort({lastUpdated: -1}).limit(1).toArray().then((data)=> data[0]),
    add: (item) => gasData.insert(item),
  }
};

export default connectors;
