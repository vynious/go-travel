import React from 'react';
import { UserProvider } from '../../../context/UserContext';

const AuthLayout = ({ children }) => {
    return (
        <UserProvider>
            {children}
        </UserProvider>
    );
};

export default AuthLayout;
