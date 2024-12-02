import { useState } from "react";
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
    const [containers, setContainers] = useState([]); // Lista de contenedores
    const [loadingContainers, setLoadingContainers] = useState(false); // Estado de carga
    const [errorContainers, setErrorContainers] = useState(null); // Estado de error

    const token = Cookies.get("token");

    // Fetch de los datos de contenedores
    const fetchContainers = async () => {
        setLoadingContainers(true);
        setErrorContainers(null); // Limpiar errores previos

        if (!token) {
            setErrorContainers("Token no encontrado. Por favor inicia sesión.");
            setLoadingContainers(false);
            return;
        }

        try {
            const response = await fetch("http://host.docker.internal:8004/services", {
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
            setContainers(containerData.services || []); // Aseguramos la estructura
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
                    <VStack spacing={4}>
                        <Button
                            w="100%"
                            onClick={fetchContainers}
                            isLoading={loadingContainers}
                            style={{ fontFamily: "Spoof Trial, sans-serif" }}
                        >
                            Ver contenedores
                        </Button>
                        {errorContainers && (
                            <Box w="100%" color="red.500">
                                {errorContainers}
                            </Box>
                        )}
                        {!loadingContainers && containers.length > 0 && (
                            <Box w="100%" className="containerGrid">
                                {containers.map((container, index) => (
                                    <Box key={index} className="containerCard">
                                        <p>
                                            <strong>Nombre:</strong>{" "}
                                            {container.name.join(", ")}
                                        </p>
                                        <p>
                                            <strong>ID:</strong>{" "}
                                            {container.containers.join(", ")}
                                        </p>
                                    </Box>
                                ))}
                            </Box>
                        )}
                    </VStack>
                </DrawerBody>
            </DrawerContent>
        </Drawer>
    );
};

export default DashAdminPopup;
