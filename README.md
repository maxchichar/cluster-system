# 🏛️ CLUSTER SYSTEM (House Verification System)

A lightweight Discord-based verification system built in Go that controls user onboarding using invite codes, House assignments, and automated role allocation.

This system ensures every member is verified, tracked, and assigned to a structured House before accessing the server.

---

# ⚙️ SYSTEM OVERVIEW

Cluster system is built around 4 core Houses:

- 🔵 Kernel (Leadership & coordination)
- 🟠 Compiler (Builders & creators)
- 🔴 Runtime (Execution & operations)
- 🟣 Algorithm (Strategy & logic)

Each user must enter a valid invite code before gaining access.

---

# 🧠 CORE FLOW

1. User joins Discord server
2. Bot assigns `UNVERIFIED` role
3. User submits invite code via `/verify CODE`
4. Go backend validates code
5. System assigns:
   - House
   - Table (1–25 mapping system)
   - Slot (S1–S4)
6. Discord role is automatically assigned
7. User gains access to server

---

# 🧱 TECH STACK

- Go (backend + bot logic)
- SQLite (database)
- DiscordGo (Discord bot library)

---

# 📁 PROJECT STRUCTURE

```
cluster-system/
├── cmd/bot/               # Discord bot entry
├── internal/
│   ├── config/            # Config loader
│   ├── db/                # Database connection + schema
│   ├── discord/          # Bot handlers + role logic
│   ├── houses/           # House-to-table mapping rules
│   ├── models/           # Data structures
│   └── services/         # Core logic (verify, generate)
├── scripts/              # Code seeding tools
├── data/                 # SQLite database
└── go.mod
```


# 🗄️ DATABASE SCHEMA

### invite_codes
- code (unique)
- house (KER/COM/RNT/ALG)
- table_id
- slot (S1–S4)
- used (0/1)
- used_by

### users
- discord_id
- house
- table_id
- slot
- verified_at

---

# 🔐 INVITE CODE FORMAT

```

HOUSE-TABLE-SEAT-RANDOM

```

Examples:
```

KER-T03-S2-X9K2LD
COM-T11-S4-ZP91QW
RNT-T14-S1-MC82XA

````

---

# 🧠 HOUSE MAPPING

| House | Tables |
|------|--------|
| Kernel | 1–6 |
| Compiler | 7–12, 25 |
| Runtime | 13–18 |
| Algorithm | 19–24 |

---

# 🚀 HOW TO RUN

## 1. Install dependencies
```bash
go mod tidy
````

## 2. Run bot

```bash
go run cmd/bot/main.go
```

---

# 🔑 DISCORD BOT FEATURES

* Auto-assign `UNVERIFIED` role on join
* `/verify <code>` command
* Automatic role assignment after validation
* House-based access control

---

# ⚠️ RULES OF THE SYSTEM

* Each invite code can only be used once
* Every user must be verified before access
* Users cannot bypass verification
* All assignments are logged in database

---

# 🧨 IMPORTANT DESIGN PRINCIPLE

> The Discord bot does not make decisions.
> The backend is the source of truth.

Bot only:

* receives input
* sends it to backend
* applies result

---

# 🔮 FUTURE UPGRADES (NOT IN SYSTEM)

* Admin dashboard (generate/revoke codes)
* Live table occupancy tracker
* Anti-leak detection system
* Analytics per House
* Web-based verification portal

---

# 🏁 GOAL

To build a controlled, structured Discord ecosystem where every user is:

* verified
* assigned
* tracked
* organized into a functional House system

---

# 📌 AUTHOR NOTE

This system is designed for structured communities, coding groups, and competitive learning environments.

```