# ASU Course Explorer 🔱
High-Performance API for ASU Course Catalog written in Go.

### About The Project
Arizona State University doesn't have an openly available API to access their course catalog. As it stands, students must navigate through the university's [website](https://catalog.apps.asu.edu/catalog/courses/courselist) to find this information, which is often buried in several layers of webpages.
`ASU Course Explorer` aims to solve the missing piece to the problem. 

Built this as a fun side project and a heavily-modified production version written in TypeScript is what's running the backend for [Courseer](https://courseer.co/)!
Feel free to implement this project into your personal open-source projects!

### Built With
- [Go](https://go.dev/) - Programming Language
- [Rod](https://github.com/go-rod/rod) - Driver for DevTools Protocol
- [Gin](https://github.com/gin-gonic/gin) - Web Framework
- [Redis](https://github.com/redis/go-redis) - In-Memory Data Structure Store (for caching)

### Get Started
Clone the repository and make sure Go is installed before running the following commands.

```shell
cd asu-course-explorer/

# Install dependencies
go mod download

# Run the server
go run .
```

### API Endpoints (tentative)

`/classes` - Returns all classes for a given subject and term.

`/classes/:subject` - Returns all classes for a given subject and term.

`/classes/:subject/:course` - Returns all classes for a given subject, course, and term.

`/classes/:subject/:course/:term` - Returns all classes for a given subject, course, and term.

`/classes/:subject/:course/:term/:section` - Returns all classes for a given subject, course, term, and section.

`/subjects` - Returns all subjects for a given term.

`/subjects/:term` - Returns all subjects for a given term.

`/terms` - Returns all terms.

`/terms/:term` - Returns all terms.

### Roadmap
See the [open issues](https://github.com/spabolu/asu-course-explorer/issues)

- [x] Add specific class endpoint
- [ ] Add semesters endpoint
- [ ] Add course descriptions endpoint
- [ ] Add Redis caching
- [ ] Add Docker support

### Contributing
If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

- Fork the Project
- Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
- Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
- Push to the Branch (`git push origin feature/AmazingFeature`)
- Open a Pull Request

### License
Distributed under the GPL-3.0. See `LICENSE.txt` for more information.
