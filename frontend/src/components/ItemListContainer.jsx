import { useState, useEffect } from 'react';
import ItemList from './ItemList';
import SearchBar from './SearchBar';

export const ItemListContainer = () => {
    const [courses, setCourses] = useState([]);
    const [filteredCourses, setFilteredCourses] = useState([]);
    const [loading, setLoading] = useState(false);
    const [showMyCourses, setShowMyCourses] = useState(false);

    useEffect(() => {
        fetch(`http://localhost:8080/search`)
            .then(response => response.json())
            .then(data => {
                setCourses(data);
                setFilteredCourses(data);
            })
            .catch(error => console.error('Error fetching courses:', error));
    }, []);

    const handleSearchResults = (results) => {
        setFilteredCourses(results);
    };


    return (
        <div>
            <SearchBar onSearchResults={handleSearchResults} />
            {loading ? (
                <p>Cargando cursos...</p>
            ) : (
                <ItemList courses={filteredCourses} />
            )}
        </div>
    );
};