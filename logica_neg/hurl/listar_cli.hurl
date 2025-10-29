GET http://localhost:8080/clientes
HTTP/1.1 200
[Asserts]
header "Content-Type" contains "application/json"