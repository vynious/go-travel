import React, { useState, useEffect } from "react";
import NavBar from "../components/Navbar";
import { useUser } from '../../context/UserContext';


export default function Connections() {
    const { user } = useUser()
    const [data, setData] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {

        const fetchData = async () => {
            try {
                const response = await fetch(`${process.env.GATSBY_BACKEND_URL}/connection/${user.id}`);
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                setData(data);
                console.log(data.Connections);
                setIsLoading(false);
            } catch (error) {
                setError(error.message);
                setIsLoading(false);
            }
        };

        fetchData();
    }, []); // Empty dependency array means this effect runs once on mount

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <section>
            <NavBar />
            <div style={{ display: 'flex', flexWrap: 'wrap', gap: '20px', justifyContent: 'center', padding: '20px' }}>
                {data.Users.map((user) => (
                    <div key={user.id} style={{ border: '1px solid #ddd', borderRadius: '10px', padding: '20px', width: '300px', textAlign: 'center' }}>
                        <div style={{ marginBottom: '10px' }}>
                            {/* Placeholder for profile picture */}
                            <img src={user.profile_picture.Valid ? user.profile_picture.String : "https://via.placeholder.com/150"} alt="profile" style={{ width: '100px', height: '100px', borderRadius: '50%' }} />
                        </div>
                        <div>
                            <h3>{user.name}</h3>
                            <p>@{user.username}</p>
                            <p>{user.email}</p>
                            <p>Joined: {new Date(user.creation_date).toLocaleDateString()}</p>
                        </div>
                    </div>
                ))}
            </div>
        </section>
    );
}