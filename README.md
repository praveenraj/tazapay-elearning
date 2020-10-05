# Tazapay eLearning

Adopted _micro_ Service Oriented Architecture to build the eLearning application.

> The microservice architecture enables the rapid, frequent and reliable delivery of large, complex applications. It also enables an organization to evolve its technology stack.

## Project Structure

```bash
├── common
│   ├── constants
│   ├── driver
│   ├── handlers
│   ├── Makefile
│   ├── models
│   ├── repository
│   └── utils
├── course-service
│   ├── controller
│   ├── logic
│   ├── Makefile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── user-service
└── README.md
```

## Prepare library/build

**Common :** Use `make library` command to prepare the library for the service(s).

_Note:_ Not essential to prepare common package to build service(s). But in order to run the ci-lint, unit test, etc. then `make`. Also common library has facility to build all the services at single shot by execute `make build-all-svc ENV=dev`.

**Service :** Use the following commands effectively with respect to environment builds:

- Development: `make dev`
- Stage || QA: `make stage`
- UAT: `make release`
- Production: `make prod`

> Type `make` to view all the available targets.

## Use Case & FLow

### 1. Register new user

1. Get user basic details, validate and save it in the system
2. To do authentication, use OTP process either via mobile or email

### 2. Courses will be available to users to get enrolled, only If

1. Courses which are in active state
2. No of active sections per course > 0
3. Active section -> no of (active and saved state) lessons per section > 0 and section is_active=1

### 3. Add new course

1. The newly added course will be in inactive state
2. The new lesson content under section will be treated as drafted one (even though if it is a master branch)
3. Author who has admin/maintain rights able to merge the changes to the master branch and so it will be in ready to publish state
4. Author who has admin rights can able to publish(save) the merged changes
5. The admin author will decide to move the course in active state to get users enrollment
   - The course will be available to users even it has only one section & one lesson in it

_Note:_ If there is only one author for the course then they would have admin rights by default

### 4. Update course/section/lesson metadata

1. Author can move the course, section or lesson to inactive state and vice-versa
2. Author can arrange the section/lesson order in which position it should appear
3. Have privileges to edit the course/section/lesson name and its attributes

### 5. Update lesson content (copy/local branch to master branch **state-machine**)

1. When begins to update the lesson content (i.e save the edited content), the edited content from master will be saved to cloud storage.
   - If it is a new version, then the new cloud storage link will be saved in `lesson_content` table with `parent_id` as master branch's id and `state` as draft
   - If it is an existing version and state is in draft, then update the content in the cloud storage
2. Only if the content is in draft state, then can able to move to merge
3. Once the edited content is ready to merge
   - If the course has only one author then they can straight away merge the changes
   - If it has more than one author then only the author who has admin/maintain rights can able to merge
4. When the admin author save the changes to get published, update the master branch cloud storage link with the updated link

_Note_: Just followed high level git mechanism. The above approach have few **limitations** in it

- The previous version of master branch can't be retained at this time. Need additional effort to do so

## Execution

### Unit Test & Coverage

Unit test functions will be available in `*_test.go` file and it present under every child folder.

To execute the unit test, direct to the parent folder (like common, course-service, etc.) and execute `make cover` command. Once the test execution done, open `coverage.html` file in the browser to view the coverage report.

```bash
cd common
make cover
------------
cd course-service
make cover
```

### API

Implemented the following APIs

1. Fetch all courses (use case #2)
2. Create user (use case #1)
3. Fetch lesson content
4. Update lesson content (state-machine) (use case #5)

Import [this](https://www.getpostman.com/collections/fc44f1e8a654d4567314) postman collection to validate the APIs

### CURL Instructions

1. Fetch all courses (use case #2)
   - Method: GET
   - Path: /courses
   - Request: nil
   - Responses
     - 200 Success
     - 500 Internal Server Error
2. Create user (use case #1)
   - Method: PUT
   - Path: /user
   - Request: {"first_name": "", "last_name": "", "mobile": "", email: ""} # last_name is optional
   - Responses
     - 200 Success
     - 400 Bad Request
     - 409 Conflict (user already exists)
     - 500 Internal Server Error
3. Fetch lesson content
   - Method: GET
   - Path: /courses/{course_id}/sections/{section_id}/lessons/{lesson_id}
   - Request: nil
   - Responses
     - 200 Success
     - 400 Bad Request
     - 500 Internal Server Error
4. Update lesson content (state-machine) (use case #5)
   - Method: POST
   - Path: /courses/{course_id}/sections/{section_id}/lessons/{lesson_id}
   - Requests
     - Draft: {"action": "draft", "version": "", "time_required": 1500,"content": {}}
     - Merge: {"action": "merge", "version": ""}
     - Save/Publish: {"action": "save", "version": ""}
   - Responses
     - 200 Success
     - 400 Bad Request
     - 409 Conflict (either the state is not in draft to get merged or it is not in merged to get saved)
     - 500 Internal Server Error
   - Steps to verify
     - Fetch lesson content (master branch content)
     - Draft lesson content
     - Fetch lesson content (master branch content)
     - Merge drafted lesson content
     - Fetch lesson content (master branch content)
     - Save/Publish merged lesson content
     - Fetch lesson content (updated branch content in master branch)
