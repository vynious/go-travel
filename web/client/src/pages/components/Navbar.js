import React from 'react';
import { Link } from 'gatsby';
import {
    navBar,
    link,
} from './Navbar.module.css'; // Import the styles

export default function NavBar () {
    return (
        <nav className={navBar}>
            <Link to="/connections" className={link}>Connections</Link>
            <Link to="/trips" className={link}>Trips</Link>
            <Link to="/profile" className={link}>Profile</Link>
        </nav>
    );
};


