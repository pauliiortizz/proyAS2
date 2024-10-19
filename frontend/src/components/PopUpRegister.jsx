import React from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton
} from "@chakra-ui/react";
import RegisterUser from './RegisterUser.jsx';

const PopupRegister = ({ isOpen, onClose }) => {
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader style={{fontFamily: 'Spoof Trial, sans-serif'}}>Registrarme</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <RegisterUser onClose={onClose} />
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};

export default PopupRegister;