# URL Shortener in Go using Gin

Welcome to the URL Shortener service built with Go and the Gin web framework! This project allows users to shorten URLs and retrieve the original URLs using a short identifier.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [API Endpoints](#api-endpoints)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/URL-Shortener_Go.git
    cd URL-Shortener_Go
    ```

2. **Install dependencies**:
    Ensure you have Go installed on your machine. Then, run:
    ```sh
    go mod tidy
    ```

3. **Set up MongoDB and Redis**:
    - Install and start MongoDB.
    - Install and start Redis.

4. **Environment variables**:
    Create a `.env` file and add your MongoDB and Redis configurations.

5. **Run the application**:
    ```sh
    go run main.go
    ```

## Usage

- **Signup**: Create a new user account.
- **Login**: Authenticate an existing user.
- **Shorten URL**: Generate a shortened URL.
- **Redirect**: Use the shortened URL to redirect to the original URL.

## Features

- User authentication (signup and login)
- URL shortening
- URL redirection
- Visit tracking and caching with Redis

## API Endpoints

### Auth Routes

- **POST /auth/signup**
    - Request body:
        ```json
        {
            "name": "your name",
            "email": "your email",
            "password": "your password"
        }
        ```
    - Response: Status and message indicating the result of the signup.

- **POST /auth/login**
    - Request body:
        ```json
        {
            "email": "your email",
            "password": "your password"
        }
        ```
    - Response: Status, session information, and session ID.

### URL Routes

- **POST /url/shorten?sessionID=YOUR_SESSION_ID**
    - Request body:
        ```json
        {
            "redirecturl": "URL to be shortened",
            "createdby": "user's email"
        }
        ```
    - Response: The shortened URL information.

- **POST /url/customshorten?sessionID=YOUR_SESSION_ID**
    - Request body:
        ```json
        {
            "shortid": "custom short ID",
            "redirecturl": "URL to be shortened",
            "createdby": "user's email"
        }
        ```
    - Response: The shortened URL information.

- **GET /:shortid**
    - Redirects to the original URL corresponding to the given short ID.

## Technologies Used

- **Gin**: A web framework for Go. Used for building the HTTP server and defining routes.
- **MongoDB**: A NoSQL database used for storing user information, URLs, and session data.
- **Redis**: An in-memory data store used for caching and session management.

## Contributing

1. **Fork the repository**.
2. **Create a new branch** for your feature or bug fix.
3. **Commit your changes**.
4. **Push to the branch**.
5. **Create a Pull Request**.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or feedback, please contact [me](mailto:harshhvstech1975@gmail.com).
