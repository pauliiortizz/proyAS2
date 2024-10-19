import React from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton
} from "@chakra-ui/react";
import EditCourse from './EditCourse.jsx';

const PopupEdit = ({ isOpen, onClose, courseId }) => {
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader style={{fontFamily: 'Spoof Trial, sans-serif'}}>Editar curso</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <EditCourse courseId={courseId} onClose={onClose} />
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};

export default PopupEdit;
