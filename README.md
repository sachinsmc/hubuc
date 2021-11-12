# hubuc-task
Create one endpoint with add user functionality

How to run :

- Clone the repo or 
``$ go get https://github.com/sachinsmc/hubuc-task``



- `$ cd hubuc-task`


- `$ go mod tidy`


- Create user request : ``curl --location --request POST 'http://127.0.0.1:3000/api/v1/users' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "John Doe",
  "email":"john@local.host",
  "password":"XasdSz42#21"
  }'``


- Check validation : ``curl --location --request POST 'http://127.0.0.1:3000/api/v1/users' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "email":"john@local.host",
  "password":"XasdSz42#21"
  }'``
