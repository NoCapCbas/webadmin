page: 46


# Database
Default database is postgres, if you want to use mongo, you can ...

# Testing 
## Unit Testing
Once you have your project setup, you can run the following command to run all the unit tests:
```bash
go test ./...
```

## Integration Testing
```bash
go test -tags integration ./...
```

## Integration Testing with MongoDB
```bash
go test --tags="integration mongo" ./data
```

## Integration Testing with Postgres
```bash
go test -tags="integraion" ./data
```





