package db

const (
	mongoDbConnectionString string = "mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority"
)

// Error messages
const (
	errorMongoConnFailMsg string = "connection with mongo db failed with error: %s"
	invColNameMsg         string = "invalid collection %s"
)
