# Go URL Shortener — Production-Ready Backend API

A high-performance, persistent URL shortening service written in Go, featuring full CRUD operations, click analytics, and production-grade architecture.


Made for https://roadmap.sh/projects/url-shortening-service

**Status:** Complete and ready for deployment  
**Live demo:** Will be deployed on Render or Fly.io the moment I begin applying (one-click deployment with Docker)

## Features

- `POST   /shorten`          → Create short URL (JSON API)  
- `GET    /{code}`           → 301 Permanent Redirect to original URL  
- `GET    /stats/{code}`     → Detailed analytics (click count, creation time, original URL)  
- `PUT    /{code}`           → Update destination URL  
- `DELETE /{code}`           → Remove URL  
- Click tracking on every redirect  
- Collision-resistant 6-character codes  
- Persistent storage with PostgreSQL  
- Type-safe database access via sqlc (zero runtime reflection)

## Technology Stack

- Go 1.23+
- PostgreSQL
- pgx/v5 driver
- sqlc (compile-time SQL code generation)
- net/http (modular and ready for Chi/Fiber ready)

## Local Development

```bash
# Clone the repository
git clone https://github.com/Asendar1/go-url-shortener.git
cd go-url-shortener

# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Create database
psql postgres -c "CREATE DATABASE urlshortener;"
psql postgres -c "CREATE USER dev WITH PASSWORD 'dev';"

# Generate type-safe queries
sqlc generate

# Run
go run .
```

# Create short URL
curl -X POST http://localhost:8080/shorten -d "url=https://github.com"

# Redirect (open in browser)
http://localhost:8080/Xy9pK2

# View statistics
curl http://localhost:8080/shorten/stats/Xy9pK2
