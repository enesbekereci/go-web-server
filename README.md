# WORK IN PROGRESS

## How to Run
1. go run dev

## How to Build and Run Container
1. Run Docker Daemon
2. docker build --tag docker-go-app .
3. docker run -p 8080:8080 docker-go-app

Example Queries for Testing:
```bash
curl --location 'http://127.0.0.1:8080'
curl --location 'http://127.0.0.1:8080' --header 'Content-Type: application/json' --data '{"key1":"data","key2":"data"}'
curl --location --request DELETE 'http://127.0.0.1:8080'
```
## API

### 🔵GET /
Reserved for checking if the service is available.

### 🟡POST /
Description: 
Return: 
#### Request
Example:
```json
{
    "key1":"value",
    "key2":"value"
}
```
#### Response
Example:
```json
{
    "result": "resut"
}
```
#### Alias Fields
1. key1: description

### 🔵GET /

### 🔴DELETE /



## Database Tables


### devices

|1|2|3|4|5|
|---|---|---|---|---|
|a|b|c|d|e|

