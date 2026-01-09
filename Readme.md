# 🎶 Real-Time Music Request System

A living, breathing music request platform for live venues—where audiences don’t just listen, they participate.

Customers request songs straight from their tables. Bands curate the flow. Everyone stays in sync in real time.

---

## 📋 Problem Statement

Live music venues thrive on energy and momentum—but song requests are often chaotic, manual, or disruptive.

This system creates a clean, real-time bridge between:

* **Customers**, who want to request songs effortlessly
* **Bands/Admins**, who need full control over what gets played and when
* **The Venue**, which benefits from smoother operations and happier guests

No shouting. No scraps of paper. Just music, moving at the speed of the room.

---

## ✨ Core Features

### For Customers

* Browse a curated song catalog
* Search by title, artist, or genre
* Request songs directly from table tablets
* Track request status in real time
* View the upcoming queue and estimated wait time
* Get notified when their song is up next

### For Admins / Bands

* Live dashboard of all song requests
* Approve, reject, or reorder requests instantly
* Drag-and-drop playlist management
* Mark songs as played and maintain history
* Add notes for band members
* Import and manage songs via Genius API

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
```bash
# Start development server
docker-compose -f ./database/docker-compose.yml up -d



```

```bash
cd my-app

# Install dependencies
npm install

# Edit .env.local as needed

# Start development frontend
npm run dev
```

```bash
cd backend

# Install dependencies
go mod tidy

# Start development server
go run main.go
```

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

---

## 🧍 Customer Flow

### 1. Access Table Interface

* Tablet loads venue-specific UI

### 2. Browse & Discover

* Search songs by title or artist
* Filter by genre or category
* View detailed song metadata

### 3. Submit a Request

* Tap **Request** on a song
* Optionally add name or notes
* Status starts as **⏳ Pending**

### 4. Track Status in Real Time

* ⏳ **Pending** — Awaiting approval
* ✅ **Approved** — Added to the queue
* ❌ **Rejected** — Declined by admin
* 🎵 **Played** — Performed live

### 5. View the Queue

* See all approved upcoming songs
* Estimated wait time by position
* Highlight when your song is next

---

## 🎛️ Admin Flow

### 1. Admin Dashboard

* Central control room for the night

### 2. Monitor Requests

* View all incoming requests
* See table number and timestamps
* Filter by status or table

### 3. Approve or Reject

* ✓ Approve to add to playlist
* ✗ Reject with optional reason

### 4. Manage the Playlist

* Drag and drop to reorder songs
* Control performance flow
* Add internal notes for the band

### 5. Mark Songs as Played

* One click when a song is performed
* Moves to history automatically
* Next song becomes active

### 6. Song Management

* Add or edit songs manually
* Import directly from Genius
* Manage genres and categories

---

## 🎵 Final Note

This project is built with care, clarity, and a love for live music.

It’s not just software—it’s rhythm, timing, and trust between performers and their audience.

If you need help setting up or extending the system, feel free to open an issue on GitHub.

**Enjoy the music. Let the room decide the next song.** 🎸🎹🥁
