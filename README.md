# DistriCache
A proof-of-concept distributed cache system implemented in golang. 
A redis like system with eventual consistency.

The consensus of the roles in the cluster will not be implemented for now.

## Methods 
- [x] Set
- [x] Get
- [x] Delete
- [x] Has

## Usage
Start the server with the following command
```shell
go run cmd/cache/main.go
```


