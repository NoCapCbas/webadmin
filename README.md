page: 56, 5.5 Get user detail example

# Development

Runs app with hot reload via air
```bash
./scripts/dev.sh
```
To make the script executable, use the chmod command:
```shell
chmod +x scripts/dev.sh
```

# Testing 
For testing, make sure development docker containers are running
If not, run the following command to start them
```bash
./scripts/dev.sh
```

## Seed Data
Seed data for postgres
```bash
docker exec -it <postgres-container-name> ./scripts/seed_data.sh
```
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

# Database
Default database is postgres, if you want to use mongo, you can ...

This code base is for educational purposes, 
sourced from "build saas app in go" by dominic st-pierre



# TODO
[] add a way to insert seed data
