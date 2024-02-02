import React, { createContext, useContext, useState, useEffect } from 'react';
import { getAuth, onAuthStateChanged, signInWithEmailAndPassword } from 'firebase/auth';
import firebaseApp from '../firebase/firebaseConfig';

const UserContext = createContext();

export const useUser = () => useContext(UserContext);

export const UserProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const auth = getAuth(firebaseApp);
        const unsubscribe = onAuthStateChanged(auth, (firebaseUser) => {
            if (firebaseUser) {
                firebaseUser.getIdToken().then((token) => {
                    fetch(`${process.env.GATSBY_BACKEND_URL}/users/${firebaseUser.uid}`, {
                        method: 'GET',
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': `Bearer ${token}`,
                        },
                    })
                        .then(response => response.json())
                        .then(data => {
                            setUser(data.User); // if backend response has a User object
                            setIsLoading(false);
                        })
                        .catch(error => {
                            console.error('Error fetching user details:', error);
                            setIsLoading(false);
                        });
                });
            } else {
                setUser(null);
                setIsLoading(false);
            }
        });

        // Cleanup subscription on unmount
        return () => unsubscribe();
    }, []);

    const signIn = async (email, password) => {
        const auth = getAuth(firebaseApp);
        try {
            await signInWithEmailAndPassword(auth, email, password);
            return user;
        } catch (error) {
            throw error; // calling component handle the error
        }
    };

    const signOut = () => {
        const auth = getAuth(firebaseApp);
        auth.signOut().then(() => setUser(null));
    };

    // Check the current route and allow unrestricted access to certain pages
    const isUnrestrictedRoute = () => {
        const currentPath = window.location.pathname;
        // Define an array of unrestricted routes
        const unrestrictedRoutes = ['/'];

        return unrestrictedRoutes.includes(currentPath);
    };

    if (!isLoading || isUnrestrictedRoute()) {
        return (
            <UserContext.Provider value={{ user, signIn, setUser, signOut, isLoading }}>
                {children}
            </UserContext.Provider>
        );
    } else {
        return <div>Loading...</div>;
    }
};
