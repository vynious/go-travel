import './src/styles/global.css'
import React from 'react';
import { UserProvider } from './src/context/UserContext'; // Ensure the path is correct

export const wrapRootElement = ({ element }) => {
    return <UserProvider>{element}</UserProvider>;
};
