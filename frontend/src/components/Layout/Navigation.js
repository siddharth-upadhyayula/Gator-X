import React from 'react';
import {Link} from 'react-router-dom';
import './Navigation.css';
 
const Navigation = () => {
    return (
        <nav className="col-md-2">
        	<ul>
                <li><div className="sidebar_item" ><Link to={'/create'}>Add New Review</Link></div></li>
				<li><div className="sidebar_item" ><Link to={'/'}>List All Reviews</Link></div></li>
            </ul>
        </nav>
     );
}
 
export default Navigation;