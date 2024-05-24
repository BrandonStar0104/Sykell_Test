# Website Information Analyzer

This project is a web application that takes a website URL as an input and provides general information about the contents of the page, such as HTML version, page title, heading tags count by level, amount of internal and external links, inaccessible links, and whether the page contains a login form.

## Features

- Analyze and display HTML version
- Show page title
- Count HTML heading tags by level
- Count internal and external links
- Identify inaccessible links
- Detect presence of login forms

## Technologies Used

- Backend: Golang
- Frontend: React
- Middleware: CORS and Authorization

## Prerequisites

- Go 1.16 or later
- Node.js 14 or later
- npm 6 or later

## Setup Instructions

### Backend Setup

1. **Clone the repository**

    ```sh
    git clone https://github.com/yourusername/website-info-analyzer.git
    cd website-info-analyzer/backend
    ```

2. **Install dependencies**

    ```sh
    go mod download
    ```

3. **Run the backend server**

    ```sh
    go run main.go
    ```

    The backend server will start on `http://localhost:8080`.

### Frontend Setup

1. **Navigate to the frontend directory**

    ```sh
    cd ../frontend
    ```

2. **Install dependencies**

    ```sh
    npm install
    ```

3. **Run the frontend development server**

    ```sh
    npm start
    ```

    The frontend server will start on `http://localhost:3000`.

### Running the Application

1. **Start the backend server**

    Make sure you have the backend server running on `http://localhost:8080` by following the backend setup instructions.

2. **Start the frontend server**

    Make sure you have the frontend server running on `http://localhost:3000` by following the frontend setup instructions.

3. **Open your browser**

    Navigate to `http://localhost:3000` in your web browser to use the application.

## API Endpoints

### Process URL

- **URL**: `/process`
- **Method**: `POST`
- **Headers**: 
  - `Authorization`: Bearer your-secure-token
  - `Content-Type`: application/json
- **Body**: 

    ```json
    {
      "url": "http://example.com"
    }
    ```

### Get Results

- **URL**: `/results`
- **Method**: `GET`
- **Headers**: 
  - `Authorization`: Bearer your-secure-token
  - `Content-Type`: application/json

### Get Result by URL

- **URL**: `/result/{url}`
- **Method**: `GET`
- **Headers**: 
  - `Authorization`: Bearer your-secure-token
  - `Content-Type`: application/json

### Stop Processing

- **URL**: `/stop`
- **Method**: `POST`
- **Headers**: 
  - `Authorization`: Bearer your-secure-token
  - `Content-Type`: application/json
- **Body**: 

    ```json
    {
      "url": "http://example.com"
    }
    ```

## Example Requests

### Process URL

```sh
curl -X POST http://localhost:8080/process \
    -H "Authorization: Bearer your-secure-token" \
    -H "Content-Type: application/json" \
    -d '{"url": "http://example.com"}'

curl -X GET http://localhost:8080/results \
    -H "Authorization: Bearer your-secure-token" \
    -H "Content-Type: application/json"

curl -X GET http://localhost:8080/result/http://example.com \
    -H "Authorization: Bearer your-secure-token" \
    -H "Content-Type: application/json"

curl -X POST http://localhost:8080/stop \
    -H "Authorization: Bearer your-secure-token" \
    -H "Content-Type: application/json" \
    -d '{"url": "http://example.com"}'
