page: 56, 5.5 Get user detail example

# Development

Runs app with hot reload via air
```bash
make dev
```

## Development Testing 
For testing, make sure development docker containers are running
If not, run the following command to start them
```bash
make dev
```

### Unit Testing
Once you have your project setup, you can run the following command to run all the unit tests:
```bash
go test ./...
```

### Integration Testing
```bash
go test -tags integration ./...
```

### Integration Testing with MongoDB
```bash
go test --tags="integration mongo" ./data
```

### Integration Testing with Postgres
```bash
go test -tags="integraion" ./data
```

# Database
Default database is mongo.

This code base is for educational purposes, 
sourced from "build saas app in go" by dominic st-pierre