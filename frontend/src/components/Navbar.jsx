import BurgerMenu from './BurgerMenu'
import Cookies from 'js-cookie';
import {useEffect, useState} from "react";
import '../estilos/Navbar.css';

export const Navbar = () => {
  const [userId, setUserId] = useState(null); // Inicialmente null
  const [isAdmin, setIsAdmin] = useState(false);

  useEffect(() => {
    const storedUserId = Cookies.get('user_id');
    if (storedUserId) {
      setUserId(parseInt(storedUserId, 10));
    }

    const storedAdmin = Cookies.get('admin');
    if (storedAdmin) {
      setIsAdmin(storedAdmin === "1"); // Convert to boolean
    }
  }, []);

  const handleLogout = () => {
    setUserId(null); // Actualizar el estado cuando se cierre la sesión
    window.location.reload(); //Recargo la pagina porque sino me sigue dejando inscribirme a cursos
  };

  const validUserId = typeof userId === 'number' ? userId : parseInt(userId, 10);

  return (
      <header>
        <div className='contenedor'>
          <div className='logoContainer'>
            <img src="/icon.png" alt="cafe y libros" width="100px" height="40px"/>
          </div>
          <div className='titleContainer'>
            Coffee&Learn
          </div>
          <div id='welcomeMessage'>
            <h1 style={{fontFamily: 'Spoof Trial, sans-serif'}}>
              {validUserId ? (
                  isAdmin ? 'Bienvenido administrador' : 'Bienvenido alumno'
              ) : (
                  'Inicie sesión'
              )}
            </h1>
          </div>
          <div className='menu'>
            <BurgerMenu onLogout={handleLogout}/>
          </div>
        </div>
      </header>
  );
}