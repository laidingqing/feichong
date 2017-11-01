package managers

import mgo "gopkg.in/mgo.v2"

const userCollectionName = "users"
const orderCollectionName = "orders"
const businessCollectionName = "business"
const orderTaxsCollectionName = "taxs"

//Database connection
var (
	mgoSession   *mgo.Session
	databaseName = "localDb"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	mgoSession, err := mgo.Dial("localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return mgoSession.Clone()
}

func executeQuery(collectionName string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	return s(c)
}
