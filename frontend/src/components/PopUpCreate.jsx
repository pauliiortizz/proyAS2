import React from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton
} from "@chakra-ui/react";
import CreateCourse from './CreateCourse.jsx';

const PopupCreate = ({ isOpen, onClose }) => {
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader style={{fontFamily: 'Spoof Trial, sans-serif'}}>Crear Curso</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <CreateCourse onClose={onClose} />
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};

export default PopupCreate;