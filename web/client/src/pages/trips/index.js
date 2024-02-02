import React, { useState, useEffect } from 'react';
import { useUser } from '../../context/UserContext';
import ASidebar from '../components/ASidebar';
import UserInviteModal from './UserInviteModal';

const Trips = () => {
    const { user } = useUser();
    const [trips, setTrips] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);

    useEffect(() => {
        // Assuming you have a function to fetch trips
        fetchTrips();
    }, []);

    const fetchTrips = async () => {
        try {
            const response = await fetch(`${process.env.GATSBY_BACKEND_URL}/trips`, {
                headers: {
                    Authorization: `Bearer ${user.token}`, // Make sure to send the correct auth token
                },
            });
            const data = await response.json();

            setTrips(data.Trips || []);
        } catch (error) {
            console.error('Error fetching trips:', error);
        }
    };

    return (
        <div className='flex'>
            <ASidebar user={user} />
            <section className="flex-grow bg-white dark:bg-gray-900 p-4">
                <button onClick={() => setIsModalOpen(true)} className="px-6 py-2 text-white bg-blue-600 rounded-md">
                    Add User To Trip
                </button>
                {trips.map((trip) => (
                    <div key={trip.id} className="p-8 my-5 space-y-3 border-2 border-blue-400 dark:border-blue-300 rounded-xl">
                        
                        <h1 className="mt-6 text-xl font-semibold text-gray-800 dark:text-white">
                            {trip.title}
                        </h1>
                        <hr className="w-32 my-6 text-blue-500" />
                        <p className="text-sm text-gray-500 dark:text-gray-400">
                            {trip.country}
                        </p>
                        <a href="#" className="inline-block mt-4 text-blue-500 underline hover:text-blue-400">Read more</a>
                    </div>
                ))}
                <UserInviteModal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} onAddUser={() => {/* Add user to trip logic here */ }} />
            </section>
        </div>
    );
};

export default Trips;
