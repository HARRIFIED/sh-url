# sh-url

A URL shortener backend service written in Go, backed by PostgreSQL and designed for easy deployment with Docker Compose.

## Features

- **Shorten URLs** – Quickly generate short links for any valid URL.
- **Redirects** – Seamless redirection from short links to original URLs.
- **RESTful API** – HTTP endpoints for creating and resolving short URLs.
- **Beautiful Landing page** – Simple Frontend Interface for testing.
- **Persistent Storage** – Uses PostgreSQL to reliably store URL mappings.
- **Containerized** – Ready-to-run with Docker and Docker Compose.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) [Go](https://golang.org/) for local development

### Quick Start with Docker Compose

Clone the repository:

```bash
git clone https://github.com/HARRIFIED/sh-url.git
cd sh-url
```

Start the backend service and database with Docker Compose:

```bash
docker-compose up --build
```

The backend will be available at http://localhost:8080

### Configuration
Environment variables can be set in your .env file (running locally):
Check .env.example for required envs, then create a .env and add them

## API Endpoints

### Frontend interface
**GET** `/`
![Sh-url landing page](image.png)

### Shorten a URL

**POST** `/shorten`

```json
// request
{
  "URL": "https://example.com/ioeroroierouerfioero"
}

//response
{
  "code":"efd3350d"
}
```
### Redirect

**GET** /<shortcode>

Redirects to the original long URL on the browser.

---

## Contributing

Contributions are welcome! If you have ideas, suggestions, or bug fixes, feel free to open an issue or submit a pull request. Please follow the existing code style and add tests where appropriate.

## Roadmap / Planned Features

Key features planned for future releases include:

- **Authentication:** Allow users to register and manage their own links in one place.
- **Link Analytics:** Track and display analytics for each short link (e.g., number of redirects, referrers, etc.).
- **Hot Links:** Identify and display the most popular (most redirected) links.
- **Caching Hot Links:** Use Redis to cache hot links for faster lookups and improved performance.
- **Custom Short Links:** Enable premium users to create custom short codes for their URLs.

If you’re interested in helping implement any of these features, please reach out or submit a PR!











