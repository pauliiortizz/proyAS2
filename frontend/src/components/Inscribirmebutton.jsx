import '../estilos/Inscribirmebutton.css';
import { useState, useEffect } from "react";
import Cookies from "js-cookie";

const Inscribirmebutton = ({ courseId, fechaInicioCurso }) => {
  const [user_id, setUserId] = useState(null);
  const [isAdmin, setIsAdmin] = useState(false);
  const tokenUser = Cookies.get('token');

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

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Validación de usuario y token
    if (!user_id || user_id === -1 || user_id === 0 || !tokenUser) {
      alert("Debes registrarte para inscribirte a un curso");
      return;
    }

    // Validación de la fecha de inicio del curso
    const currentDate = new Date();
    const courseStartDate = new Date(fechaInicioCurso);

    if (courseStartDate <= currentDate) {
      alert("El curso ya comenzó");
      return;
    }

    // Realizar la inscripción
    try {
      const response = await fetch(`http://localhost:8083/insertinscripcion`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          id_course: courseId,
          id_user: user_id
        }),
      });

      if (response.ok) {
        alert("Inscripción exitosa! :)");
        window.location.reload();  // Alternativa: utiliza un estado o redirección sin recarga
      } else if (response.status === 500) {
        alert("Ya estás inscrito en este curso");
      } else {
        alert("Error en la inscripción. Inténtalo de nuevo.");
      }
    } catch (error) {
      console.log('Error al realizar la solicitud al backend:', error);
      alert("Error al realizar la inscripción. Inténtalo de nuevo más tarde.");
    }
  };

  // Si el usuario es admin o no está registrado, no mostrar el botón
  if (!user_id || isAdmin) {
    return null;
  }

  return (
      <button className="subscribe-button" onClick={handleSubmit}>INSCRIBIRME</button>
  );
}

export default Inscribirmebutton;
