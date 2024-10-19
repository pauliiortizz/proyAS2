import React, { useEffect, useState } from 'react';

function SeeReviews({ isOpen, courseId }) {
    const [comments, setComments] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        if (!isOpen) return; // Evitar solicitudes innecesarias si el modal no está abierto

        const fetchComments = async () => {
            try {
                const response = await fetch(`http://localhost:8080/comentarios/${courseId}`);
                if (!response.ok) {
                    throw new Error('Error fetching comments');
                }
                const data = await response.json();
                setComments(Array.isArray(data) ? data : []); // Asegurar que data es un array
            } catch (error) {
                setError(error.message);
                console.error('Error fetching comments:', error);
            }
        };

        fetchComments();
    }, [isOpen, courseId]);

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div>
            <h2>Comentarios del Curso</h2>
            {comments.length === 0 ? (
                <p>No hay comentarios para este curso.</p>
            ) : (
                <ul>
                    {comments.map(comment => (
                        <li key={comment.id_inscripcion}>
                            <p>{comment.comentario}</p>
                            <p>Valoración: {comment.valoracion}</p>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}

export default SeeReviews;
