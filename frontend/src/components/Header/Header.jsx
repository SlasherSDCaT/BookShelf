import classes from './Header.module.css'
import { FaUserAlt } from "react-icons/fa";
import { Link } from 'react-router-dom';
import Cookie from 'universal-cookie'

const Header = () => {
  
  const cookie = new Cookie()

  return (
    <header className={classes.header}>
        <Link to='/'><h2>BookShelf</h2></Link>
        <div className={classes.acc}>
            <span>
                <FaUserAlt/>
            </span>
            <span>
              <Link to='/myshelf'>Моя полка</Link>
              <span onClick={() => {
                cookie.remove("token")
                window.location.reload (); 
              }}>Выйти</span>
            </span>
        </div>
    </header>
  )
}

export default Header