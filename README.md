# Vending Machine App

Me trying to learn Go and Next.js with Surrealdb.

## Description

This application simulates a vending machine. Here are the key technologies
used:

### Backend

| Area      | Technology |
| --------- | ---------- |
| Language  | Go         |
| Framework | Gin        |
| Database  | SurrealDB  |

### Frontend

| Area             | Technology               |
| ---------------- | ------------------------ |
| Language         | TypeScript               |
| Framework        | React.js and Next.js     |
| UI Design        | Tailwind CSS and DailyUI |
| State Management | Nano Stores              |

## how to run

How to Run To run the application locally, use Docker Compose:

```bash
docker-compose up -d
```

The application will run at port 3000, and the REST API will run at port 8080.

Please remember to update the `.env` file in the root directory and the
`/frontend/.env`file as well. Next.js uses a separate `.env` file at build time.
