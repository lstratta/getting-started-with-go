# User-Defined Helper Functions for MongoDB in Go

## Connection URI

``` Go 
func connect(uri string)
(*mongo.Client, context.Context, context.CancelFunc, error) {
 
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
func close(client *mongo.Client, ctx context.Context,
           cancel context.CancelFunc) {
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
func query(client *mongo.Client, ctx context.Context,
dataBase, col string, query, field interface{})
(result *mongo.Cursor, err error) {
 
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
func insertOne
(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{})
(*mongo.InsertOneResult, error) {
 
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
func insertMany
(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{})
(*mongo.InsertManyResult, error) {
 
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

## Insert One

``` Go

```

************************

## Insert Many

``` Go

```

************************