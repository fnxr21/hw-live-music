# 🎶 Real-Time Music Request System

A living, breathing music request platform for live venues—where audiences don’t just listen, they participate.

Customers request songs straight from their tables. Bands curate the flow. Everyone stays in sync in real time.

---


## 🏗️ Architecture Overview

This system is built on an **event-driven architecture**, designed for immediacy and reliability.

Every action ripples through the system in real time.

```
[User Action] → [Database Change] → [Trigger Event]
     → [WebSocket Broadcast] → [All Clients Update]
```

### High-Level Components

```
┌─────────────────┐    HTTP / WebSocket    ┌─────────────────┐
│                 │◄─────────────────────►│                 │
│   Frontend      │                       │    Backend      │
│   (React)       │                       │     (Go)        │
│                 │                       │                 │
└─────────────────┘                       └────────┬────────┘
                                                   │
                                                   │ PostgreSQL
                                                   │ LISTEN / NOTIFY
                                            ┌──────▼──────┐
                                            │             │
                                            │  Database   │
                                            │ (Postgres)  │
                                            │             │
                                            └─────────────┘
```

### Event Flow


```
1. Customer or Admin performs an action
2. Frontend sends API request
3. Backend updates PostgreSQL
4. Database trigger fires
5. pg_notify() emits an event
6. Backend listens via LISTEN
7. WebSocket broadcasts update
8. All connected clients update instantly
```

The result: **zero refreshes, zero confusion**.

---

## ✨ Key Features of Event-Driven Design

### Benefits

✅ **Real-time Updates** — Perubahan muncul seketika tanpa perlu polling

✅ **Database as Source of Truth** — Semua alur perubahan berasal dari trigger database

✅ **Scalable** — PostgreSQL menangani distribusi event dengan efisien

✅ **Reliable** — Jaminan ACID dari database menjaga konsistensi data

✅ **Simple** — Tidak perlu message queue atau event bus tambahan

✅ **Efficient** — Hanya data yang berubah yang dikirim ke klien

---

### How It’s Different from Polling

#### Traditional Polling

```
Frontend → (setiap 2 detik) → Backend → Database → Backend → Frontend
```

❌ Boros bandwidth
❌ Update terlambat (tergantung interval polling)
❌ Beban server tinggi

---

#### Event-Driven dengan Database Triggers

```
Database Change → Trigger → Notify → WebSocket → All Clients
```

✅ Update instan
✅ Bandwidth minimal
✅ Beban server rendah

---


## 🚀 Getting Started
Before you begin, ensure you have the following installed:
- **Docker** and **Docker Compose**
- **Node.js** (v18 or higher) and **npm**
- **Go** (v1.21 or higher)
- **golang-migrate** CLI tool

Install golang-migrate:
```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Windows
choco install golang-migrate
```

## ⚙️ Environment Setup

### Step 1: Database Environment

Create `database/docker/.env` file:

```env
APP_NAME=live-music

PGSQL_POSTGRES_USER=postgres
PGSQL_POSTGRES_PASSWORD=pwpostgres!
PGSQL_POSTGRES_DB=live_music
PGSQL_DIR=./data
PGSQL_INIT=./init
PGSQL_PORT=5432
```

### Step 2: Backend Environment

Create `backend/.env` file:

```env
DB_USER=postgres
DB_PASSWORD=pwpostgres!
DB_HOST=localhost
DB_PORT=5432
DB_NAME=live_music

PORT=8080
```

### Step 3: Frontend Environment

Create `frontend/.env.local` file:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## 🚀 Getting Started

### Option 1: Using Make Commands (Recommended)

```bash
# 1. Start PostgreSQL database
make db-up

# 2. Run database migrations
make migrate-up

# 3. Start backend (in a new terminal)
make run-go

# 4. Start frontend (in another new terminal)
make run-next
```

### Option 2: Manual Setup

#### 1. Start the Database

```bash
docker-compose -f database/postgres/docker-compose.yml up -d
```

Verify the database is running:
```bash
docker ps
```

#### 2. Run Database Migrations

```bash
migrate -path ./backend/db/migrations -database "postgres://postgres:pwpostgres!@localhost:5432/live_music?sslmode=disable" up
```

#### 3. Start the Backend Server

Open a new terminal window:

```bash
cd backend
go mod tidy
go run main.go
```

The backend API will be available at `http://localhost:8080`

#### 4. Start the Frontend Development Server

Open another new terminal window:

```bash
cd frontend
npm install
npm run dev
```

The frontend will be available at `http://localhost:3000`

## 🛠️ Available Make Commands

| Command | Description |
|---------|-------------|
| `make db-up` | Start PostgreSQL database container |
| `make db-down` | Stop and remove database container |
| `make migrate-up` | Run all pending database migrations |
| `make migrate-down` | Rollback the last migration |
| `make run-go` | Start Go backend server |
| `make run-next` | Start Next.js frontend server |
| `make reset` | ⚠️ **DANGER**: Full reset - deletes all data |

## 🗄️ Database Management

### View Database Logs
```bash
docker-compose -f database/postgres/docker-compose.yml logs -f
```

### Access PostgreSQL CLI
```bash
docker exec -it <container_name> psql -U postgres -d live_music
```

### Stop Database
```bash
make db-down
```

### Reset Everything (⚠️ Warning: Deletes All Data)
```bash
make reset
```

This will:
- Stop and remove the database container
- Delete all data from `database/postgres/data`
- You'll need to run `make db-up` and `make migrate-up` again

## 📍 Application URLs

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Database**: localhost:5432

## 🔧 Troubleshooting

### Database connection errors
- Ensure PostgreSQL container is running: `docker ps`
- Check credentials in `backend/.env` match `database/docker/.env`
- Verify port 5432 is not in use: `lsof -i :5432` (macOS/Linux)

### Migration errors
- Ensure database is running before migrating
- Check migration files in `backend/db/migrations`
- To fix failed migrations, you may need to run `make migrate-down` then `make migrate-up`

### Port already in use
- Frontend (3000): Check if another Next.js app is running
- Backend (8080): Check if another service is using port 8080
- Database (5432): Check if PostgreSQL is installed locally







---

## 🎼 Genius API Integration

Songs are powered by the Genius API for rich metadata and discovery.

### Search Songs

```
GET /api/genius/search?q=Bohemian+Rhapsody
```

### Get Song Details

```
GET /api/genius/songs/:id
```

### Sample Response

```json
{
  "success": true,
  "data": {
    "id": 12345,
    "title": "Bohemian Rhapsody",
    "artist": "Queen",
    "album": "A Night at the Opera",
    "release_date": "1975-10-31",
    "duration": 354,
    "genius_url": "https://genius.com/Queen-bohemian-rhapsody-lyrics",
    "thumbnail": "https://images.genius.com/..."
  }
}
```

## 🎵 Final Note

This project is built with care, clarity, and a love for live music.

It’s not just software—it’s rhythm, timing, and trust between performers and their audience.

If you need help setting up or extending the system, feel free to open an issue on GitHub.

**Enjoy the music. Let the room decide the next song.** 🎸🎹🥁
