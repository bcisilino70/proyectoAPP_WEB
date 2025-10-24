# ============================================
# SECCIÓN ELIMINACIONES
# ============================================

# 17. Eliminar Reseña #2 del Cliente #1
DELETE http://localhost:8080/resenas?id=2&cliente_id=1
HTTP/1.1 204

# 18. Verificar eliminación listando reseñas del Cliente #1
GET http://localhost:8080/resenas?cliente_id=1
HTTP/1.1 200

# 19. Eliminar Cliente #3
DELETE http://localhost:8080/clientes?id=3
HTTP/1.1 204

# 20. Verificar eliminación listando todos los clientes
GET http://localhost:8080/clientes
HTTP/1.1 200