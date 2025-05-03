# Match-Me Web Application

*A recommendation platform to connect people based on shared interests*

## Table of Contents

- Features
    
- Technologies
    
- Installation
    
- Configuration
    
- Usage
    
- API Endpoints
    
- Screenshots
    
- Contributing
    
- License
    

## Features

✅ **User Authentication**

- Secure JWT-based registration/login
    
- Password hashing with bcrypt
    
- Session management
    

💼 **Profile System**

- Complete user profiles with bios
    
- Profile pictures upload
    
- 5+ matching criteria (interests, location, etc.)
    

🔍 **Smart Recommendations**

- Algorithm-based matching
    
- Location-based filtering
    
- Swipe-style interface
    

🤝 **Connection Management**

- Send/accept connection requests
    
- View current connections
    
- Disconnect from users
    

💬 **Real-time Chat**

- Instant messaging between connected users
    
- Message history
    
- Unread message indicators
    

## Technologies

**Backend**:

- Go (Gin framework)
    
- PostgreSQL
    
- JWT authentication
    
- WebSocket for real-time chat
    

**Frontend**:

- React.js
    
- React Router
    
- Context API for state management
    
- Axios for API calls
    
- CSS for styling
    

**DevOps**:

- Docker
    
- Docker Compose
    
- PostgreSQL
    

## Installation

### Prerequisites

- Docker and Docker Compose installed
    
- Node.js (v14+) for frontend development
    

### Setup

1.  Clone the repository:

`git clone https://gitea.koodsisu.fi/ihorshaposhnik/match-me`  
`cd match-me`

2.  Set up environment variables:

`cp backend/.env.example backend/.env`  
`cp frontend/.env.example frontend/.env`

Edit the files with your configuration.

3.  Build and start the containers:

`docker-compose up --build`

4.  The application will be available at:

- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

## Configuration

### Backend Environment Variables

`PORT=8080`  
`DB_HOST=db`  
`DB_PORT=5432`  
`DB_USER=matchme`  
`DB_PASSWORD=securepassword`  
`DB_NAME=matchme`  
`JWT_SECRET=your_jwt_secret_here`

### Frontend Environment Variables

`REACT_APP_API_URL=http://localhost:8080`

## Usage

1.  **Registration**: Create a new account with your email
    
2.  **Profile Setup**: Complete your profile with interests and preferences
    
3.  **Discover**: Browse recommended matches
    
4.  **Connect**: Send connection requests to interesting profiles
    
5.  **Chat**: Message with your connections in real-time
    

## API Endpoints

| Endpoint | Method | Description |
| --- | --- | --- |
| `/register` | POST | Register new user |
| `/login` | POST | User login |
| `/profile` | GET/PUT | Get/update user profile |
| `/bio` | GET/PUT | Get/update biographical data |
| `/recommendations` | GET | Get match recommendations |
| `/connect` | POST | Send connection request |
| `/connections` | GET | List user connections |
| `/chat` | GET/POST | Get/send chat messages |
| `/ws` | GET | WebSocket endpoint |

## License

Distributed under the MIT License. See `LICENSE` for more information.