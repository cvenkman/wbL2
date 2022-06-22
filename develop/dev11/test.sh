curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "first"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "second"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "third"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-09", "title": "first"}' http://localhost:8080/create_event