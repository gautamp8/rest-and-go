use store;
var bulk = db.products.initializeUnorderedBulkOp();
bulk.insert(   { title: "Apple iMac Pro", image: "http:://example.com/p1.jpg", price: 5000, rating: 4, id: 1 });
bulk.insert(   { title: "Google Pixel 2", image: "http:://example.com/p2.jpg", price: 2000, rating: 5, id: 2 });
bulk.insert(   { title: "Apple iPhone X", image: "http:://example.com/p3.jpg", price: 3000, rating: 5, id: 3 });
bulk.insert(   { title: "Google Chromebook", image: "http:://example.com/p4.jpg", price: 4000, rating: 5, id: 4 });
bulk.insert(   { title: "Microsoft Holo Lens", image: "http:://example.com/p5.jpg", price: 1000, rating: 4, id: 5 });
bulk.insert(   { title: "Samsung Galaxy S8", image: "http:://example.com/p6.jpg", price: 3000, rating: 3, id: 6 });
bulk.execute();