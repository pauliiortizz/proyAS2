import React from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalCloseButton
} from "@chakra-ui/react";
import SeeReviews from "./SeeReviews.jsx";

const PopupSeeReview = ({ isOpen, onClose, courseId }) => {
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader>Reviews</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <SeeReviews isOpen={isOpen} courseId={courseId} />
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};

export default PopupSeeReview;