import React, { useState } from 'react';
import firebaseApp from '../../firebase/firebaseConfig';
import { getAuth, signInWithEmailAndPassword } from 'firebase/auth';
import {Link, navigate} from 'gatsby'

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    const handleLogin = async (e) => {
        e.preventDefault();
        setError('');
        try {
            const auth = getAuth(firebaseApp);
            await signInWithEmailAndPassword(auth, email, password);
            // todo: Handle successful login here
            // Redirect or set token in header as needed
            navigate('/home');
        } catch (error) {
            console.log(error.message);
            setError(error.message);
        }
    };

    return (
        <div className="bg-white dark:bg-gray-900 min-h-screen flex justify-center items-center">
            <div className="flex items-center w-full max-w-md px-6 mx-auto">
                <div className="flex-1">
                    <div className="text-center">
                        <h2 className="text-2xl font-bold text-gray-900 dark:text-white">Login to goTravel</h2>
                        <p className="mt-3 text-gray-500 dark:text-gray-300">Sign in to access your account</p>
                    </div>

                    <div className="mt-8">
                        <form onSubmit={handleLogin}>
                            {error && <p className="text-sm text-red-500">{error}</p>}
                            <div>
                                <label htmlFor="email" className="block mb-2 text-sm text-gray-600 dark:text-gray-200">Email Address</label>
                                <input type="email" name="email" id="email" placeholder="example@example.com"
                                    className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-lg dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                />
                            </div>

                            <div className="mt-6">
                                <label htmlFor="password" className="block mb-2 text-sm text-gray-600 dark:text-gray-200">Password</label>
                                <input type="password" name="password" id="password" placeholder="Your Password"
                                    className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-lg dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                />
                            </div>

                            <div className="mt-6">
                                <button type="submit"
                                    className="w-full px-4 py-2 text-white bg-blue-500 rounded-lg hover:bg-blue-700 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-800"
                                >
                                    Sign in
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Login;
