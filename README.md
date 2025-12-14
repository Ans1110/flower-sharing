# Flower Sharing

A full-stack social media application for sharing flowers and connecting with others.

## 🌐 Live Deployment

- **Frontend**: [flower-sharing.vercel.app](https://flower-sharing.vercel.app)
- **Backend API**: [flower-sharing.zeabur.app](https://flower-sharing.zeabur.app)

## 🛠️ Tech Stack

### Frontend
- **Framework**: Next.js 16 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS 4
- **State Management**: Zustand
- **Data Fetching**: TanStack Query (React Query)
- **Forms**: React Hook Form + Zod
- **UI Components**: shadcn/ui, Radix UI

### Backend
- **Language**: Go 1.24.5
- **Framework**: Gin
- **Database**: MySQL (via GORM)
- **Authentication**: JWT + OAuth2
- **File Storage**: Cloudinary
- **Documentation**: Swagger
- **Logging**: Zap

## 📁 Project Structure

```
flower-sharing/
├── frontend/          # Next.js frontend application
│   ├── src/
│   │   ├── app/      # Next.js app router pages
│   │   ├── components/  # React components
│   │   ├── hooks/    # Custom React hooks
│   │   ├── lib/      # Utility libraries
│   │   ├── service/  # API service layer
│   │   ├── store/    # Zustand stores
│   │   └── types/    # TypeScript type definitions
│   └── package.json
│
└── backend/          # Go backend API
    ├── controllers/  # HTTP handlers
    ├── services/    # Business logic
    ├── repositories/ # Data access layer
    ├── models/      # Database models
    ├── middlewares/ # HTTP middlewares
    ├── routes/      # API routes
    ├── dto/         # Data transfer objects
    └── go.mod
```

## 🚀 Getting Started

### Prerequisites

- **Node.js** 20+ (for frontend)
- **Go** 1.24.5+ (for backend)
- **MySQL** database
- **Cloudinary** account (for image uploads)

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
bun install
# or
npm install
```

3. Create a `.env.local` file with your environment variables:
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
# Add other environment variables as needed
```

4. Run the development server:
```bash
bun dev
# or
npm run dev
```

5. Open [http://localhost:3000](http://localhost:3000) in your browser.

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file with your environment variables:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=flower_sharing
JWT_SECRET=your_jwt_secret
CLOUDINARY_URL=your_cloudinary_url
# Add other environment variables as needed
```

4. Run the server:
```bash
go run main.go
```

5. API will be available at [http://localhost:8080](http://localhost:8080)
6. Swagger documentation at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## 📝 Features

- User authentication (JWT + OAuth2)
- User profiles and following system
- Post creation and management
- Image uploads via Cloudinary
- Admin panel for user and post management
- Responsive design with dark mode support

## 🔧 Development

### Frontend Scripts

- `bun dev` - Start development server
- `bun build` - Build for production
- `bun start` - Start production server
- `bun lint` - Run ESLint

### Backend

- `go run main.go` - Run development server
- `go build` - Build binary
- `swag init` - Generate Swagger documentation
