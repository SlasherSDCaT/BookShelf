import React, { useState } from 'react'
import classes from './LogPage.module.css'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios'
import { BiErrorCircle } from 'react-icons/bi'
import Cookies from 'universal-cookie'

const LogPage = () => {
    const navigate = useNavigate()
    const cookie = new Cookies()
    
    const [inf, setInf] = useState({
        email: "",
        password: ""
    })

    const [error, setError] = useState({
        visibility: false,
        message: ""
    })

    const action = async (event) => {
        event.preventDefault()
        if (inf.password.length > 7) {
            try {
                const data = await axios.post("http://0.0.0.0:3001/auth/sing-in", {
                    Username: inf.email,
                    Password: inf.password
                })
                cookie.set("token", data.data.Token)
                cookie.set("user_id", data.data.UserID)
                navigate("/")
            } catch (e) {
                setInf({...inf, password: ""})
                console.log(e.message)
                if (e.response && e.response.status === 400) {
                    setError({visibility: true, message: "Неверные учетные данные"})
                    setTimeout(() => setError({message: "", visibility: false}), 5000)
                }
            }
        } else {
            setInf({...inf, password: ""})
            if (inf.password.length < 8) setError({visibility: true, message: "Длина пароля менее 8 символов"})
            setTimeout(() => setError({message: "", visibility: false}), 5000)
        }
    }

    return (
        <>
            <main className={classes.main}>
                <div style={{top: error.visibility ? "20px" : "-100%"}} className={classes.err_mes}>
                    <span>
                        <BiErrorCircle/>
                    </span>
                    {error.message}
                </div>
                <div className={classes.reg}>
                    <h1>Вход</h1>
                    <form>
                        <label htmlFor="email">E-mail:</label>
                        <input type="email" name='email' required value={inf.email}
                            onChange={(event) => setInf({...inf, email: event.target.value})}/>
                        <label htmlFor="pass">Пароль:</label>
                        <input type="password" name='pass' required value={inf.password}
                            onChange={(event) => setInf({...inf, password: event.target.value})}/>
                        <button onClick={action}>Войти</button>
                    </form>
                </div>
                <div className={classes.navs}>
                    <Link to='/'>&laquo; Вернуться на главную</Link>
                    <Link to='/register'>Желаете зарегистрироваться?&raquo;</Link>
                </div>
            </main>
        </>
    )
}

export default LogPage
