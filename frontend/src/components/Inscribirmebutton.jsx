import '../estilos/Inscribirmebutton.css';
import { useState, useEffect } from "react";
import Cookies from "js-cookie";

const Inscribirmebutton = ({ courseId }) => {
  const [userId, setUserId] = useState(null);
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

    if (!userId || userId === -1 || userId === 0 || !tokenUser) {
      alert("Debes registrarte para inscribirte a un curso");
    } else {
      const currentDate = new Date();
      const storedDate = Cookies.get('Fecha_inicio');

      if (storedDate) {
        const [storedYear, storedMonth, storedDay] = storedDate.split('-').map(Number);
        const storedDateObject = new Date(storedYear, storedMonth - 1, storedDay);

        if (storedDateObject <= currentDate) {
          alert("El curso ya comenzó");
          return;
        }
      } else {
        console.log('No se encontró la cookie de Fecha_inicio');
      }

      try {
        const response = await fetch(`http://localhost:8080/inscripcion`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            fecha_inscripcion: currentDate,
            id_course: courseId,
            id_user: userId
          }),
        });

        if (response.ok) {
          alert("Inscripción exitosa! :)");
          window.location.reload()
        } else if (response.status === 500) {
          alert("Ya estás inscrito en este curso");
        }
      } catch (error) {
        console.log('Error al realizar la solicitud al backend:', error);
        alert("Error al realizar la inscripción. Inténtalo de nuevo más tarde.");
      }
    }
  };

  if (!userId || isAdmin) {
    return null;
  }

  return (
      <button className="subscribe-button" onClick={handleSubmit}>INSCRIBIRME</button>
  );
}

export default Inscribirmebutton;
