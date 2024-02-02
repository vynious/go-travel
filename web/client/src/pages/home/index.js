import React, { useEffect } from 'react';
import { navigate } from 'gatsby';
import { getAuth, onAuthStateChanged } from 'firebase/auth';
import firebaseApp from '../../firebase/firebaseConfig';
import { useUser } from '../../context/UserContext';
import ASidebar from '../components/ASidebar'; // Ensure this path is correct

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
                        setUser(data.User); // Make sure the data structure matches your context expectation
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
    }, [setUser]); // Adding setUser as a dependency is good practice to ensure useEffect has access to the latest setUser function

    return (
        <ASidebar user={user} />
    );

    // Now `user` is correctly passed to ASidebar
};

export default Home;
