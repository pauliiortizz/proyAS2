import { useState, useEffect } from "react";
import {
    Drawer,
    DrawerOverlay,
    DrawerContent,
    DrawerCloseButton,
    DrawerHeader,
    DrawerBody,
    Button,
    VStack,
    Box,
} from "@chakra-ui/react";
import Cookies from "js-cookie";

const DashAdminPopup = ({ isOpen, onClose }) => {
    const [userData, setUserData] = useState(null); // Datos del usuario
    const [containers, setContainers] = useState([]); // Lista de contenedores
    const [loadingUser, setLoadingUser] = useState(true); // Estado de carga del usuario
    const [loadingContainers, setLoadingContainers] = useState(false); // Estado de carga de los contenedores
    const [errorUser, setErrorUser] = useState(null); // Error al obtener usuario
    const [errorContainers, setErrorContainers] = useState(null); // Error al obtener contenedores

    const token = Cookies.get("token");
    const user_id = Cookies.get("user_id");

    // Fetch de los datos del usuario
    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await fetch(`http://localhost:8080/users/${user_id}`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                if (!response.ok) {
                    throw new Error("Failed to fetch admin data");
                }

                const userData = await response.json();
                setUserData(userData);

                if (!userData.Admin) {
                    setErrorUser("Acceso denegado. No eres administrador.");
                }
            } catch (error) {
                console.error("Error fetching user:", error);
                setErrorUser("Error al obtener datos del administrador.");
            } finally {
                setLoadingUser(false);
            }
        };

        fetchUser();
    }, [user_id, token]);

    // Fetch de los datos de contenedores
    const fetchContainers = async () => {
        setLoadingContainers(true);
        try {
            const response = await fetch("http://localhost:8004/admin/containers", {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || "Error fetching containers");
            }

            const containerData = await response.json();
            setContainers(containerData);
        } catch (error) {
            console.error("Error fetching containers:", error);
            setErrorContainers("Error al obtener la lista de contenedores.");
        } finally {
            setLoadingContainers(false);
        }
    };

    return (
        <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
            <DrawerOverlay />
            <DrawerContent>
                <DrawerCloseButton />
                <DrawerHeader>Tablero Administrativo</DrawerHeader>
                <DrawerBody>
                    {loadingUser ? (
                        <p>Cargando datos del administrador...</p>
                    ) : errorUser ? (
                        <p style={{ color: "red" }}>{errorUser}</p>
                    ) : (
                        <VStack spacing={4}>
                            <Box>
                                <p><strong>Nombre:</strong> {userData.name}</p>
                                <p><strong>Email:</strong> {userData.email}</p>
                                <p><strong>Rol:</strong> Administrador</p>
                            </Box>
                            <Button
                                w="100%"
                                onClick={fetchContainers}
                                isLoading={loadingContainers}
                                style={{ fontFamily: "Spoof Trial, sans-serif" }}
                            >
                                Ver contenedores
                            </Button>
                            {errorContainers && <p style={{ color: "red" }}>{errorContainers}</p>}
                            {!loadingContainers && containers.length > 0 && (
                                <Box w="100%" className="containerGrid">
                                    {containers.map((container) => (
                                        <Box key={container.ID} className="containerCard">
                                            <p><strong>ID:</strong> {container.ID}</p>
                                            <p><strong>Imagen:</strong> {container.Image}</p>
                                            <p><strong>Estado:</strong> {container.Status}</p>
                                        </Box>
                                    ))}
                                </Box>
                            )}
                        </VStack>
                    )}
                </DrawerBody>
            </DrawerContent>
        </Drawer>
    );
};

export default DashAdminPopup;