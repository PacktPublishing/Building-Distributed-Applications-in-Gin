import React from 'react';
import { useAuth0 } from "@auth0/auth0-react";
import './Profile.css'

const Profile = () => {
    const { user, logout } = useAuth0();
    return (
        <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                <div class="user">
                    <img src={user.picture} class="rounded-circle" />
                    <span>{user.name}</span>
                </div>
            </a>
            <div class="dropdown-menu" aria-labelledby="navbarDropdown">
                <a class="dropdown-item" onClick={() => logout()}> Logout</a>
            </div>
        </li>
    )
}

export default Profile;