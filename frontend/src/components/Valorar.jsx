import React, { useEffect, useState } from "react";
import Cookies from 'js-cookie';
import { Box, Button, FormControl, FormLabel, Textarea, Select, useToast } from "@chakra-ui/react";
import axios from 'axios';

const ValorarCourse = ({ courseId, onClose }) => {
    const [inscripcion_id, setInscripcion] = useState('');
    const user_id = parseInt(Cookies.get('user_id'));
    const [valoracion, setValoracion] = useState('');
    const [comentario, setComentario] = useState('');
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

        try {
            await axios.post(`http://localhost:8080/valorar/${inscripcion_id}`, {
                Id_user: user_id,
                Id_course: courseId,
                Comentario: comentario,
                Valoracion: parseFloat(valoracion),  // Convertir a número antes de enviar
            });
            toast({
                title: "Valoración enviada.",
                description: "Tu valoración ha sido enviada con éxito.",
                status: "success",
                duration: 5000,
                isClosable: true,
            });
            onClose();
        } catch (error) {
            console.error('Error submitting valoracion:', error);
            toast({
                title: "Error.",
                description: "Hubo un problema al enviar tu valoración.",
                status: "error",
                duration: 5000,
                isClosable: true,
            });
        }
    };

    return (
        <Box p={4}>
            <form onSubmit={handleSubmit}>
                <FormControl>
                    <FormLabel>Valoración</FormLabel>
                    <Select
                        placeholder="Selecciona una valoración"
                        value={valoracion}
                        onChange={(e) => setValoracion(e.target.value)}
                        required
                    >
                        <option value="1">1 estrella</option>
                        <option value="2">2 estrellas</option>
                        <option value="3">3 estrellas</option>
                        <option value="4">4 estrellas</option>
                        <option value="5">5 estrellas</option>
                    </Select>
                </FormControl>
                <FormControl mt={4}>
                    <FormLabel>Comentario</FormLabel>
                    <Textarea
                        value={comentario}
                        onChange={(e) => setComentario(e.target.value)}
                        required
                    />
                </FormControl>
                <Button mt={4} colorScheme="teal" type="submit">Enviar Valoración</Button>
            </form>
        </Box>
    );
};

export default ValorarCourse;
