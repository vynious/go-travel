import React, { useState, useEffect } from "react";
import { getAuth, onAuthStateChanged } from "firebase/auth";
import { Link, navigate } from "gatsby";
import firebaseApp from "../../firebase/firebaseConfig";
import ASidebar from "../components/ASidebar";

const Home = () => {
    const [user, setUser] = useState(null);

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
                        console.log(data);
                        setUser(data.User);
                    })
                    .catch(error => {
                        console.error('Error fetching user data:', error);
                        setUser(null);
                    });
            } else {
                // User is signed out
                setUser(null);
                navigate("/login")
            }
        });
    }, []);


    return (
        <div>
            <ASidebar user={user} />
        </div>
    )

}

export default Home;