# Neo4jStore - gorilla/sessions

A session storage backend for [gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions) - [src](https://github.com/gorilla/sessions).

This implementation is based on the [postgresql implementation](https://github.com/antonlindstrom/pgstore). In fact, I copy the whole pgs backend implementation tests with a few modifications to fit the context of Neo4j.

[![Build Status](https://travis-ci.org/nathandao/neo4jstore.svg?branch=master)](https://travis-ci.org/nathandao/neo4jstore) [![GoDoc](https://godoc.org/github.com/nathandao/neo4jstore?status.svg)](https://godoc.org/github.com/nathandao/neo4jstore)

# Usage

```go
//...

const(
  DbUrl = "http://user:password@localhost:7474"
  SecretKey = "something very secret"
)

// Fetch new store.
store, err := NewNeo4jStore(DbUrl, []byte(SecretKey))
if (err != nil) {
  panic(err)
}

// Get a session.
session, err = store.Get(req, "session-key")
if err != nil {
  log.Error(err.Error())
}

// Add a value.
session.Values["foo"] = "bar"

// Save.
if err = sessions.Save(req, rsp); err != nil {
  t.Fatalf("Error saving session: %v", err)
}

// Delete session.
session.Options.MaxAge = -1
if err = sessions.Save(req, rsp); err != nil {
  t.Fatalf("Error saving session: %v", err)
}
```

## Testing

Before running the tests, make sure you have neo4j installed and running at ```localhost:7474``` with ```username: neo4j``` and ```password: foobar```

WARNING: make sure you are runny the test on a separate database. Since there is a clean up function that wipes out the whole database after each test.

## Thanks

I've stolen, borrowed and gotten inspiration from the other backends available - especially pgstore.

* [pgstore](https://github.com/antonlindstrom/pgstore)
* [redistore](https://github.com/boj/redistore)
* [mysqlstore](https://github.com/srinathgs/mysqlstore)
* [babou dbstore](https://github.com/drbawb/babou/blob/master/lib/session/dbstore.go)

Thank you all for sharing your code!
