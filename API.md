# VIPTV-Server API Documentation

This document provides detailed information about the API endpoints available in the VIPTV-Server application. The API is designed to handle user authentication, manage playlists, and stream video content.

## Base URL

```
http://localhost:8080/api
```

## Authentication

### Register

- **Endpoint**: `/auth/register`
- **Method**: `POST`
- **Description**: Registers a new user.
- **Request Body**:

  ```json
  {
    "username": "string",
    "password": "string",
    "email": "string"
  }
  ```

- **Response**:

  - `201 Created`:

    ```json
    {
      "message": "User registered successfully"
    }
    ```

  - `400 Bad Request`:

    ```json
    {
      "error": "Bad Request"
    }
    ```

### Login

- **Endpoint**: `/auth/login`
- **Method**: `POST`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**:

  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- **Response**:

  - `200 OK`:

    ```json
    {
      "token": "jwt_token"
    }
    ```

  - `401 Unauthorized`:

    ```json
    {
      "error": "Invalid username or password"
    }
    ```

## Playlist Management

### Get All Playlists

- **Endpoint**: `/playlist`
- **Method**: `GET`
- **Description**: Retrieves a list of all playlists.
- **Response**:

  - `200 OK`:

    ```json
    [
      {
        "id": "integer",
        "name": "string",
        "description": "string"
      }
    ]
    ```

### Get Playlist by ID

- **Endpoint**: `/playlist/:id`
- **Method**: `GET`
- **Description**: Retrieves a specific playlist by its ID.
- **Response**:

  - `200 OK`:

    ```json
    {
      "id": "integer",
      "name": "string",
      "description": "string",
      "videos": [
        {
          "id": "integer",
          "title": "string",
          "url": "string"
        }
      ]
    }
    ```

  - `404 Not Found`:

    ```json
    {
      "error": "Playlist not found"
    }
    ```

### Create Playlist

- **Endpoint**: `/playlist`
- **Method**: `POST`
- **Description**: Creates a new playlist.
- **Request Body**:

  ```json
  {
    "name": "string",
    "description": "string"
  }
  ```

- **Response**:

  - `201 Created`:

    ```json
    {
      "message": "Playlist created successfully"
    }
    ```

  - `400 Bad Request`:

    ```json
    {
      "error": "Bad Request"
    }
    ```

### Update Playlist

- **Endpoint**: `/playlist/:id`
- **Method**: `PUT`
- **Description**: Updates an existing playlist.
- **Request Body**:

  ```json
  {
    "name": "string",
    "description": "string"
  }
  ```

- **Response**:

  - `200 OK`:

    ```json
    {
      "message": "Playlist updated successfully"
    }
    ```

  - `404 Not Found`:

    ```json
    {
      "error": "Playlist not found"
    }
    ```

### Delete Playlist

- **Endpoint**: `/playlist/:id`
- **Method**: `DELETE`
- **Description**: Deletes a playlist by its ID.
- **Response**:

  - `200 OK`:

    ```json
    {
      "message": "Playlist deleted successfully"
    }
    ```

  - `404 Not Found`:

    ```json
    {
      "error": "Playlist not found"
    }
    ```

## Video Management

### Get Video by ID

- **Endpoint**: `/video/:id`
- **Method**: `GET`
- **Description**: Retrieves a specific video by its ID.
- **Response**:

  - `200 OK`:

    ```json
    {
      "id": "integer",
      "title": "string",
      "url": "string",
      "playlist_id": "integer"
    }
    ```

  - `404 Not Found`:

    ```json
    {
      "error": "Video not found"
    }
    ```

### Create Video

- **Endpoint**: `/video`
- **Method**: `POST`
- **Description**: Adds a new video to a playlist.
- **Request Body**:

  ```json
  {
    "title": "string",
    "url": "string",
    "playlist_id": "integer"
  }
  ```

- **Response**:

  - `201 Created`:

    ```json
    {
      "message": "Video added successfully"
    }
    ```

  - `400 Bad Request`:

    ```json
    {
      "error": "Validation error message"
    }
    ```

### Update Video

- **Endpoint**: `/video/:id`
- **Method**: `PUT`
- **Description**: Updates the details of an existing video.
- **Request Body**:

  ```json
  {
    "title": "string",
    "url": "string"
  }
  ```

- **Response**:

  - `200 OK`:

    ```json
    {
      "message": "Video updated successfully"
    }
    ```

  - `404 Not Found`:

    ```json
    {
      "error": "Video not found"
    }
    ```

### Delete Video

- **Endpoint**: `/video/:id`
- **Method**: `DELETE`
- **Description**: Deletes a video by its ID.
- **Response**:

  - `200 OK`:

    ```json
    {
      "message": "Video deleted successfully"
    }
    ```

  - `404 Not Found`:

    ```json
    {
      "error": "Video not found"
    }
    ```

## Error Handling

- **400 Bad Request**: Returned when the request is invalid or missing required parameters.
- **401 Unauthorized**: Returned when the user is not authenticated.
- **404 Not Found**: Returned when the requested resource is not found.
- **500 Internal Server Error**: Returned when an unexpected error occurs on the server.

## Authentication Header

For endpoints requiring authentication, include the following header in your requests:

```
Authorization: Bearer <jwt_token>
```

## Example Requests

### cURL

#### Register

```sh
curl -X POST http://localhost:8080/api/auth/register \
-H "Content-Type: application/json" \
-d '{"username": "user1", "password": "pass123", "email": "user1@example.com"}'
```

#### Login

```sh
curl -X POST http://localhost:8080/api/auth/login \
-H "Content-Type: application/json" \
-d '{"username": "user1", "password": "pass123"}'
```

#### Get All Playlists

```sh
curl -X GET http://localhost:8080/api/playlist
```

## Conclusion

This documentation provides a comprehensive guide to the VIPTV-Server API endpoints. For any questions or problems, consult the Wiki or Issues section or contact the project maintainers.

---
