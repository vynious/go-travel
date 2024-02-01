import React, { useState } from 'react';
import firebaseApp from '../../firebase/firebaseConfig';
import {
    loginContainer,
    loginForm,
    formTitle,
    errorMessage,
    formInput,
    formButton
} from './Login.module.css';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    const handleLogin = async (e) => {
        e.preventDefault();
        setError('');
        try {
            await firebaseApp.auth().signInWithEmailAndPassword(email, password);
            // todo:  Handle successful login here


        } catch (error) {
            setError(error.message);
        }
    };

    return (
        <div className={loginContainer}>
            <form onSubmit={handleLogin} className={loginForm}>
                <h1 className={formTitle}>Login</h1>
                {error && <p className={errorMessage}>{error}</p>}
                <input
                    className={formInput}
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    className={formInput}
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button type="submit" className={formButton}>Login</button>
            </form>
        </div>
    );
};

export default Login;
