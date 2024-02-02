
import React, { useEffect } from 'react';
import { navigate } from 'gatsby';
import { useUser } from '../../context/UserContext';

const withAuthentication = (WrappedComponent) => {
    const WithAuthComponent = (props) => {
        const { user, isLoading } = useUser();

        useEffect(() => {
            if (!isLoading && !user) {
                navigate('/login');
            }
        }, [user, isLoading]);

        if (isLoading || !user) return <div>Loading...</div>; 
        
        return <WrappedComponent {...props} />;
    };

    return WithAuthComponent;
};

export default withAuthentication;
