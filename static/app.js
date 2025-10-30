document.addEventListener('DOMContentLoaded', () => {
    initRegistro();
    initLogin();
    initUsuarioPanel();
});

let resenaEditandoId = null;

function initRegistro() {
    const formRegistro = document.getElementById('form-registro');
    if (!formRegistro) return;

    formRegistro.addEventListener('submit', async (event) => {
        event.preventDefault();
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

            if (!res.ok) {
                const error = await extraerMensajeError(res);
                alert(error || 'Error al registrar usuario');
                return;
            }

            const cliente = await res.json();
            guardarClienteEnLocalStorage(cliente);
            window.location.href = 'usuario.html';
        } catch (err) {
            alert('Error de red al registrar usuario');
        }
    });
}

function initLogin() {
    const formLogin = document.getElementById('form-login');
    if (!formLogin) return;

    formLogin.addEventListener('submit', async (event) => {
        event.preventDefault();
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

            if (!res.ok) {
                const error = await extraerMensajeError(res);
                alert(error || 'Usuario o contraseña incorrectos');
                return;
            }

            const cliente = await res.json();
            guardarClienteEnLocalStorage(cliente);
            window.location.href = 'usuario.html';
        } catch (err) {
            alert('Error de red al iniciar sesión');
        }
    });
}

function initUsuarioPanel() {
    const formResena = document.getElementById('form-resena');
    if (!formResena) return;

    const clienteID = localStorage.getItem('cliente_id');
    if (!clienteID) {
        window.location.href = '/';
        return;
    }

    mostrarDetalleCliente();
    cargarResenas(clienteID);
    cargarResenasRecientes();

    const contenedorResenas = document.getElementById('mis-resenas');
    if (contenedorResenas) {
        contenedorResenas.addEventListener('click', manejarClickResenas);
    }

    const btnCancelarEdicion = document.getElementById('btn-cancelar-edicion');
    if (btnCancelarEdicion) {
        btnCancelarEdicion.addEventListener('click', cancelarEdicionResena);
    }

    const formPerfil = document.getElementById('form-actualizar-email');
    if (formPerfil) {
        const inputEmail = formPerfil.elements.email;
        const emailActual = localStorage.getItem('cliente_email') || '';
        if (inputEmail && emailActual) {
            inputEmail.value = emailActual;
        }

        formPerfil.addEventListener('submit', async (event) => {
            event.preventDefault();
            const nuevoEmail = inputEmail?.value.trim();
            if (!nuevoEmail) {
                mostrarMensajePerfil('Ingresá un email válido.', 'error');
                return;
            }

            try {
                const res = await fetch(`/clientes?id=${encodeURIComponent(clienteID)}`, {
                    method: 'PUT',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ email: nuevoEmail })
                });

                if (!res.ok) {
                    const error = await extraerMensajeError(res);
                    mostrarMensajePerfil(error || 'No se pudo actualizar el email.', 'error');
                    return;
                }

                localStorage.setItem('cliente_email', nuevoEmail);
                mostrarDetalleCliente();
                mostrarMensajePerfil('Email actualizado correctamente.', 'success');
            } catch (err) {
                mostrarMensajePerfil('Error de red al actualizar el email.', 'error');
            }
        });
    }

    formResena.addEventListener('submit', async (event) => {
        event.preventDefault();
        const payload = {
            titulo: formResena.titulo.value,
            descripcion: formResena.descripcion.value,
            nota: Number(formResena.nota.value),
            cliente_id: Number(clienteID)
        };

        try {
            const esEdicion = Number.isInteger(resenaEditandoId) && resenaEditandoId > 0;
            let url = '/resenas';
            if (esEdicion) {
                payload.id = resenaEditandoId;
                url += `?id=${encodeURIComponent(resenaEditandoId)}&cliente_id=${encodeURIComponent(clienteID)}`;
            }

            const res = await fetch(url, {
                method: esEdicion ? 'PUT' : 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });

            if (!res.ok) {
                const error = await extraerMensajeError(res);
                mostrarMensajeResena(error || 'No se pudo guardar la reseña', 'error');
                return;
            }

            await res.json();
            restaurarFormularioResena();
            mostrarMensajeResena(esEdicion ? 'Reseña actualizada correctamente' : 'Reseña guardada correctamente', 'success');
            cargarResenas(clienteID);
            cargarResenasRecientes();
        } catch (err) {
            mostrarMensajeResena('Error de red al crear la reseña', 'error');
        }
    });
}

async function cargarResenas(clienteID) {
    const contenedor = document.getElementById('mis-resenas');
    if (!contenedor) return;

    contenedor.classList.add('empty-state');
    contenedor.textContent = 'Cargando reseñas...';

    try {
        const res = await fetch(`/resenas?cliente_id=${encodeURIComponent(clienteID)}`);
        if (!res.ok) {
            const error = await extraerMensajeError(res);
            contenedor.textContent = error || 'No se pudieron obtener las reseñas.';
            return;
        }

        const resenas = await res.json();
        renderizarResenas(resenas);
    } catch (err) {
        contenedor.textContent = 'Error de red al obtener las reseñas.';
    }
}

