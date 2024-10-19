import React, { useEffect, useState } from "react";
import axios from "axios";
import Cookies from "js-cookie";
import { Button, Stack, Card, CardBody, CardFooter, Text, Image, useDisclosure } from "@chakra-ui/react";
import Inscribirmebutton from "./Inscribirmebutton.jsx";
import EliminarButton from "./EliminarButton.jsx";
import PopupEdit from "./PopUpEdit.jsx";
import '../estilos/Inscribirmebutton.css';
import '../estilos/Course.css';
import PopupValorar from "./PopUpValorar.jsx"
import PopupSubirArchivo from "./PopUpArchivo.jsx";
import PopupSeeReview from "./PopUpSeeReview.jsx";

const Item = ({ course, bandera }) => {
    const [userId, setUserId] = useState(null);
    const [isAdmin, setIsAdmin] = useState(false);
    const [isEnrolled, setIsEnrolled] = useState(false);
    const { isOpen: isPopupOpenEdit, onOpen: onOpenPopupEdit, onClose: onClosePopupEdit } = useDisclosure();
    const { isOpen: isPopupOpenValorar, onOpen: onOpenPopupValorar, onClose: onClosePopupValorar } = useDisclosure();
    const { isOpen: isPopupOpenSubirArchivo, onOpen: onOpenPopupSubirArchivo, onClose: onClosePopupSubirArchivo } = useDisclosure();
    const { isOpen: isPopupOpenSeeReview , onOpen: onOpenPopupSeeReview , onClose: onClosePopupSeeReview } = useDisclosure();

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
                    const response = await axios.get(`http://localhost:8080/inscripciones/${userId}`);
                    const inscripciones = response.data;
                    const enrolled = inscripciones.some(inscripcion => inscripcion.id_course === course.course_id);
                    setIsEnrolled(enrolled);
                } catch (error) {
                    console.error('Error checking enrollment:', error);
                }
            }
        };

        checkEnrollment();
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

    const handleValorarCourse =()=>{
        onOpenPopupValorar();
    }

    const handleSubirArchivo = () => {
        onOpenPopupSubirArchivo();
    };

    const handleSeeReview = () => {
        onOpenPopupSeeReview();
    };


    return (
        <Card direction={{ base: 'column', sm: 'row' }} overflow='hidden' variant='outline'>
            {bandera !== 1 ? (
                <Image
                    objectFit='cover'
                    maxW={{ sm: '250px' }}
                    src={course.url_image}
                    alt='Imagen Curso'
                />
            ) : null}

            <Stack>
                <CardBody className='body'>
                    <h1 style={{ fontFamily: 'Spoof Trial, sans-serif', fontWeight: 800, fontSize: 30 }}>{course.nombre}</h1>

                    <Text py='2' className="card-text">{course.descripcion}</Text>
                    <Text py='2' className="card-text">{course.categoria}</Text>
                    <Text marginBottom='3px' display='flex' py='2' alignItems='center' className="card-text">
                        <img src="/estrella.png" alt="estrella" width="20px" height="20px" style={{ marginRight: '5px' }} />
                        {course.valoracion}/5
                    </Text>
                    <Text className="card-textt">Duracion: {course.duracion}hs</Text>
                    <Text className="card-textt">Fecha de inicio: {formattedDate}</Text>
                    <Text className="card-textt">Requisito: Nivel {course.requisitos}</Text>
                    <Text className="card-textt">Profesor: {getProfesorName(course.profesor_id)}</Text>
                </CardBody>
                <CardFooter>
                    {userId && (
                        isAdmin ? (
                            <>
                                <Button w="40%" style={{ fontFamily: 'Spoof Trial, sans-serif' }} onClick={handleEditCourse}>Editar</Button>
                                <EliminarButton courseId={course.course_id} />
                                <Button w="75%" style={{ fontFamily: 'Spoof Trial, sans-serif' }} onClick={handleSeeReview}>Ver review</Button>
                            </>
                        ) : (
                            isEnrolled ? (
                                bandera !== 1 ? (
                                    <>
                                        <Button w="75%" style={{ fontFamily: 'Spoof Trial, sans-serif', margin: '0 10px' }} onClick={handleValorarCourse}>Valorar</Button>
                                        <Button w="75%" style={{ fontFamily: 'Spoof Trial, sans-serif' }} onClick={handleSubirArchivo}>Subir archivo</Button>
                                        <Button w="75%" style={{ fontFamily: 'Spoof Trial, sans-serif' }} onClick={handleSeeReview}>Ver review</Button>
                                    </>
                                ) : null
                            ) : (
                                bandera !== 1 ? (
                                    <>
                                    <Inscribirmebutton courseId={course.course_id} />
                                    <Button w="75%" style={{ fontFamily: 'Spoof Trial, sans-serif' }} onClick={handleSeeReview}>Ver review</Button>
                                    </>
                ) : null
                            )
                        )
                    )}

                </CardFooter>
            </Stack>
            <PopupEdit isOpen={isPopupOpenEdit} onClose={onClosePopupEdit} courseId={course.course_id} />
            <PopupValorar isOpen={isPopupOpenValorar} onClose={onClosePopupValorar} courseId={course.course_id} />
            <PopupSubirArchivo isOpen={isPopupOpenSubirArchivo} onClose={onClosePopupSubirArchivo} courseId={course.course_id} />
            <PopupSeeReview isOpen={isPopupOpenSeeReview} onClose={onClosePopupSeeReview} courseId={course.course_id} />
        </Card>
    );
};

export default Item;
