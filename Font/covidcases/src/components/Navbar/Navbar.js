import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { SidebarData } from "./SidebarData";
import { IconContext } from 'react-icons';
import * as FaIcons from 'react-icons/fa';
import * as AiIcons from 'react-icons/ai';
import './Navbar.css';

function Navbar(props) {
    const [sidebar, setSidebar] = useState(true);

    const showSidebar = () => setSidebar(!sidebar);
    return (
        <div>
            <IconContext.Provider value={{ color:'white' }}>
                {/* <div className='navbar'>
                    <Link to='#' className='menu-bars'>
                        <FaIcons.FaBars onClick={showSidebar} />
                    </Link>
                </div> */}
                <nav className={sidebar ? 'nav-menu active' : 'nav-menu'}>
                    <div className='nav-menu-items'>
                        {/* <div className='navbar-toggle'>
                            <Link to='#' className='menu-bars'>
                                <AiIcons.AiOutlineClose />
                            </Link>
                        </div> */}
                        {SidebarData.map((item, index) => {
                            return (
                                <div key={index} className={item.cName}>
                                    <Link to={item.path}>
                                        {item.icon}
                                    </Link>
                                </div>
                            )
                        })}
                    </div>
                </nav>
            </IconContext.Provider>
        </div>

    )
}

export default Navbar;