
GET http://localhost:8080/resenas/recientes?limit=100
HTTP/1.1 200
[Asserts]
header "Content-Type" contains "application/json"