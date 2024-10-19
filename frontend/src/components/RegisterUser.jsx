import React from 'react';
import {Button, Input, FormControl, FormLabel, Select, InputRightElement, InputGroup} from "@chakra-ui/react";
import '../estilos/RegisterUser.css'

const RegisterUser = ({ onClose }) => {
    const [nombre, setNombre] = React.useState('');
    const [apellido, setApellido] = React.useState('');
    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [admin, setAdmin] = React.useState(''); // Mantener como string para el dropdown
    const [errorMessage, setErrorMessage] = React.useState('');
    const [show, setShow] = React.useState(false);
    const handleClick = () => setShow(!show);

    const handleSubmit = async (e) => {
        e.preventDefault();

        console.log('Formulario enviado');
        console.log('Datos del usuario:', { nombre, apellido, username, password, admin });

        if (nombre === '' || apellido === '' || username === '' || password === '' || admin === '') {
            setErrorMessage('Todos los campos son obligatorios.');
            return;
        }

        const data = {
            username,
            password,
            nombre,
            apellido,
            admin: admin === 'true' // Convertir string a boolean
        };

        console.log('Enviando datos:', data);

        try {
            const response = await fetch('http://localhost:8080/CreateUser', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    //'Authorization': `Bearer ${Cookies.get('token')}`
                },
                body: JSON.stringify(data)
            });

            console.log('Respuesta del servidor:', response);

            if (!response.ok) {
                const errorData = await response.json();
                setErrorMessage(`Error al crear el usuario: ${errorData.message}`);
            } else {
                alert('Usuario creado exitosamente');
                onClose(); // Cierra el formulario
            }
        } catch (error) {
            console.log(`Error de red al crear el usuario: ${error.message}`);
            alert("Error al crear el usuario");
        }
    };

    return (
        <form id="formRegisterUser" onSubmit={handleSubmit}>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Nombre</FormLabel>
                <Input value={nombre} onChange={(e) => setNombre(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Apellido</FormLabel>
                <Input value={apellido} onChange={(e) => setApellido(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Correo Electrónico</FormLabel>
                <Input value={username} onChange={(e) => setUsername(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Contraseña</FormLabel>
                <InputGroup size='md' className='inputContrasena'>
                    <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
                    <InputRightElement width='4.5rem'>
                        <Button h='1.75rem' size='sm' onClick={handleClick} style={{fontFamily: 'Spoof Trial, sans-serif'}}>
                            {show ? 'Ocultar' : 'Mostrar'}
                        </Button>
                    </InputRightElement>
                </InputGroup>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Rol</FormLabel>
                <Select value={admin} onChange={(e) => setAdmin(e.target.value)}>
                    <option value="" style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}>Seleccione un rol</option>
                    <option value="true" style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}>Profesor</option>
                    <option value="false" style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}>Alumno</option>
                </Select>
            </FormControl>
            {errorMessage && <p style={{fontFamily: 'Spoof Trial, sans-serif'}}>{errorMessage}</p>}
            <Button type="submit">Registrar usuario</Button>
        </form>
    );
};

export default RegisterUser;