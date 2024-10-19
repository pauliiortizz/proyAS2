import React, { useEffect } from 'react';
import Cookies from 'js-cookie';
import { Button, Input, FormControl, FormLabel } from "@chakra-ui/react";
import '../estilos/EditCourse.css';

const EditCourse = ({ courseId, onClose }) => {
    const [nombre, setNombre] = React.useState('');
    const profesor_id = parseInt(Cookies.get('user_id'), 10);
    const tokenUser = Cookies.get('token');
    const [categoria, setCategoria] = React.useState('');
    const [descripcion, setDescripcion] = React.useState('');
    const [duracion, setDuracion] = React.useState('');
    const [requisitos, setRequisitos] = React.useState('');
    const [url_image, setUrl_image] = React.useState('');
    const [fecha_inicio, setFecha_inicio] = React.useState('');
    const [errorMessage, setErrorMessage] = React.useState('');

    useEffect(() => {
        const fetchCourseData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/editCourse/${courseId}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });
                if (response.ok) {
                    const courseData = await response.json();
                    setNombre(courseData.nombre);
                    setCategoria(courseData.categoria);
                    setDescripcion(courseData.descripcion);
                    setDuracion(courseData.duracion.toString());
                    setRequisitos(courseData.requisitos);
                    setUrl_image(courseData.url_image);
                    setFecha_inicio(new Date(courseData.fecha_inicio).toISOString().split('T')[0]); // Formato yyyy-mm-dd
                } else {
                    console.error('Error al obtener los datos del curso');
                }
            } catch (error) {
                console.error(`Error de red al obtener los datos del curso: ${error.message}`);
            }
        };

        fetchCourseData();
    }, [courseId]);

    const handleSubmit = async (e) => {
        e.preventDefault();


        const data = {
            nombre: nombre || undefined,
            profesor_id,
            categoria: categoria || undefined,
            descripcion: descripcion || undefined,
            duracion: duracion ? parseInt(duracion, 10) : undefined,
            requisitos: requisitos || undefined,
            url_image: url_image || undefined,
            fecha_inicio: fecha_inicio ? new Date(fecha_inicio).toISOString() : undefined
        };

        try {
            const response = await fetch(`http://localhost:8080/edit/${courseId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            });

            if (!response.ok) {
                const errorData = await response.json();
                setErrorMessage(`Error al editar el curso: ${errorData.message}`);
            } else {
                alert('Curso editado exitosamente');
                onClose(); // Cierra el formulario
                window.location.reload()
            }
        } catch (error) {
            console.log(`Error de red al editar el curso: ${error.message}`);
            alert("Error al editar el curso");
        }
    };

    return (
        <form id="formEditCourse" onSubmit={handleSubmit}>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Nombre del curso</FormLabel>
                <Input value={nombre} onChange={(e) => setNombre(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}} />
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
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Requisitos</FormLabel>
                <Input value={requisitos} onChange={(e) => setRequisitos(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>URL de la imagen</FormLabel>
                <Input value={url_image} onChange={(e) => setUrl_image(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            <FormControl>
                <FormLabel style={{fontFamily: 'Spoof Trial, sans-serif'}}>Fecha de inicio</FormLabel>
                <Input type="date" value={fecha_inicio} onChange={(e) => setFecha_inicio(e.target.value)} style={{border:'2px solid black',fontFamily: 'Spoof Trial, sans-serif'}}/>
            </FormControl>
            {errorMessage && <p style={{fontFamily: 'Spoof Trial, sans-serif'}} className="error">{errorMessage}</p>}
            <Button type="submit">Editar curso</Button>
        </form>
    );
};

export default EditCourse;
