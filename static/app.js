document.addEventListener('DOMContentLoaded', () => {
    const formRegistro = document.getElementById('form-registro');
    if (formRegistro) {
        formRegistro.onsubmit = async (e) => {
            e.preventDefault();
            const data = {
                Nombre: formRegistro.nombre.value,
                Apellido: formRegistro.apellido.value,
                Usuario: formRegistro.usuario.value,
                Pass: formRegistro.contrasena.value,
                Email: formRegistro.mail.value
            };
            try {
                const res = await fetch('/clientes', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(data)
                });
                if (res.ok) {
                    // Redirige a la página de usuario con las 4 opciones
                    window.location.href = 'usuario.html';
                } else {
                    alert('Error al registrar usuario');
                }
            } catch (err) {
                alert('Error de red');
            }
        };
    }
});

document.addEventListener('DOMContentLoaded', () => {
    // Registro
    const formRegistro = document.getElementById('form-registro');
    if (formRegistro) {
        formRegistro.onsubmit = async (e) => {
            e.preventDefault();
            const data = {
                Nombre: formRegistro.nombre.value,
                Apellido: formRegistro.apellido.value,
                Usuario: formRegistro.usuario.value,
                Pass: formRegistro.contrasena.value,
                Email: formRegistro.mail.value
            };
            try {
                const res = await fetch('/clientes', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(data)
                });
                if (res.ok) {
                    window.location.href = 'usuario.html';
                } else {
                    alert('Error al registrar usuario');
                }
            } catch (err) {
                alert('Error de red');
            }
        };
    }

    // Login
    const formLogin = document.getElementById('form-login');
    if (formLogin) {
        formLogin.onsubmit = async (e) => {
            e.preventDefault();
            const data = {
                usuario: formLogin.usuario.value,
                pass: formLogin.contrasena.value
            };
            try {
                const res = await fetch('/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(data)
                });
                if (res.ok) {
                    const cliente = await res.json();
                    localStorage.setItem('cliente_id', cliente.id);
                    window.location.href = 'usuario.html';
                } else {
                    alert('Usuario o contraseña incorrectos');
                }
            } catch (err) {
                alert('Error de red');
            }
        };
    }
});