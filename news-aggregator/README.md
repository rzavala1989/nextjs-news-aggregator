# News Aggregator Application

## Overview

The News Aggregator Application is a robust platform that consolidates news articles from various sources, provides trending search queries, and delivers personalized news recommendations. It consists of a backend built with Go for fetching and processing news, and a frontend built with Next.js for user interaction. The application uses JWT for authentication and stores data in a SQLite database.

## Features

- **User Authentication**: Register and login with secure JWT authentication.
- **News Aggregation**: Fetch news articles from multiple sources concurrently.
- **Trending Searches**: Display trending search queries updated daily.
- **Personalized Recommendations**: Provide personalized news recommendations based on user preferences.
- **Responsive Design**: User-friendly interface built with Next.js.

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: Next.js
- **Database**: SQLite
- **Authentication**: JWT (JSON Web Tokens)
- **APIs**: Currents API, TheNewsAPI, NewsData.io, MediaStack, SerpAPI

## Project Structure
```
news-aggregator-app/
├── cmd/
│ └── server/
│ └── main.go
├── internal/
│ ├── auth/
│ │ ├── handler.go
│ │ ├── jwt.go
│ │ └── middleware.go
│ ├── db/
│ │ ├── db.go
│ │ ├── trending.go
│ │ └── user.go
│ ├── models/
│ │ ├── trending.go
│ │ └── user.go
│ ├── news/
│ │ ├── fetch.go
│ │ ├── handler.go
│ │ └── trending.go
│ └── user/
│ └── handler.go
├── .env
└── README.md
```


## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)
- SQLite3

### Installation

**Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/news-aggregator-app.git
   cd news-aggregator-app
   ```
### Set up environment variables

```bash
CURRENTS_API_KEY=your_currents_api_key
THENEWSAPI_API_KEY=your_thenewsapi_api_key
NEWSDATAIO_API_KEY=your_newsdataio_api_key
MEDIASTACK_API_KEY=your_mediastack_api_key
SERPAPI_API_KEY=your_serpapi_key
JWT_SECRET=your_jwt_secret
```

### Install dependencies

```bash
go mod tidy
```

### Run the Go server

```bash
go run cmd/server/main.go
```