# GoDam Backend

A robust Go-based backend service for the GoDam inventory management system, providing RESTful APIs, Firebase integration, and image processing capabilities.

## Features

- 🔐 Firebase Authentication
- 📦 Product Management
- 🖼️ Image Processing
- 📊 Inventory Tracking
- 🔍 Barcode Scanning
- 🔄 Real-time Updates
- 🔒 Role-based Access Control

## Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin
- **Database**: Firebase Firestore
- **Storage**: Firebase Storage
- **Authentication**: Firebase Auth
- **Testing**: Go Testing Framework

## Project Structure

```
godam-backend/
├── cmd/              # Application entry points
├── config/           # Configuration files
├── controllers/      # Request handlers
├── handlers/         # HTTP route handlers
├── interfaces/       # Interface definitions
├── libraries/        # Shared utilities
├── middleware/       # HTTP middleware
├── models/           # Data models
├── services/         # Business logic
├── types/            # Type definitions
└── tests/            # Test files
```

## Prerequisites

- Go 1.21 or higher
- Firebase project with Firestore and Storage enabled
- Firebase service account credentials
- Python 3.8+ (for image recognition service)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/godam-backend.git
cd godam-backend
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your Firebase credentials
```

4. Run the server:
```bash
go run main.go
```

## API Documentation

### Authentication
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration
- `POST /api/auth/refresh` - Token refresh

### Products
- `GET /api/products` - List products
- `POST /api/products` - Create product
- `PUT /api/products/:id` - Update product
- `DELETE /api/products/:id` - Delete product
- `GET /api/products/barcode/:barcode` - Find product by barcode

### Image Processing
- `POST /api/images/upload` - Upload product image
- `POST /api/images/process` - Process image for recognition
- `POST /api/products/scan` - Scan product with image

## Testing

Run tests:
```bash
go test ./...
```

## Security

- JWT-based authentication
- Firebase security rules
- CORS configuration
- Input validation
- Rate limiting

## Deployment

1. Build the application:
```bash
go build -o godam-backend
```

2. Deploy to your server:
```bash
./godam-backend
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
