
# Creación de par de claves (pública/privada) y subida de clave pública a GitHub.

Para generar las claves públicas y privadas en mi sistema operativo macOS y subir la pública a GitHub, he seguido los siguientes pasos, gracias al sistema de claves SSH, que permite autenticarte sin necesidad de ingresar tus credenciales en cada interacción con GitHub.

### Paso 1: Genero una nueva clave SSH

1. Abrir la terminal.
2. Ejecutar comando para crear una nueva clave SSH:

   ```bash
   ssh-keygen -t ed25519 -C mmnuria@correo.ugr.es
   ```

3. Pulso `Enter` para guardar la clave generada.

4. Vuelvo a pulsar `Enter` para no generar una frase de seguridad.

### Paso 2: Añado la clave pública a GitHub

Buscamos la clave **publica** generada para copiarla y añadirla a GitHub.

```bash
pbcopy < ~/.ssh/id_ed25519.pub
```

#### Añado la clave en GitHub:
1. Inicio sesión en mi cuenta de [GitHub](https://github.com).
2. Voy a la configuración de mi cuenta (Settings).
3. En la barra lateral izquierda, selecciono **SSH and GPG keys**.
4. Hago clic en **New SSH key** .
5. Pego la clave pública copiada en el campo "Key", asigno un nombre que identifique esta clave (por ejemplo: "Clave publica mac"), y hago clic en **Add SSH key**.

Finalmente, están generadas las claves y subida la clave publica a GitHub.

# Configuración correo y usuario

He usado los siguientes comandos para realizarlo:

```bash
git config --global user.name mmnuria
git config --global user.email mmnuria@correo.ugr.es
```
