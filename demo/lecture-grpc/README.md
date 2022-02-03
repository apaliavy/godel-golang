# gRPC Communication example 

Here I have a tiny mono-repo with 3 services: 
- gateway 
- auth 
- users

The idea is - we have an HTTP middleware in `gateway` service which calls `auth` service to validate a request. 
Once we received confirmation from `auth` service we make a gRPC request to the users service. 

## How to Run 

Just use the Makefile: 
```
make run
```

Then you can send http requests to the following endpoints: 
- POST http://127.0.0.0:8080/users 
- GET http://127.0.0.0:8080/users
- GET http://127.0.0.0:8080/users/1

## Notes

Those ports are needed: 
- 8080 (gateway service)
- 9000 (for auth service)
- 9001 (for users service)

