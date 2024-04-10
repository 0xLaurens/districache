# DistriCache
DistriCache 🌐 is a proof-of-concept distributed in-memory key/value cache implemented implemented using go. It provides several methods for storing, retrieving and deleting values. DistriCache uses a custom byte based communication protocol over TCP.

## Features
- 📦 Key/Value data storage
- 🎚️ SET, GET and DEL operations
- ⏰ Optional TTL for Keys.
- 🖥️ Leader - Follower model
- 🤖 Core functions tested
## Running

Start the server with the following command
```shell
make server
```

Run a test client
```shell
make client
```

Running test
```shell
make test
```

Running test + benchmark
```shell
make bench
```
## Methods
The methods that are available over TCP
### SET
Command for inserting values into the cache
```TCP
SET key value [ttl]
```
#### Options
**TTL** - how long a key/value pair is valid before it's deleted. The TTL is specified in milliseconds. _Default value is 0 -> no expiration_
#### Response
**OK**, the request was successful

### GET
Command for retrieving values from the cache
```TCP
GET key
```
#### Response
**value**, the request was successful
**(nil)**, no key found


## DEL
Command for deleting values in the cache
```TCP
DEL key
```
#### Response
**OK**, the request was successful.

