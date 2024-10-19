import React, { useEffect, useState } from "react";
import Cookies from 'js-cookie';
import { Box, Button, FormControl, FormLabel, Input, useToast } from "@chakra-ui/react";
import axios from 'axios';

const SubirArchivo = ({ courseId, onClose }) => {
    const [inscripcion_id, setInscripcion] = useState('');
    const user_id = parseInt(Cookies.get('user_id'));
    const [archivo, setArchivo] = useState(null);
    const toast = useToast();

    useEffect(() => {
        const fetchInscripcion = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/inscripciones/${user_id}`);
                const inscripciones = response.data;
                const inscripcion = inscripciones.find(inscripcion => inscripcion.id_course === courseId);
                if (inscripcion) {
                    setInscripcion(inscripcion.id_inscripcion);
                }
            } catch (error) {
                console.error('Error fetching inscripcion:', error);
            }
        };

        fetchInscripcion();
    }, [user_id, courseId]);

    const handleSubmit = async (e) => {
        e.preventDefault();

        const formData = new FormData();
        formData.append("id_inscripcion", inscripcion_id);
        formData.append("id_user", user_id);
        formData.append("id_course", courseId);
        formData.append("archivo", archivo);

        try {
            await axios.post(`http://localhost:8080/archivo/${inscripcion_id}`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            });
            toast({
                title: "Archivo subido.",
                description: "Tu archivo ha sido subido con éxito.",
                status: "success",
                duration: 5000,
                isClosable: true,
            });
            onClose();
        } catch (error) {
            console.error('Error submitting archivo:', error);
            toast({
                title: "Error.",
                description: "Hubo un problema al subir tu archivo.",
                status: "error",
                duration: 5000,
                isClosable: true,
            });
        }
    };

    const handleFileChange = (e) => {
        setArchivo(e.target.files[0]);
    };

    return (
        <Box p={4}>
            <form onSubmit={handleSubmit}>
                <FormControl>
                    <FormLabel>Subir Archivo</FormLabel>
                    <Input
                        type="file"
                        onChange={handleFileChange}
                        required
                    />
                </FormControl>
                <Button mt={4} colorScheme="teal" type="submit">Enviar Archivo</Button>
            </form>
        </Box>
    );
};

export default SubirArchivo;