# Project Structure

The project structure is defined as the following:
```
├─api              - api exposed to clients
├─internal         - internal packages
│  ├─entity        - definition of tables
│  ├─model         - models for input and output
│  └─service       - business logic
└─utest            - unit tests
```
# How to run?

To run this application in container, you can execute the following commands:

```
$ docker build -t gogolook .
$ docker run -d -p 8080:8080  gogolook
```

# How to do unit test?

To perform unit test, you can execute the following commands:

```
$ cd utest
$ go test -run TestCreateTask
```

Replace 'TestCreateTask' to whatever you want to test.

# Coding Exercise

Implement a Restful task list API as well as run this application in container.

- Spec
  - Fields of task:
      - name
          - Type: String
      - status
          - Type: Bool
          - Value
              - 0=Incomplete
              - 1=Complete
  - Reponse headers
      - Content-Type=application/json
  - Unit Test
  - Manage codebase on Github

- Runtime Environment Requirement
    - If you choose Python
        - Python 3.7+
        - Flask 2.0.x
    - If you choose Golang
        - Go 1.17.8+
        - Gin 1.7.7
    - Docker

- About Database
  - DB is a complex component that you can use in-memory mechanism to handle data storage problem.


### 1.  GET /tasks (list tasks)
```
{
    "result": [
        {"id": 1, "name": "name", "status": 0}
    ]
}
```

### 2.  POST /task  (create task)
```
request
{
  "name": "買晚餐"
}

response status code 201
{
    "result": {"name": "買晚餐", "status": 0, "id": 1}
}
```

### 3. PUT /task/<id> (update task)
```
request
{
  "name": "買早餐",
  "status": 1
  "id": 1
}

response status code 200
{
  "result":{
    "name": "買早餐",
    "status": 1,
    "id": 1
  }
}
```

### 4. DELETE /task/<id> (delete task)
response status code 200
