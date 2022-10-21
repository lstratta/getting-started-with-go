# Connecting to MongoDB and Core Functionality

## Contents
1. [Database Connection](#database-connection)


***************************

## Database Connection

Open connection, then defer disconnection once the query has been completed

``` Go
// Load dotenv data for MONGODB_URI
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Assign URI value; if empty, log.Fatal()
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}


	// Connect to MongoDB; Perform query; Disconnect once complete
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// DB disconnect deferred to the end of the function
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

```

***************************

## Find One

``` Go
// begin findOne
	coll := client.Database("sample_mflix").Collection("movies")
	
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", "The Room"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	// end findOne

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
```

## Find Multiple


***************************

## Insert Document



***************************

## Insert Multiple