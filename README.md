# Key-Value Cache Service

## Overview
This project implements an in-memory key–value cache service using Go. It provides two RESTful endpoints:
- **POST /put**: Inserts or updates a key-value pair.
- **GET /get**: Retrieves the value associated with a given key.

The service is containerized with Docker and is designed to run on an AWS t3.small instance (2 cores, 2GB RAM). Load testing is performed using Locust.

## Folder Structure
keyvalue-cache/ ├── main.go ├── go.mod ├── Dockerfile ├── locustfile.py └── README.md


## Setup & Build Instructions

### Prerequisites
- **Go** (1.20 or later)
- **Docker**
- **Git** (if cloning from a repository)
- **Python 3 & pip** (for Locust)
- **Locust** (install using `pip install locust`)

### Environment Setup on Linux
For Ubuntu/Debian:
```bash
# Update package lists
sudo apt update

# Install Git and Python3 if not installed
sudo apt install -y git python3 python3-pip

# Install Docker
sudo apt install -y docker.io

# Verify Docker installation
docker --version

# Install Go (using snap for simplicity)
sudo snap install go --classic

# Install Locust using pip
pip3 install locust

