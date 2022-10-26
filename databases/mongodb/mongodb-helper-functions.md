# User-Defined Helper Functions for MongoDB in Go

## Connection URI

``` Go 
// This is a user defined method that returns
// a mongo.Client, context.Context,
// context.CancelFunc and error.
// mongo.Client will be used for further database
// operation. context.Context will be used set
// deadlines for process. context.CancelFunc will
// be used to cancel context and resource
// associated with it.
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
 
    ctx, cancel := context.WithTimeout(
                            context.Background(), 
                            30 * time.Second
                        )

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    
    return client, ctx, cancel, err
}
```

************************

## Close Connection

``` Go
func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
    defer cancel()
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()
}
```

************************

## Fine One

``` Go

```

************************

## Find Many

``` Go
// query is user defined method used to query MongoDB,
// that accepts mongo.client,context, database name,
// collection name, a query and field.
 
//  database name and collection name is of type
// string. query is of type interface.
// field is of type interface, which limits
// the field being returned.
 
// query method returns a cursor and error.
func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
 
    // select database and collection.
    collection := client.Database(dataBase).Collection(col)
     
    // collection has an method Find,
    // that returns a mongo.cursor
    // based on query and field.
    result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
    return
}
```

************************


## Insert One

``` Go
// insertOne is a user defined method, used to insert
// documents into collection returns result of InsertOne
// and error if any.
func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
 
    // select database and collection ith Client.Database method
    // and Database.Collection method
    collection := client.Database(dataBase).Collection(col)
     
    // InsertOne accept two argument of type Context
    // and of empty interface  
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}
```

************************

## Insert Many

``` Go
// insertMany is a user defined method, used to insert
// documents into collection returns result of
// InsertMany and error if any.
func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {
 
    // select database and collection ith Client.Database
    // method and Database.Collection method
    collection := client.Database(dataBase).Collection(col)
     
    // InsertMany accept two argument of type Context
    // and of empty interface  
    result, err := collection.InsertMany(ctx, docs)
    return result, err
}
```

************************

## Update One

``` Go
// UpdateOne is a user defined method, that update
// a single document matching the filter.
// This methods accepts client, context, database,
// collection, filter and update filter and update
// is of type interface this method returns
// UpdateResult and an error if any.
func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
                
    // select the database and the collection
    collection := client.Database(dataBase).Collection(col)
     
    // A single document that match with the
    // filter will get updated.
    // update contains the filed which should get updated.
    result, err = collection.UpdateOne(ctx, filter, update)
    return
}
```

************************

## Update Many

``` Go
// UpdateMany is a user defined method, that update
// a multiple document matching the filter.
// This methods accepts client, context, database,
// collection, filter and update filter and update
// is of type interface this method returns
// UpdateResult and an error if any.
func UpdateMany(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
 
    // select the database and the collection
    collection := client.Database(dataBase).Collection(col)
     
    // All the documents that match with the filter will
    // get updated.
    // update contains the filed which should get updated.
    result, err = collection.UpdateMany(ctx, filter, update)
    return
}
```

************************

## Delete One

``` Go
// deleteOne is a user defined function that delete,
// a single document from the collection.
// Returns DeleteResult and an  error if any.
func deleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
 
    // select document and collection
    collection := client.Database(dataBase).Collection(col)
     
    // query is used to match a document  from the collection.
    result, err = collection.DeleteOne(ctx, query)
    return
}
```

************************

## Delete Many

``` Go
// deleteMany is a user defined function that delete,
// multiple documents from the collection.
// Returns DeleteResult and an  error if any.
func deleteMany(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
 
    // select document and collection
    collection := client.Database(dataBase).Collection(col)
     
     // query is used to match  documents  from the collection.
    result, err = collection.DeleteMany(ctx, query)
    return
}
```

************************

