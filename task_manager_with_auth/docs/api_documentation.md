# Task Management API Documentation

## Endpoints

### GET /tasks

**Description:** Get a list of all tasks.

**Response:**
- **Status Code:** 200 OK
- **Body:**
    ```json
    [
        {
            "id": 1,
            "title": "Task 1",
            "description": "Description of Task 1",
            "duedate": "2024-08-08T00:00:00Z",
            "status": "pending"
        },
        ...
    ]
    ```

### GET /tasks/:id

**Description:** Get the details of a specific task.

**Parameters:**
- `id` (path): ID of the task.

**Response:**
- **Status Code:** 200 OK
- **Body:**
    ```json
    {
        "id": 1,
        "title": "Task 1",
        "description": "Description of Task 1",
        "duedate": "2024-08-08T00:00:00Z",
        "status": "pending"
    }
    ```
- **Status Code:** 404 Not Found (if task does not exist)

### POST /tasks

**Description:** Create a new task.

**Request Body:**
- **Content-Type:** application/json
- **Body:**
    ```json
    {
        "title": "New Task",
        "description": "Description of New Task",
        "duedate": "2024-08-08T00:00:00Z",
        "status": "pending"
    }
    ```

**Response:**
- **Status Code:** 201 Created
- **Body:**
    ```json
    {
        "id": 1,
        "title": "New Task",
        "description": "Description of New Task",
        "duedate": "2024-08-08T00:00:00Z",
        "status": "pending"
    }
    ```

### PUT /tasks/:id

**Description:** Update a specific task.

**Parameters:**
- `id` (path): ID of the task.

**Request Body:**
- **Content-Type:** application/json
- **Body:**
    ```json
    {
        "title": "Updated Task",
        "description": "Updated Description",
        "duedate": "2024-08-09T00:00:00Z",
        "status": "completed"
    }
    ```

**Response:**
- **Status Code:** 200 OK
- **Body:**
    ```json
    {
        "id": 1,
        "title": "Updated Task",
        "description": "Updated Description",
        "duedate": "2024-08-09T00:00:00Z",
        "status": "completed"
    }
    ```
- **Status Code:** 404 Not Found (if task does not exist)

### DELETE /tasks/:id

**Description:** Delete a specific task.

**Parameters:**
- `id` (path): ID of the task.

**Response:**
- **Status Code:** 204 No Content
- **Status Code:** 404 Not Found (if task does not exist)
