import React, { useEffect } from 'react';
import { navigate } from 'gatsby';
import { getAuth, onAuthStateChanged } from 'firebase/auth';
import firebaseApp from '../../firebase/firebaseConfig';
import { useUser } from '../../context/UserContext';
import ASidebar from '../components/ASidebar'; 

const Home = () => {
    const { user, setUser } = useUser();

    useEffect(() => {
        const auth = getAuth(firebaseApp);
        onAuthStateChanged(auth, async (firebaseUser) => {
            if (firebaseUser) {
                const token = await firebaseUser.getIdToken();
                fetch(`${process.env.GATSBY_BACKEND_URL}/users/${firebaseUser.uid}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                })
                    .then(response => {
                        if (!response.ok) {
                            throw new Error('Network response was not ok');
                        }
                        return response.json();
                    })
                    .then(data => {
                        console.log(data.User);
                        setUser(data.User);
                    })
                    .catch(error => {
                        console.error('Error fetching user data:', error);
                        setUser(null);
                    });
            } else {
                setUser(null);
                navigate("/login");
            }
        });
    }, [setUser]); 
    return (
        <ASidebar user={user} />
    );

    
};

export default Home;
