import PropTypes from 'prop-types';
import { Input, InputGroup, InputLeftElement, Box, Button } from '@chakra-ui/react';
import { SearchIcon } from '@chakra-ui/icons';
import React from "react";
import '../estilos/SearchBar.css'

const SearchBar = ({ onSearchResults }) => {
    const [searchTerm, setSearchTerm] = React.useState('');

    const handleSearch = async (e) => {
        e.preventDefault();
        if (searchTerm.trim() === '') {
            // If the search term is empty, fetch all courses
            try {
                const response = await fetch(`http://localhost:8080/search`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });

                if (response.ok) {
                    const data = await response.json();
                    onSearchResults(data);
                } else {
                    alert("No se encontraron cursos.");
                    onSearchResults([]);
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
                alert("Error al buscar cursos. Inténtalo de nuevo más tarde.");
                onSearchResults([]);
            }
            return;
        }

        try {
            const response = await fetch(`http://localhost:8080/search/${searchTerm}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            if (response.ok) {
                const data = await response.json();
                onSearchResults(data);
            } else {
                alert("No se encontraron cursos con ese nombre.");
                onSearchResults([]);
            }
        } catch (error) {
            console.log('Error al realizar la solicitud al backend:', error);
            alert("Error al buscar cursos. Inténtalo de nuevo más tarde.");
            onSearchResults([]);
        }
    };

    return (
        <Box className='search' id='caja'>
            <form onSubmit={handleSearch}>
                <InputGroup>
                    <InputLeftElement pointerEvents="none">
                        <SearchIcon id='icono' />
                    </InputLeftElement>
                    <Input
                        className='input'
                        type="text"
                        placeholder="Buscar cursos por nombre..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                    />
                </InputGroup>
                {/*<Button type="submit" mt={2} width="100%">Buscar</Button>*/}
            </form>
        </Box>
    );
};

SearchBar.propTypes = {
    onSearchResults: PropTypes.func.isRequired,
};

export default SearchBar;