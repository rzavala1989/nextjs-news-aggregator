* UPDATE 8/2/2024 - Backend code complete, working on NextJS code to put it all together in terms of best practices and adding features*

# News Aggregator Application

## Overview

The News Aggregator Application is a robust platform that consolidates news articles from various sources, provides trending search queries, and delivers personalized news recommendations. It consists of a backend built with Go for fetching and processing news, a frontend built with Next.js for user interaction, and a recommendation service in Python. The application uses JWT for authentication and stores data in a SQLite database.

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
- **Recommendation Engine**: Python (Flask)

## Project Structure

```
news-aggregator-app/
├── frontend/
│ ├── .idea/
│ ├── .next/
│ ├── node_modules/
│ ├── public/
│ ├── src/
│ │ ├── app/
│ │ │ ├── login/
│ │ │ │ └── page.jsx
│ │ │ ├── news/
│ │ │ │ └── page.jsx
│ │ │ ├── register/
│ │ │ │ └── page.jsx
│ │ │ └── page.jsx
│ │ ├── components/
│ │ │ ├── Layout.jsx
│ │ │ ├── Navbar.jsx
│ │ │ ├── Recommendations.jsx
│ │ │ └── TrendingSearches.jsx
│ │ ├── hooks/
│ │ │ └── useToast.js
│ │ ├── store/
│ │ │ ├── newsStore.js
│ │ │ └── userStore.js
│ │ └── styles/
│ │ └── globals.css
│ ├── .env.local
│ ├── .eslintrc.json
│ ├── .gitignore
│ ├── jsconfig.json
│ ├── next.config.js
│ ├── package.json
│ ├── package-lock.json
│ ├── postcss.config.mjs
│ ├── tailwind.config.js
│ └── README.md
├── news-aggregator/
│ ├── .idea/
│ ├── cmd/
│ │ └── server/
│ │ └── main.go
│ ├── data/
│ │ ├── database.db
│ │ └── recommendations.csv
│ ├── internal/
│ │ ├── auth/
│ │ │ ├── handler.go
│ │ │ ├── jwt.go
│ │ │ └── middleware.go
│ │ ├── db/
│ │ │ ├── articles.go
│ │ │ ├── db.go
│ │ │ ├── recommendation.go
│ │ │ ├── trending.go
│ │ │ └── user.go
│ │ ├── models/
│ │ │ ├── article.go
│ │ │ ├── py_article.go
│ │ │ ├── trending.go
│ │ │ └── user.go
│ │ ├── news/
│ │ │ ├── fetch.go
│ │ │ ├── handler.go
│ │ │ ├── sync.go
│ │ │ └── trending.go
│ │ └── user/
│ │ └── handler.go
│ ├── .env
│ ├── .env_sample
│ ├── .gitignore
│ ├── app.db
│ ├── go.mod
│ ├── go.sum
│ ├── identifier.sqlite
│ ├── README.md
│ └── start_servers.sh
├── news_aggregator_recommendations/
│ ├── .idea/
│ ├── .venv/
│ ├── model/
│ │ └── model.py
│ ├── models/
│ │ ├── add_dummy_data.py
│ │ ├── recreate_articles_table.py
│ ├── static/
│ ├── templates/
│ ├── .env
│ ├── app.py
│ ├── requirements.txt
│ └── README.md
├── .gitignore
└── README.md
```


## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)
- [Node.js](https://nodejs.org/) (version 12+)
- [Python](https://www.python.org/downloads/) (version 3.6+)
- SQLite3

### Installation

**Clone the repository**:

```bash
git clone https://github.com/yourusername/news-aggregator-app.git
cd news-aggregator-app


IMPORTANT: Run command to make start script execuatable (this runs both servers):
```bash
chmod +x start_servers.sh
```

Then run the script:
```bash
./start_servers.sh
```

## Manual Backend Setup (Go)
- Navigate to the `news-aggregator` directory
```
cd news-aggregator
```

- Install dependencies
```
go mod tidy
```

- Run the Go server
```
go run cmd/server/main.go
```

## Manual Frontend Setup (Flask)

- Navigate to the `news_aggregator_recommendations` directory
```
(from root)

cd news_aggregator_recommendations
```

- Create a virtual environment
```
python3 -m venv .venv
source .venv/bin/activate
```

- Install dependencies
```
pip install -r requirements.txt
```

- Run the Flask server
```
python app.py
```

## Run the Frontend (Next.js)
- Navigate to the `frontend` directory
```
cd frontend
```

- Install dependencies
```
npm install
```

- Set up environment variables for local development
```
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_RECOMMENDATION_URL=http://localhost:5000

NEXT_PUBLIC_CURRENTS_API_KEY=your_currents_api_key
NEXT_PUBLIC_THENEWSAPI_API_KEY=your_thenewsapi_api_key
NEXT_PUBLIC_NEWSDATAIO_API_KEY=your_newsdataio_api_key
NEXT_PUBLIC_MEDIASTACK_API_KEY=your_mediastack_api_key
NEXT_PUBLIC_SERPAPI_API_KEY=your_serpapi_key
NEXT_PUBLIC_JWT_SECRET=your_jwt_secret

```

- Run the Next.js server
```
npm run dev
```

*MORE DETAILS ARE IN THE README.md FILES IN THE RESPECTIVE DIRECTORIES*
