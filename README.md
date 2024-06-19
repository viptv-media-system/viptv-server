# VIPTV-Server

VIPTV-Server is a backend server application built with Go and Gin framework, designed to serve IPTV shows using an M3U playlist. It handles video streaming, authentication, and metadata storage for a self-hosted streaming service.

## Features

- **Video Streaming**: Stream IPTV shows using M3U playlist URLs.
- **Authentication**: Endpoints with user authentication.
- **Database**:  Use PostgreSQL to store metadata and user information.
- **Scalability**: Built with Go's concurrency for handling multiple video streams efficiently.

## Technologies Used

- **Backend Language**: Go (Golang)
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Video Streaming**: FFmpeg

## Folder Structure

```
viptv-server/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── auth.go
│   │   │   ├── playlist.go
│   │   │   └── video.go
│   │   ├── middleware/
│   │   │   └── auth.go
│   │   └── routes.go
│   │
│   ├── database/
│   │   └── postgres.go
│   │
│   ├── models/
│   │   ├── user.go
│   │   └── video.go
│   │
│   ├── services/
│   │   ├── ffmpeg.go
│   │   ├── playlist.go
│   │   └── video.go
│   │
│   └── util/
│       └── errors.go
│
├── config/
│   └── config.yaml
│
├── pkg/
│   └── (third-party packages)
│
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
```

## Getting Started

### Prerequisites

- Go installed on your machine. [Install Go](https://golang.org/doc/install)
- PostgreSQL database set up and running. [Install PostgreSQL](https://www.postgresql.org/download/)
- FFmpeg installed for video processing. [Install FFmpeg](https://ffmpeg.org/download.html)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/viptv-media-system/viptv-server.git
   cd viptv-server
   ```

2. Set up the database:

   - Create a PostgreSQL database and note down the credentials.
   - Execute the initial schema migration:

     <!-- ```bash
     psql -U your_username -d your_database -a -f migrations/001_initial_schema.up.sql
     ``` -->

3. Set up environment variables:

   Create a copy `.env-example` to `.env` file in the root directory and add the following:

   ```
   DB_USER=your_postgres_username
   DB_PASSWORD=your_postgres_password
   DB_NAME=your_database_name
   ```

4. Build and run the server:

   ```bash
   go build ./cmd/server
   ./viptv-server
   ```

   The server should now be running at `http://localhost:8080`.

## Usage

- **Endpoints**:
  - `/api/auth`: Authentication endpoints.
  - `/api/playlist`: Playlist management endpoints.
  - `/api/video`: Video streaming endpoints.

- **Documentation**:
  - Detailed API documentation can be found in [API.md](API.md).

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE Version 2 - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gin](https://gin-gonic.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [FFmpeg](https://ffmpeg.org/)

---

### Notes:

- Replace placeholders like `your-username`, `your-email`, `your_postgres_username`, `your_postgres_password`, and `your_database_name` with your actual details.