function renderizarResenas(resenas) {
    const contenedor = document.getElementById('mis-resenas');
    if (!contenedor) return;

    if (!Array.isArray(resenas) || resenas.length === 0) {
        contenedor.classList.add('empty-state');
        contenedor.textContent = 'Todavía no cargaste reseñas.';
        return;
    }

    contenedor.classList.remove('empty-state');
    const frag = document.createDocumentFragment();

    const ordenadas = [...resenas].sort((a, b) => new Date(b.fecha) - new Date(a.fecha));

    ordenadas.forEach((resena) => {
        const item = document.createElement('article');
        item.className = 'resena-item';
        if (resena.id != null) {
            item.dataset.resenaId = String(resena.id);
        }
        if (resena.titulo != null) item.dataset.titulo = resena.titulo;
        if (resena.descripcion != null) item.dataset.descripcion = resena.descripcion;
        if (resena.nota != null) item.dataset.nota = String(resena.nota);

        const titulo = document.createElement('h3');
        titulo.textContent = resena.titulo;
        item.appendChild(titulo);

        const descripcion = document.createElement('p');
        descripcion.textContent = resena.descripcion;
        item.appendChild(descripcion);

        const meta = document.createElement('p');
        const fecha = resena.fecha ? new Date(resena.fecha) : null;
        const fechaLegible = fecha ? fecha.toLocaleDateString() : 'Sin fecha';
        const idTexto = resena.id != null ? ` · ID: #${resena.id}` : '';
        meta.textContent = `Nota: ${resena.nota} · Fecha: ${fechaLegible}${idTexto}`;
        item.appendChild(meta);

        if (resena.id != null) {
            const acciones = document.createElement('div');
            acciones.className = 'resena-acciones';

            const btnEditar = document.createElement('button');
            btnEditar.type = 'button';
            btnEditar.className = 'btn-editar';
            btnEditar.textContent = 'Editar';

            const btnEliminar = document.createElement('button');
            btnEliminar.type = 'button';
            btnEliminar.className = 'btn-eliminar';
            btnEliminar.textContent = 'Eliminar';

            acciones.appendChild(btnEditar);
            acciones.appendChild(btnEliminar);
            item.appendChild(acciones);
        }

        frag.appendChild(item);
    });

    contenedor.innerHTML = '';
    contenedor.appendChild(frag);
}

async function cargarResenasRecientes(limite = 10) {
    const contenedor = document.getElementById('resenas-recientes');
    if (!contenedor) return;

    contenedor.classList.add('empty-state');
    contenedor.textContent = 'Cargando reseñas...';

    try {
        const res = await fetch(`/resenas/recientes?limit=${encodeURIComponent(limite)}`);
        if (!res.ok) {
            const error = await extraerMensajeError(res);
            contenedor.textContent = error || 'No se pudieron obtener las reseñas recientes.';
            return;
        }

        const resenas = await res.json();
        renderizarResenasRecientes(resenas);
    } catch (err) {
        contenedor.textContent = 'Error de red al obtener las reseñas recientes.';
    }
}

function renderizarResenasRecientes(resenas) {
    const contenedor = document.getElementById('resenas-recientes');
    if (!contenedor) return;

    if (!Array.isArray(resenas) || resenas.length === 0) {
        contenedor.classList.add('empty-state');
        contenedor.textContent = 'Aún no hay reseñas cargadas.';
        return;
    }

    contenedor.classList.remove('empty-state');
    const frag = document.createDocumentFragment();

    resenas.forEach((resena) => {
        const item = document.createElement('article');
        item.className = 'resena-item';

        const titulo = document.createElement('h3');
        titulo.textContent = resena.titulo;
        item.appendChild(titulo);

        const descripcion = document.createElement('p');
        descripcion.textContent = resena.descripcion;
        item.appendChild(descripcion);

        const meta = document.createElement('p');
        const fecha = resena.fecha ? new Date(resena.fecha) : null;
        const fechaLegible = fecha ? fecha.toLocaleDateString() : 'Sin fecha';
        const autor = resena.usuario ? ` · Autor: ${resena.usuario}` : '';
        meta.textContent = `Nota: ${resena.nota} · Fecha: ${fechaLegible}${autor}`;
        item.appendChild(meta);

        frag.appendChild(item);
    });

    contenedor.innerHTML = '';
    contenedor.appendChild(frag);
}

function mostrarMensajeResena(mensaje, tipo) {
    const destino = document.getElementById('mensaje-resena');
    if (!destino) return;

    destino.textContent = mensaje;

    if (!mensaje) {
        destino.className = '';
        return;
    }

    let clase = '';
    if (tipo === 'success') clase = 'mensaje-exito';
    else if (tipo === 'error') clase = 'mensaje-error';
    else if (tipo === 'info') clase = 'mensaje-info';

    destino.className = clase;
}

