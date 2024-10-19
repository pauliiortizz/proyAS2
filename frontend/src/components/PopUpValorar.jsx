import React from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton
} from "@chakra-ui/react";
import ValorarCourse from "./Valorar.jsx";

const PopupValorar = ({ isOpen, onClose, courseId }) => {
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader style={{fontFamily: 'Spoof Trial, sans-serif'}}>Valorar</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <ValorarCourse courseId={courseId} onClose={onClose} />
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};

export default PopupValorar;
