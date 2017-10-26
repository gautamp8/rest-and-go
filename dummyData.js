use musicstore;
var bulk = db.albums.initializeUnorderedBulkOp();
bulk.insert(   { title: "OK Computer", artist: "Radiohead", year: 1997, id: 1 });
bulk.insert(   { title: "The Queen is dead", artist: "The Smiths", year: 1986, id: 2 });
bulk.insert(   { title: "Be Here Now", artist: "Oasis", year: 1997, id: 3 });
bulk.insert(   { title: "Appetite for Destruction", artist: "Guns N Roses", year: 1987, id: 4 });
bulk.insert(   { title: "Back To Black", artist: "Amy Winehouse", year: 2006, id: 5 });
bulk.insert(   { title: "Hotel California", artist: "Eagles", year: 1976, id: 6 });
bulk.execute();