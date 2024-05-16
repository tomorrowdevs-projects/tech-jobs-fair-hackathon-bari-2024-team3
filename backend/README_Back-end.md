# Back-end

## Installations

### GoLang
Instal GoLang on your pc https://go.dev/doc/install

### Postman
From VSCode: Install Postman plugin


## Run
### Run the back-end
In your terminal:
    cd to folder tech-jobs-fair-hackathon-bari-2024-team3/backend/

Start backend by terminal: 
    go run main.go

### Run the Postman WebSocket for interactions
In VSCode open the Postman plugin.
Start a new WebSocket Request.
connect websocket to the following:
    ws://localhost:3333/ws

## Interactions

In the WebSocket you can send in requests to the backend and see the results in the terminal of the backend:
__Commands:__

- createQuiz <quizName> <userId||username>
    - createQuiz FirstQuiz Christian

- joinQuiz <quizId> <userId||username>
    - joinQuiz 05126db1-dbb2-49a3-b187-9bbe46bc48a4 Jack

- leaveQuiz <quizId> <userId||username>
    - leaveQuiz 05126db1-dbb2-49a3-b187-9bbe46bc48a4 Jack

- printing all quizzes data:
    - print
