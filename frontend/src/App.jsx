import * as React from 'react';
import { ChakraProvider } from '@chakra-ui/react';
import { Navbar } from './components/Navbar';
import { ItemListContainer } from './components/ItemListContainer';
import './App.css';

function App() {
    return (
        <ChakraProvider>
            <header>
                <Navbar />
            </header>
            <ItemListContainer />
        </ChakraProvider>
    );
}

export default App;
