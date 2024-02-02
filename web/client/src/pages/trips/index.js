import React, { useState, useEffect } from 'react';
import NavBar from '../components/Navbar';
import Footer from '../components/Footer';


const Trips = () => {
    const [trips, setTrips] = useState([]);
    const backendUrl = process.env.GATSBY_BACKEND_URL;

    useEffect(() => {
        const fetchTrips = async () => {
            try {
                const response = await fetch(`${backendUrl}/trips`);
                const data = await response.json();
                setTrips(data); // Set trip data to state
            } catch (error) {
                console.error("Error fetching data: ", error);
                // Handle error scenario
            }
        };

        fetchTrips();
    }, []);

    return (
        <section className="bg-white dark:bg-gray-900">
            <NavBar />
            <div className="container px-6 py-10 mx-auto">
                <div className="text-center">
                    <h1 className="text-2xl font-semibold text-gray-800 capitalize lg:text-3xl dark:text-white">All my trips</h1>
                    <p className="max-w-lg mx-auto mt-4 text-gray-500">
                        A record of all the trips your trips
                    </p>
                </div>

                <div className="grid grid-cols-1 gap-8 mt-8 md:mt-16 md:grid-cols-2 xl:grid-cols-3">
                    {trips.map((trip) => (
                        <div key={trip.id} className="p-8 space-y-3 border-2 border-blue-400 dark:border-blue-300 rounded-xl">
                            <div className="relative">
                                <img className="object-cover object-center w-full h-64 rounded-lg lg:h-80" src={trip.imageUrl} alt={trip.title} />
                                <div className="absolute bottom-0 flex p-3 bg-white dark:bg-gray-900">
                                    <img className="object-cover object-center w-10 h-10 rounded-full" src={trip.author.imageUrl} alt={trip.author.name} />
                                    <div className="mx-4">
                                        <h1 className="text-sm text-gray-700 dark:text-gray-200">{trip.author.name}</h1>
                                        <p className="text-sm text-gray-500 dark:text-gray-400">{trip.author.role}</p>
                                    </div>
                                </div>
                            </div>
                            <h1 className="mt-6 text-xl font-semibold text-gray-800 dark:text-white">
                                {trip.title}
                            </h1>
                            <hr className="w-32 my-6 text-blue-500" />
                            <p className="text-sm text-gray-500 dark:text-gray-400">
                                {trip.description}
                            </p>
                            <a href="#" className="inline-block mt-4 text-blue-500 underline hover:text-blue-400">Read more</a>
                        </div>
                    ))}
                </div>
            </div>
            <Footer />
        </section>
    );
};

export default Trips;