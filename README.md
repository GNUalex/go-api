# Sample Go REST API

![Go Version](https://img.shields.io/badge/Go-1.23-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview

This repository contains a sample REST API written in Go.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.23 or higher
- [Docker](https://www.docker.com/get-started) (optional, for containerization)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/GNUalex/go-api.git
    cd go-api
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Build the project:
    ```bash
    go build -o app
    ```

### Running the API

To run the API locally, use the following command:

```bash
go run .
```

The API will be accessible at `http://localhost:8080`.

### Docker

To run the API in a Docker container:

Run the Docker container:
```bash
docker run --rm -ti -v $PWD:/app -w /app -p 8080:8080 golang:alpine go run .
```

The API will be accessible at `http://localhost:8080`.

## API Endpoints

Here are some example endpoints provided by the API:

- **GET** `/posts`: Retrieve all posts.
- **GET** `/post/{slug}`: Retrieve a single post by slug.

### Request and Response Examples

**GET** `/posts`

_Response:_
```json
[
    {
        "id": 1,
        "title": "Sample blog post 1",
        "slug": "sample-blog-post-1",
        "publishedAt": "2024-09-01 16:26:03.227426841 +0200 CEST m=+4.718829503",
        "author": "GNUalex",
        "content": "Proin tristique sagittis nisl, sit amet ultricies nulla vehicula nec. Ut imperdiet lacus eu porttitor feugiat. Aenean ipsum risus, rhoncus ac tempor at, rhoncus vitae dui. Aliquam erat volutpat. Morbi non vulputate ipsum."
    }
]
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Go Official Documentation](https://golang.org/doc/) - Excellent resource for learning Go
