import loki from 'lokijs';
const db = new loki(`${__dirname}/gasApp.db`,
      {
        autoload: true,
        autoloadCallback : loadHandler,
        autosave: true,
        autosaveInterval: 10000,
      });
let gasDataCollection;
let allGasDataView;

    function loadHandler() {
      // if database did not exist it will be empty so I will intitialize here
      gasDataCollection = db.getCollection('gasData',{unique: 'lastUpdated'})
      if (gasDataCollection === null) {
        gasDataCollection = db.addCollection('gasData',{unique: 'lastUpdated'});
      }
      allGasDataView = gasDataCollection.addDynamicView('orderByLastUpdated');
      allGasDataView.applySimpleSort('lastUpdated');
    }

const connectors = {
  gasData: {
    getOne: (id) => gasDataCollection.get(id),
    get: (last) => allGasDataView.data().slice(-1*last),
    add: (item) => gasDataCollection.insert(item),
  }
};

export default connectors;
