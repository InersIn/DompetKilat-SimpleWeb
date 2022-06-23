# DompetKilat-SimpleWeb
Pre-test Golang API

# Technology
1. Golang Echo
2. Authentication JWT
3. Database Mysql
4. Docker Containerizing

# Instalation
```bash
git clone https://github.com/InersIn/DompetKilat-SimpleWeb
```

```
cd DompetKilat-SimpleWeb
```

```
docker-compose build
```

```
docker-compose up -d
```

# Usage
## Account
---
### Register
#### register new user
---
url: `/api/register`

method: `POST`

request body:
```json
{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

response:
```json
{
  "status": "string",
  "message": "string"
}
```

### Login
#### User login and generate auth token
---
url: `/api/login`

method: `POST`

request body:
```json
{
  "username": "string",
  "password": "string"
}
```

response:
```json
{
  "token": "string",
  "type": "string"
}
```

## Finance
### Financing
#### Create financing
---
url: `/api/service/createFinances`

method: `POST`

request body:
```json
{
  "name": "string",
  "count": 0,
  "sub": "string"
}
```

response:
```json
{
  "status": "string",
  "message": "string"
}
```
#### Get all financing from current user
---
url: `/api/service/createFinances`

method: `GET`

response:
```json
{
  "status": "string",
  "data": [{
    "finance_id": 0,
    "name": "string",
    "count": 0,
    "sub": "string",
    "user_id": 0
  }]
}
```
## For full documentation you can see [here](https://app.swaggerhub.com/apis/SAPU2776/dompet-kilat_simple_web/1.0.0) or you can open file `apispec.json` in vscode
