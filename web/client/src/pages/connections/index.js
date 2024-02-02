import React, { useState, useEffect } from "react";
import ASidebar from "../components/ASidebar"; // Adjust import path as needed
import { useUser } from "../../context/UserContext"; // Ensure the path is correct
import { navigate } from "gatsby";

export default function Connections() {
    const { user } = useUser();
    const [connections, setConnections] = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        if (!user) {
            navigate("/login"); // Redirect to login if no user found
            return;
        }

        const fetchConnections = async () => {
            setIsLoading(true);
            try {
                // First, fetch the list of connections
                const response = await fetch(`${process.env.GATSBY_BACKEND_URL}/connection/${user.id}`, {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('userToken')}`, 
                    },
                });
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const connectionsData = await response.json();

                // For each connection, fetch the user details
                const connectionDetailsPromises = connectionsData.Connections.map(async (connection) => {
                    const userResponse = await fetch(`${process.env.GATSBY_BACKEND_URL}/users/${connection.connected_user_id.String}`, {
                        headers: {
                            'Authorization': `Bearer ${localStorage.getItem('userToken')}`,
                        },
                    });
                    if (!userResponse.ok) {
                        throw new Error(`HTTP error! status: ${userResponse.status}`);
                    }
                    return userResponse.json();
                });

                const connectionsDetails = await Promise.all(connectionDetailsPromises);
                setConnections(connectionsDetails);
                console.log(connectionsDetails)
                setIsLoading(false);
            } catch (error) {
                setError(error.message);
                setIsLoading(false);
            }
        };

        fetchConnections();
    }, [user]); // Depend on user to re-fetch when user changes

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="flex">
            <ASidebar user={user} />
            <section className="flex-grow">
                <div style={{ display: 'flex', flexWrap: 'wrap', gap: '20px', justifyContent: 'center', padding: '20px' }}>
                    {connections.map((connection) => (
                        <div key={connection.User.id} style={{ border: '1px solid #ddd', borderRadius: '10px', padding: '20px', width: '300px', textAlign: 'center' }}>
                            <div style={{ marginBottom: '10px' }}>
                                {/* Check if profile_picture is valid and has a string; otherwise, use a placeholder */}
                                <img src={connection.User.profile_picture.Valid ? connection.User.profile_picture.String : "https://via.placeholder.com/150"} alt="profile" style={{ width: '100px', height: '100px', borderRadius: '50%' }} />
                            </div>
                            <div>
                                <h3>{connection.User.name}</h3>
                                <p>@{connection.User.username}</p>
                                <p>{connection.User.email}</p>
                                {/* Ensure creation_date exists before attempting to format it */}
                                <p>Joined: {connection.User.creation_date ? new Date(connection.User.creation_date).toLocaleDateString() : "Unknown"}</p>
                            </div>
                        </div>
                    ))}

                </div>
            </section>
        </div>
    );
}
