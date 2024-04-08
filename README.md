# DistriCache
A proof-of-concept distributed cache system implemented in golang. 
A redis like system with eventual consistency.

The consensus of the roles in the cluster will not be implemented for now.
## Running
Start the server with the following command
```shell
go run cmd/cache/main.go
```

A testing client can be started using
```shell
go run cmd/client/main.go
```

## Commands 
The syntax for the implemented methods
### GET
```
GET key
```
#### Response
**value**, the request was successful 

**(nil)**, no key found

### SET 
```
SET key value [TTL]
```
#### options
**TTL** is an optional parameter for setting the expiration defined in milliseconds

#### response
**OK**, the request was successful

### DELETE
delete a key
```
DEL key
```

#### response
**OK**, the request was successful


