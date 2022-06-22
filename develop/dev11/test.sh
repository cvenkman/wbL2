
# в одном дне
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "firstInMonth"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "secondInMonth"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "thirdInMonth"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-09", "title": "firstInMonth"}' http://localhost:8080/create_event

# curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "dddddddddd"}' http://localhost:8080/update_event

# curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-09", "title": "first"}' http://localhost:8080/delete_event

# в одной неделе
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-01", "title": "firstInWeek"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-02", "title": "secondInWeek"}' http://localhost:8080/create_event

# в одном месяце
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-03", "title": "firstInMonth"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-03", "title": "firstInMonth"}' http://localhost:8080/create_event
curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "2", "date": "2019-09-13", "title": "secondInMonth"}' http://localhost:8080/create_event