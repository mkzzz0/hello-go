# hello-go
# Go Web App with Nginx in Docker

## Description
This repository contains a Go web application with Nginx page that displays "Hello World!".

## Running the Project

Follow these steps to run the application locally:

1. Clone or download this repository:
   git clone <REPOSITORY_URL>
   cd hello-go
   
2. Install Docker and Docker Compose if they are not already installed:
   
   apt install docker docker-compose  # For Debian/Ubuntu
   yum install docker docker-compose  # For CentOS/RHEL
   
3. Build the images and start the containers:
   
   docker-compose up -d
   
4. Open `http://localhost` in your browser.
