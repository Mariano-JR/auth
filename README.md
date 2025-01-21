# Authentication System with OAuth2

This project is a back-end implementation of a secure and scalable authentication system built with **Golang**. The system is designed to leverage OAuth2 authentication with popular social login providers like Google and GitHub, minimizing the need for storing sensitive user data and focusing on core security and logic.

## Features

- **OAuth2 Integration**: Support for login with social providers (e.g., Google, GitHub).
- **Secure Token Validation**: Implements best practices for handling and validating OAuth tokens.
- **Modular Architecture**: Organized using a clean and scalable project structure.
- **In-Memory User Storage**: Basic user storage to quickly test the system (can be extended to databases like PostgreSQL or MongoDB).
- **Static Front-End Integration**: A simple HTML form to test authentication flows.

## Project Structure

```plaintext
.
├── cmd/                # Entrypoints of the application
│   └── server/         # Main application logic
│       └── main.go     # Server initialization
├── internal/           # Core logic and services
│   ├── auth/           # Authentication module
│   │   ├── handler.go  # Handlers for authentication routes
│   │   ├── service.go  # Authentication business logic
│   │   └── model.go    # Data models for authentication
│   ├── user/           # User management module
│   │   ├── repository.go # User data handling logic
│   │   ├── service.go    # User-related business logic
│   │   └── model.go      # User data model
│   └── utils/          # Utility functions and helpers
│       └── helpers.go  # Logging and common utilities
├── configs/            # Configuration files (e.g., YAML, JSON)
├── static/             # Static HTML files for testing
│   └── index.html      # Simple front-end for authentication
├── migrations/         # Database migration files (optional)
├── scripts/            # Automation and helper scripts
├── docs/               # Documentation and API specs
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksums
└── README.md           # Project documentation
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/auth.git
   cd auth
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   Create a `.env` file in the root directory and add the required variables:
   ```env
   PORT=3000
   GOOGLE_CLIENT_ID=your-google-client-id
   GOOGLE_CLIENT_SECRET=your-google-client-secret
   GITHUB_CLIENT_ID=your-github-client-id
   GITHUB_CLIENT_SECRET=your-github-client-secret
   ```

4. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

5. Access the application:
   Open `http://localhost:3000` in your browser.

## Usage

- Navigate to the home page to access the login buttons for Google and GitHub.
- The `/auth/:provider` route redirects users to the respective OAuth provider.
- The `/auth/callback` route handles the provider's response and validates the token.

## Future Improvements

- Add support for more OAuth providers (e.g., Facebook, LinkedIn).
- Integrate a robust database (e.g., PostgreSQL) for persistent user storage.
- Implement token refresh mechanisms.
- Add unit tests and API documentation (Swagger).

## Contributing

Feel free to fork this repository, open issues, or submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.