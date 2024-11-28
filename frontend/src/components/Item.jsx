import React, { useEffect, useState } from "react";
import axios from "axios";
import Cookies from "js-cookie";
import { Button, Stack, Card, CardBody, CardFooter, Text, Image, useDisclosure } from "@chakra-ui/react";
import Inscribirmebutton from "./Inscribirmebutton.jsx";
import PopupEdit from "./PopUpEdit.jsx";
import '../estilos/Inscribirmebutton.css';
import '../estilos/Course.css';

// eslint-disable-next-line react/prop-types
const Item = ({ course, bandera }) => {
    const [userId, setUserId] = useState(null);
    const [isAdmin, setIsAdmin] = useState(false);
    const [isEnrolled, setIsEnrolled] = useState(false);
    const { isOpen: isPopupOpenEdit, onOpen: onOpenPopupEdit, onClose: onClosePopupEdit } = useDisclosure();


    // eslint-disable-next-line react/prop-types
    const formattedDate = new Date(course.fecha_inicio).toLocaleDateString('es-ES', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
    });

    useEffect(() => {
        const storedUserId = Cookies.get('user_id');
        if (storedUserId) {
            setUserId(parseInt(storedUserId, 10));
        }

        const storedAdmin = Cookies.get('admin');
        if (storedAdmin) {
            setIsAdmin(storedAdmin === "1");
        }
    }, []);

    useEffect(() => {
        const checkEnrollment = async () => {
            if (userId) {
                try {
                    const response = await axios.get(`http://localhost:8083/inscripciones/user/${userId}`);
                    const inscripciones = response.data;
                    {/* eslint-disable-next-line react/prop-types */}
                    const enrolled = inscripciones.some(inscripcion => inscripcion.Id_course === course.course_id);
                    setIsEnrolled(enrolled);
                } catch (error) {
                    console.error('Error checking enrollment:', error);
                }
            }
        };

        checkEnrollment();
        // eslint-disable-next-line react/prop-types
    }, [userId, course.course_id]);

    const getProfesorName = (profesor_id) => {
        const profesores = {
            2: 'Juan Lopez',
            4: 'Margarita de Marcos',
            8: 'Gustavo Jacobo',
            17: 'Rodolfo Perez',
            18: 'Sebastian Colidio',
            19: 'Lucas Beltran'
        };
        return profesores[profesor_id] || 'Profesor desconocido';
    };

    const handleEditCourse = () => {
        onOpenPopupEdit();
    };


    return (
        <Card direction={{ base: 'column', sm: 'row' }} overflow='hidden' variant='outline'>
            {bandera !== 1 ? (
                <Image
                    objectFit='cover'
                    maxW={{ sm: '250px' }}
                    // eslint-disable-next-line react/prop-types
                    src={course.url_image}
                    alt='Imagen Curso'
                />
            ) : null}

            <Stack>
                <CardBody className='body'>
                    {/* eslint-disable-next-line react/prop-types */}
                    <h1 style={{ fontFamily: 'Spoof Trial, sans-serif', fontWeight: 800, fontSize: 30 }}>{course.nombre}</h1>

                    {/* eslint-disable-next-line react/prop-types */}
                    <Text py='2' className="card-text">{course.descripcion}</Text>
                    {/* eslint-disable-next-line react/prop-types */}
                    <Text py='2' className="card-text">{course.categoria}</Text>
                    <Text marginBottom='3px' display='flex' py='2' alignItems='center' className="card-text">
                        <img src="/estrella.png" alt="estrella" width="20px" height="20px" style={{ marginRight: '5px' }} />
                        {/* eslint-disable-next-line react/prop-types */}
                        {course.valoracion}/5
                    </Text>
                    {/* eslint-disable-next-line react/prop-types */}
                    <Text className="card-textt">Duracion: {course.duracion}hs</Text>
                    <Text className="card-textt">Fecha de inicio: {formattedDate}</Text>
                    {/* eslint-disable-next-line react/prop-types */}
                    <Text className="card-textt">Requisito: Nivel {course.requisitos}</Text>
                    <Text className="card-textt">Cupos Disponibles: {course.capacidad}</Text>
                    {/* eslint-disable-next-line react/prop-types */}
                    <Text className="card-textt">Profesor: {getProfesorName(course.profesor_id)}</Text>
                </CardBody>
                <CardFooter>
                    {userId && (
                        isAdmin ? (
                            <Button w="40%" style={{ fontFamily: 'Spoof Trial, sans-serif' }} onClick={handleEditCourse}>Editar</Button>
                        ) : (
                            !isEnrolled && bandera !== 1 && course.capacidad > 0 && (
                                <Inscribirmebutton courseId={course.course_id} />
                            )
                        )
                    )}
                </CardFooter>
            </Stack>
            {/* eslint-disable-next-line react/prop-types */}
            <PopupEdit isOpen={isPopupOpenEdit} onClose={onClosePopupEdit} courseId={course.course_id} />
        </Card>
    );
};

export default Item;
