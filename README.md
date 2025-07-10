# 📰 Gator RSS

Gator RSS is a lightweight RSS reader backend written in Go. It lets you manage RSS feeds, fetch and store posts, and serve them efficiently via a Postgres-backed API.

## 🚀 Features

- Add and manage RSS feeds  
- Periodically fetch latest posts  
- Store feeds and posts in a Postgres database  
- Easily extendable for frontend integration  

## 🛠 Requirements

Before running the project, make sure you have the following installed:

- **Go** (version 1.20 or higher recommended)  
- **PostgreSQL** (any modern version, tested with 14+)

## 🛠 Installing the Gator CLI

You can install the `gator` command-line tool globally using the `go install` command:

```bash
go install github.com/yourusername/gator-rss/cmd/gator@latest
```
