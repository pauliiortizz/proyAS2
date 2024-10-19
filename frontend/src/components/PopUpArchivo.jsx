import React from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton
} from "@chakra-ui/react";
import SubirArchivo from './SubirArchivo.jsx';

const PopupSubirArchivo = ({ isOpen, onClose, courseId }) => {
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader style={{fontFamily: 'Spoof Trial, sans-serif'}}>Subir Archivo</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <SubirArchivo courseId={courseId} onClose={onClose} />
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};

export default PopupSubirArchivo;