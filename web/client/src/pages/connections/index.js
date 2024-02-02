import React, { useState, useEffect } from "react";
import ASidebar from "../components/ASidebar";
import { useUser } from "../../context/UserContext";
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
                // fetch the list of connections
                const response = await fetch(`${process.env.GATSBY_BACKEND_URL}/connection/${user.id}`, {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('userToken')}`,
                    },
                });
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const connectionsData = await response.json();

                // for each connection, fetch the user details
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
    }, [user]); // re-fetch when user changes

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="flex">
            <ASidebar user={user} />
            <section class="container px-4 mx-auto">
                <div class="sm:flex sm:items-center sm:justify-between">
                    <div>
                        <div class="flex items-center gap-x-3">
                            <h2 class="text-lg font-medium text-gray-800 dark:text-white">Connections</h2>

                            <span class="px-3 py-1 text-xs text-blue-600 bg-blue-100 rounded-full dark:bg-gray-800 dark:text-blue-400">{connections.length} connections</span>
                        </div>

                        <p class="mt-1 text-sm text-gray-500 dark:text-gray-300">These are your connections.</p>
                    </div>

                    <div class="flex items-center mt-4 gap-x-3">


                        <button class="flex items-center justify-center w-1/2 px-5 py-2 text-sm tracking-wide text-white transition-colors duration-200 bg-blue-500 rounded-lg shrink-0 sm:w-auto gap-x-2 hover:bg-blue-600 dark:hover:bg-blue-500 dark:bg-blue-600">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v6m3-3H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>

                            <span>Add Connection</span>
                        </button>
                    </div>
                </div>

                <div class="mt-6 md:flex md:items-center md:justify-between">
                    <div class="inline-flex overflow-hidden bg-white border divide-x rounded-lg dark:bg-gray-900 rtl:flex-row-reverse dark:border-gray-700 dark:divide-gray-700">
                        <button class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 bg-gray-100 sm:text-sm dark:bg-gray-800 dark:text-gray-300">
                            Connected
                        </button>

                        <button class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 sm:text-sm dark:hover:bg-gray-800 dark:text-gray-300 hover:bg-gray-100">
                            Suggested
                        </button>

                        <button class="px-5 py-2 text-xs font-medium text-gray-600 transition-colors duration-200 sm:text-sm dark:hover:bg-gray-800 dark:text-gray-300 hover:bg-gray-100">
                            Others
                        </button>
                    </div>

                    <div class="relative flex items-center mt-4 md:mt-0">
                        <span class="absolute">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 mx-3 text-gray-400 dark:text-gray-600">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
                            </svg>
                        </span>

                        <input type="text" value="Search" placeholder="Search" class="block w-full py-1.5 pr-5 text-gray-700 bg-white border border-gray-200 rounded-lg md:w-80 placeholder-gray-400/70 pl-11 rtl:pr-11 rtl:pl-5 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 dark:focus:border-blue-300 focus:ring-blue-300 focus:outline-none focus:ring focus:ring-opacity-40" />
                    </div>
                </div>

                <div class="flex flex-col mt-6">
                    <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                        <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
                            <div class="overflow-hidden border border-gray-200 dark:border-gray-700 md:rounded-lg">
                                <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                                    <thead class="bg-gray-50 dark:bg-gray-800">
                                        <tr>
                                            <th scope="col" class="py-3.5 px-4 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400">Name</th>
                                            <th scope="col" class="px-12 py-3.5 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400">Username</th>
                                            <th scope="col" class="px-4 py-3.5 text-sm font-normal text-left rtl:text-right text-gray-500 dark:text-gray-400">Email</th>
                                        </tr>
                                    </thead>
                                    <tbody class="bg-white divide-y divide-gray-200 dark:divide-gray-700 dark:bg-gray-900">
                                        {connections.map((connection) => (
                                            <tr key={connection.User.id}>
                                                <td class="px-4 py-4 text-sm font-medium whitespace-nowrap">
                                                    <div>
                                                        <h2 class="font-medium text-gray-800 dark:text-white">{connection.User.name}</h2>
                                                    </div>
                                                </td>
                                                <td class="px-4 py-4 text-sm whitespace-nowrap">
                                                    <div>
                                                        <h2 class="font-medium text-gray-800 dark:text-white">{connection.User.username}</h2>
                                                    </div>
                                                </td>
                                                <td class="px-4 py-4 text-sm whitespace-nowrap">
                                                    <div>
                                                        <h2 class="font-medium text-gray-800 dark:text-white">{connection.User.email}</h2>
                                                    </div>
                                                </td>

                                            </tr>
                                        ))}
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="mt-6 sm:flex sm:items-center sm:justify-between ">
                    <div class="text-sm text-gray-500 dark:text-gray-400">
                        Page <span class="font-medium text-gray-700 dark:text-gray-100">1 of 10</span>
                    </div>

                    <div class="flex items-center mt-4 gap-x-4 sm:mt-0">
                        <a href="#" class="flex items-center justify-center w-1/2 px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md sm:w-auto gap-x-2 hover:bg-gray-100 dark:bg-gray-900 dark:text-gray-200 dark:border-gray-700 dark:hover:bg-gray-800">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 rtl:-scale-x-100">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18" />
                            </svg>

                            <span>
                                previous
                            </span>
                        </a>

                        <a href="#" class="flex items-center justify-center w-1/2 px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md sm:w-auto gap-x-2 hover:bg-gray-100 dark:bg-gray-900 dark:text-gray-200 dark:border-gray-700 dark:hover:bg-gray-800">
                            <span>
                                Next
                            </span>

                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 rtl:-scale-x-100">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3" />
                            </svg>
                        </a>
                    </div>
                </div>
            </section>

        </div>
    );
}
