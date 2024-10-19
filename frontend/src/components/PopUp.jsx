import React from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton
} from "@chakra-ui/react";
import Login from './Login';

const Popup = ({ isOpen, onClose }) => {
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader style={{fontFamily: 'Spoof Trial, sans-serif'}}>Inicio de sesión</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <Login onClose={onClose} />
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};

export default Popup;