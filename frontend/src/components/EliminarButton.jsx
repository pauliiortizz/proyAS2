import React from 'react';
import Cookies from 'js-cookie';
import '../estilos/Inscribirmebutton.css';

const DeleteCourse = ({ courseId, onClose }) => {
    const tokenUser = Cookies.get('token');
    const [isOpen, setIsOpen] = React.useState(false);
    const onCloseAlert = () => setIsOpen(false);
    const cancelRef = React.useRef();

    const handleDelete = async () => {
        try {
            const response = await fetch(`http://localhost:8080/courses/${courseId}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                   // 'Authorization': `Bearer ${tokenUser}`
                }
            });

            if (response.ok) {
                alert('Curso eliminado exitosamente');
                window.location.reload(); // Recargar la página
                } else {
                const errorData = await response.json();
                alert(`Error al eliminar el curso: ${errorData.message}`);
            }
        } catch (error) {
            console.error(`Error de red al eliminar el curso: ${error.message}`);
            alert("Error al eliminar el curso");
        }
    };



    return (
        <button className="delete-button" onClick={handleDelete}>ELIMINAR</button>
    );
};



export default DeleteCourse;