function mostrarDetalleCliente() {
    const destino = document.getElementById('cliente-detalle');
    if (!destino) return;

    const nombre = localStorage.getItem('cliente_nombre') || '';
    const apellido = localStorage.getItem('cliente_apellido') || '';
    const usuario = localStorage.getItem('cliente_usuario') || '';
    const id = localStorage.getItem('cliente_id');
    const email = localStorage.getItem('cliente_email') || '';

    const partes = [nombre, apellido].filter(Boolean).join(' ');
    const usuarioTexto = usuario ? ` (@${usuario})` : '';
    const emailTexto = email ? ` Email: ${email}` : '';
    destino.textContent = `Estás logueado como ${partes || 'usuario'}${usuarioTexto}.`;
}

function guardarClienteEnLocalStorage(cliente) {
    if (!cliente || typeof cliente !== 'object') return;
    if (cliente.id != null) localStorage.setItem('cliente_id', cliente.id);
    if (cliente.nombre) localStorage.setItem('cliente_nombre', cliente.nombre);
    if (cliente.apellido) localStorage.setItem('cliente_apellido', cliente.apellido);
    if (cliente.usuario) localStorage.setItem('cliente_usuario', cliente.usuario);
    if (cliente.email) localStorage.setItem('cliente_email', cliente.email);
}

async function extraerMensajeError(res) {
    try {
        const data = await res.json();
        if (data && typeof data === 'object') {
            return data.error || data.message || '';
        }
    } catch (err) {
        return '';
    }
    return '';
}

async function eliminarResena(resenaID) {
    const clienteID = localStorage.getItem('cliente_id');
    if (!clienteID) {
        mostrarMensajeResena('No se encontró el usuario logueado. Volvé a iniciar sesión.', 'error');
        return;
    }

    const confirmar = confirm('¿Querés eliminar esta reseña?');
    if (!confirmar) return;

    try {
        const res = await fetch(`/resenas?id=${encodeURIComponent(resenaID)}&cliente_id=${encodeURIComponent(clienteID)}`, {
            method: 'DELETE'
        });

        if (!res.ok) {
            const error = await extraerMensajeError(res);
            mostrarMensajeResena(error || 'No se pudo eliminar la reseña.', 'error');
            return;
        }

        if (resenaEditandoId === resenaID) {
            restaurarFormularioResena();
        }

        mostrarMensajeResena('Reseña eliminada correctamente.', 'success');
        cargarResenas(clienteID);
        cargarResenasRecientes();
    } catch (err) {
        mostrarMensajeResena('Error de red al eliminar la reseña.', 'error');
    }
}

function manejarClickResenas(evento) {
    const boton = evento.target.closest('button');
    if (!boton) return;

    const item = boton.closest('.resena-item');
    if (!item || !item.dataset.resenaId) return;

    const resenaID = Number(item.dataset.resenaId);
    if (!Number.isInteger(resenaID) || resenaID <= 0) return;

    if (boton.classList.contains('btn-eliminar')) {
        eliminarResena(resenaID);
    } else if (boton.classList.contains('btn-editar')) {
        prepararEdicionResena(item);
    }
}

function prepararEdicionResena(elemento) {
    const formResena = document.getElementById('form-resena');
    if (!formResena) return;

    resenaEditandoId = Number(elemento.dataset.resenaId);
    formResena.titulo.value = elemento.dataset.titulo || '';
    formResena.descripcion.value = elemento.dataset.descripcion || '';
    formResena.nota.value = elemento.dataset.nota || '';

    const botonSubmit = formResena.querySelector('button[type="submit"]');
    if (botonSubmit) botonSubmit.textContent = 'Actualizar reseña';

    const btnCancelar = document.getElementById('btn-cancelar-edicion');
    if (btnCancelar) btnCancelar.style.display = 'inline-flex';

    mostrarMensajeResena(`Editando reseña #${resenaEditandoId}`, 'info');
    formResena.scrollIntoView({ behavior: 'smooth', block: 'start' });
    formResena.titulo.focus();
}

function restaurarFormularioResena() {
    resenaEditandoId = null;
    const formResena = document.getElementById('form-resena');
    if (formResena) formResena.reset();

    const botonSubmit = formResena?.querySelector('button[type="submit"]');
    if (botonSubmit) botonSubmit.textContent = 'Enviar reseña';

    const btnCancelar = document.getElementById('btn-cancelar-edicion');
    if (btnCancelar) btnCancelar.style.display = 'none';
}

function cancelarEdicionResena() {
    restaurarFormularioResena();
    mostrarMensajeResena('Edición cancelada.', 'info');
}

function mostrarMensajePerfil(mensaje, tipo) {
    const destino = document.getElementById('mensaje-perfil');
    if (!destino) return;

    destino.textContent = mensaje;

    if (!mensaje) {
        destino.className = '';
        return;
    }

    let clase = '';
    if (tipo === 'success') clase = 'mensaje-exito';
    else if (tipo === 'error') clase = 'mensaje-error';
    else if (tipo === 'info') clase = 'mensaje-info';
    destino.className = clase;
}