# ASU Course Explorer 🔱
A high-performance API for accessing the ASU Course Catalog, built with Go.

## About The Project

Arizona State University doesn't provide an openly available API to access their course catalog. Currently, students must navigate through the university's [website](https://catalog.apps.asu.edu/catalog/courses/courselist) to find course information, which is often buried across multiple webpages.

**ASU Course Explorer** solves this problem by providing a clean, fast API interface to access ASU course data programmatically.

This project started as a fun side project, and a heavily modified, even faster production version now powers the backend for [Courseer](https://courseer.co/)!

## Built With

- [Golang](https://go.dev/) - Programming Language
- [Rod](https://github.com/go-rod/rod) - Driver for DevTools Protocol
- [Gin](https://github.com/gin-gonic/gin) - Web Framework
- [Redis](https://github.com/redis/go-redis) - In-Memory Data Structure Store (for caching)

## Getting Started

Clone the repository and make sure Go is installed before running the following commands:

```shell
cd asu-course-explorer/

# Install dependencies
go mod download

# Set up environment variables (create .env file with REDIS_ENDPOINT)
cp .env.example .env  # Edit with your Redis endpoint

# Run the server
go run .
```

The server will start on port 8080 by default.

## API Endpoints

### Classes

- **`GET /api/classes`** - Fetches a list of all classes
- **`GET /api/classes/{classId}`** - Fetches information about a specific class (5-digit class ID)
- **`GET /api/classes/course/{courseId}`** - Fetches all classes for a specific course (format: 3-letter subject + 3-digit number, e.g., "CSE110")

### Response Format

Each class object includes:

- Course code and title
- Class number and instructor
- Meeting days, times, and location
- Available seats and units
- Syllabus link (if available)
- General Studies designations

## Features

- **Redis Caching**: 30-minute cache TTL for improved performance
- **Web Scraping**: Real-time data from ASU's course catalog
- **RESTful API**: Clean, predictable endpoints
- **Error Handling**: Robust error responses and validation

## Roadmap

See the [open issues](https://github.com/spabolu/asu-course-explorer/issues) for planned features:

- [x] Add specific class endpoint
- [x] Add Redis caching
- [ ] Add semesters endpoint
- [ ] Add professors endpoint
- [ ] Add course descriptions endpoint
- [ ] Add Docker support
- [ ] Add rate limiting

## Contributing

Contributions are welcome! If you have suggestions for improvements, please fork the repo and create a pull request, or open an issue with the "enhancement" tag.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the GPL-3.0 License. See `LICENSE` for more information.

## Support

If you like what you see and want to use a powerful iteration of this program to track your classes at ASU, please support and use [Courseer](https://courseer.co/).
