import React from 'react';
import { useAuth0 } from "@auth0/auth0-react";
import Profile from './Profile';

const Navbar = () => {
    const { isAuthenticated, loginWithRedirect, logout, user } = useAuth0();
    return (
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
            <a class="navbar-brand" href="#">Recipes</a>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarTogglerDemo02" aria-controls="navbarTogglerDemo02" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbarTogglerDemo02">
                <ul class="navbar-nav ml-auto">
                    <li class="nav-item">
                        {isAuthenticated ? (<Profile />) : (
                            <a class="nav-link active" onClick={() => loginWithRedirect()}> Login</a>
                        )}
                    </li>
                </ul>
            </div>
        </nav >
    )
}

export default Navbar;