import PropTypes from 'prop-types';
import { Input, InputGroup, InputLeftElement, Box, Button } from '@chakra-ui/react';
import { SearchIcon } from '@chakra-ui/icons';
import React from "react";
import '../estilos/SearchBar.css'

const SearchBar = ({ onSearchResults }) => {
    const [searchTerm, setSearchTerm] = React.useState('');

    const handleSearch = async (e) => {
        e.preventDefault();
        const baseUrl = 'http://localhost:8082/search';

        // Parámetros comunes para ambas solicitudes
        const params = new URLSearchParams();
        params.append('limit', '20');
        params.append('offset', '1');

        // Si el campo de búsqueda no está vacío, agrega el parámetro `q`
        if (searchTerm.trim() !== '') {
            params.append('q', searchTerm);
        }

        try {
            const response = await fetch(`${baseUrl}?${params.toString()}`, {
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