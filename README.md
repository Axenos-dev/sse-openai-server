# SSE OpenAI server

This document provides instructions on how to run and interact with this server.

## Table of Contents

- [Running the API](#running-the-api)
- [API Endpoints](#api-endpoints)


## Running the API

1. ### Install dependencies:

    ```bash
    make deps
    ```

2. ### Set environment variables:

    ```bash
    export PORT=8080;
    export OPEN_AI_API_KEY=your_openai_api_key;
    export POSTGRES_HOST=127.0.0.1;
    export POSTGRES_PORT=5432;
    export POSTGRES_NAME=username;
    export POSTGRES_PASSWORD=password;
    export POSTGRES_DB=some_db;
    ```

3. ### Run the API:

    ```bash
    make build;
    make run;
    ```

    or

    ```bash
    go run main.go
    ```

    - or run it via docker (place OPEN_AI_API_KEY in .env file)
    
    ```bash
    docker-compose up
    ```
    it will run at port :8080


The API should now be running on `http://localhost:PORT`.

## API Endpoints

### /v1/chat/:topic

Send a message, then it it starts stream on -[/v1/sse/:topic]

- **Endpoint:** `/v1/chat/:topic`
- **Method:** `POST`
- **Path Parameters:**

    ```
    topic - chat topic that the client is listening to
    ```

- **Request Body:**

    ```json
    {
      "message": "Hello!"
    }
    ```

- **Response:**

    ```json
    {
      "code": 200,
      "message": "message with content 'Hello!' sent"
    }
    ```

### /v1/sse/:topic

Streams Server Send Events for specific :topic

- **Endpoint:** `/v1/sse/:topic`
- **Method:** `GET`
- **Path Parameters:**

    ```
    topic - chat topic that the client is listening to
    ```

- **Response:**

    ```json
    {
        "data": {
            "content": "Hi there! How can I assist"
        },
        "event":"message_completion"
    }
    ```
