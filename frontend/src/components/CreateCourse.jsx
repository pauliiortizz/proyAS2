import React from 'react';
import Cookies from 'js-cookie';
import {Button, Input, FormControl, FormLabel, Select} from "@chakra-ui/react";
import '../estilos/CreateCourse.css';

const CreateCourse = ({ onClose }) => {
    const [nombre, setNombre] = React.useState('');
    const profesor_id = parseInt(Cookies.get('user_id'), 10);
    const tokenUser = Cookies.get('token');
    const [categoria, setCategoria] = React.useState('');
    const [descripcion, setDescripcion] = React.useState('');
    const [duracion, setDuracion] = React.useState('');
    const [valoracion, setValoracion] = React.useState('');
    const [requisitos, setRequisitos] = React.useState('');
    const [url_image, setUrl_image] = React.useState('');
    const [fecha_inicio, setFecha_inicio] = React.useState('');
    const [errorMessage, setErrorMessage] = React.useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        console.log('Formulario enviado');
        console.log('Datos del curso:', { nombre, profesor_id, categoria, descripcion, duracion, valoracion, requisitos, url_image, fecha_inicio });

        if (nombre === '' || categoria === '' || descripcion === '' || duracion === ''|| valoracion === '' || requisitos === '' || url_image === '' || fecha_inicio === '') {
            setErrorMessage('Todos los campos son obligatorios.');
            return;
        }
        const data = {
            nombre,
            profesor_id,
            categoria,
            descripcion,
            duracion: parseInt(duracion, 10),
            valoracion: parseFloat(valoracion),
            requisitos,
            url_image,
            fecha_inicio: new Date(fecha_inicio).toISOString() // Convertir la fecha a ISO 8601
        };
        console.log('Enviando datos:', data);

        try {
            const response = await fetch('http://localhost:8080/createCourse', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    //'Authorization': `Bearer ${Cookies.get('token')}`
                },
                body: JSON.stringify (data)
            });

            console.log('Respuesta del servidor:', response);

            if (!response.ok) {
                const errorData = await response.json();
                setErrorMessage(`Error al crear el curso: ${errorData.message}`);
            } else {
                alert('Curso creado exitosamente');
                onClose(); // Cierra el formulario
                window.location.reload(); // Recargar la página
            }
        } catch (error) {
            console.log(`Error de red al crear el curso: ${error.message}`);
            alert("Error al crear el curso")
        }
    };

    return (
        <form id="formCreateCourse" onSubmit={handleSubmit}>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Nombre del curso</FormLabel>
                <Input value={nombre} onChange={(e) => setNombre(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Categoría</FormLabel>
                <Input value={categoria} onChange={(e) => setCategoria(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Descripción</FormLabel>
                <Input value={descripcion} onChange={(e) => setDescripcion(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Duración</FormLabel>
                <Input value={duracion} onChange={(e) => setDuracion(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Valoracion</FormLabel>
                <Select value={valoracion} onChange={(e) => setValoracion(e.target.value)}
                        style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>
                    <option value="" style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>seleccione la valoracion
                    </option>
                    <option value="0" style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>0
                    </option>
                    <option value="1" style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>1
                    </option>
                    <option value="2" style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>2
                    </option>
                    <option value="3"
                            style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>3
                    </option>
                    <option value="4"
                            style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>4
                    </option>
                    <option value="5"
                            style={{border: '2px solid black', fontFamily: 'Spoof Trial, sans-serif'}}>5
                    </option>
                </Select>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Requisitos</FormLabel>
                <Input value={requisitos} onChange={(e) => setRequisitos(e.target.value)}
                       style={{border: '2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>URL de la imagen</FormLabel>
                <Input value={url_image} onChange={(e) => setUrl_image(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Fecha de inicio</FormLabel>
                <Input type="date" value={fecha_inicio} onChange={(e) => setFecha_inicio(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            {errorMessage && <p style={{fontFamily: 'Spoof Trial, sans-serif'}}>{errorMessage}</p>}
            <Button type="submit">Crear curso</Button>
        </form>
    );
};

export default CreateCourse;