# Backend Assignment | Fam

This is a backend assignment for FamPay's Intern position. The goal of this project is to create an API to fetch the latest videos from YouTube for a given search query, store them in a MongoDB database, and provide a GET API for retrieving the stored videos in a paginated, sorted manner. The frontend is implemented in ReactJS.

## Project Structure

The backend is divided into three layers:
- **API Layer**: Handles incoming HTTP requests and routes them to the appropriate handlers.
- **Daos Layer**: Handles interactions with the MongoDB database.
- **Service Layer**: Implements business logic, including calling the YouTube API, storing data, and pagination.

The frontend is implemented using ReactJS:
- **API Integration**: Makes API calls to the backend for fetching video data.

## Implemented Basic Requirements

- Continuously fetches latest videos from YouTube for a predefined search query.(Using go cron)
- Stores video data (title, description, publishing datetime, thumbnails URLs, etc.) in MongoDB with proper indexes.
- Provides a GET API to retrieve stored videos in a paginated response sorted by published datetime.(See service layer code for reference)
- Frontend implemented using ReactJS.   

## Implemented Bonus Points

- Dashboard to view stored videos with filters and sorting options.

## Project Setup

### Backend (GoLang)

#### Prerequisites
- Go installed
- MongoDB installed and running or Docker installed
- NPM installed


#### Steps to run locally (Use ENV file I have attached in mail)
   ```bash
   $ git clone https://github.com/yourusername/backend.git
   $ chmod +x install_dependencies.sh
   $ ./install_dependencies.sh

   $ go build
   $ ./ytvideofetcher

   $ docker run -d -p 27017:27017 --name test-mongo mongo:latest

   $ cd frontend
   $ npm install --force
   $ npm start

Now the code would be running successfully
If any issues mail at: vishnuofficial1912@gmail.com
//This project is solely made by Vishnu


