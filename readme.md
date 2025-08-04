# URL Shortener Service

A simple and efficient URL shortening service built with Go. This service allows you to create short codes for long URLs and redirect users to the original URLs using those codes.

## Features

- **URL Shortening**: Convert long URLs into short, manageable codes
- **URL Redirection**: Redirect short codes back to original URLs
- **Persistent Storage**: JSON-based file storage for URL mappings
- **RESTful API**: Clean HTTP endpoints for easy integration
- **Environment Configuration**: Configurable port and base URL settings

## Project Structure

```
urlShortener/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── shortenUrl.go    # URL shortening endpoint handler
│   │   └── redirectUrl.go   # URL redirection endpoint handler
│   ├── storage/
│   │   └── storage.go       # JSON file storage implementation
│   └── utils/
│       └── hash.go          # Short code generation utilities
├── assets/
│   └── urlStorage.json      # Persistent storage file
├── .env                     # Environment configuration
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
└── readme.md                # Project documentation
```

## API Endpoints

### Shorten URL
**POST** `/shorten`

Create a short code for a long URL.

**Request Body:**
```json
{
  "url": "https://example.com/very/long/url"
}
```

**Response:**
```json
{
  "code": "CmAJMizX"
}
```

**Status Codes:**
- `201` - Successfully created short code
- `400` - Invalid request body or missing URL
- `405` - Method not allowed

### Redirect URL
**GET** `/{code}`

Redirect to the original URL using the short code.

**Response:**
```json
{
  "url": "https://example.com/very/long/url"
}
```

**Status Codes:**
- `302` - Found (with redirect URL in response)
- `404` - Short code not found
- `405` - Method not allowed

## Installation & Setup

### Prerequisites
- Go 1.24.4 or higher
- Git

### Clone the Repository
```bash
git clone https://github.com/RbMo7/urlLinkShortner-Service.git
cd urlLinkShortner-Service
```

### Install Dependencies
```bash
go mod download
```

### Environment Configuration
Create a `.env` file in the root directory (or copy from `.env.example`):

```env
PORT=8081
BASE_URL=http://localhost:8081
DB_HOST=localhost
```

### Run the Application
```bash
# From the project root
cd cmd
go run main.go
```

Or build and run the executable:
```bash
cd cmd
go build -o ../urlshortener main.go
cd ..
./urlshortener
```

## Usage Examples

### Shorten a URL
```bash
curl -X POST http://localhost:8081/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.example.com/very/long/url/path"}'
```

Response:
```json
{
  "code": "CmAJMizX"
}
```

### Access Shortened URL
```bash
curl http://localhost:8081/CmAJMizX
```

Response:
```json
{
  "url": "https://www.example.com/very/long/url/path"
}
```

## How It Works

1. **URL Shortening**: The service generates a unique 8-character code using SHA-1 hashing and Base64 encoding of the input URL
2. **Storage**: URL mappings are stored in a JSON file (`assets/urlStorage.json`) with thread-safe read/write operations
3. **Redirection**: When a short code is accessed, the service looks up the original URL and returns it in the response

## Dependencies

- **[godotenv](https://github.com/joho/godotenv)**: Environment variable loading from `.env` files

## Development

### Code Organization
- `cmd/`: Application entry point and main function
- `internal/handlers/`: HTTP request handlers for different endpoints
- `internal/storage/`: Data persistence layer with JSON file storage
- `internal/utils/`: Utility functions for short code generation
- `assets/`: Static files and data storage

### Key Components
- **Storage Layer**: Thread-safe in-memory map with JSON persistence
- **Hash Generation**: SHA-1 based short code generation for URL uniqueness
- **HTTP Handlers**: Clean separation of concerns for different API endpoints

## Configuration

The application uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8081` |
| `BASE_URL` | Base URL for the service | `http://localhost:8081` |
| `DB_HOST` | Database host (reserved for future use) | `localhost` |

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

**RbMo7** - [GitHub Profile](https://github.com/RbMo7)

---

**Note**: This is a learning project built while exploring Go web development. The current implementation uses file-based storage and is suitable for development and small-scale deployments.