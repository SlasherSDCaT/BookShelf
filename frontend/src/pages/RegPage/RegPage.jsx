import React, { useState } from 'react'
import classes from './RegPage.module.css'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios'
import { BiErrorCircle } from "react-icons/bi"
import Cookies from 'universal-cookie'

const RegPage = () => {
    const navigate = useNavigate()
    const cookie = new Cookies()
    
    const [inf, setInf] = useState({
        email: "",
        password: "",
        conf: ""
    })

    const [error, setError] = useState({
        visibility: false,
        message: ""
    })

    const action = async (event) => {
        event.preventDefault()
        if (inf.password === inf.conf && inf.password.length > 7) {
            try {
                const data = await axios.post("https://bookshelf-ekzd.onrender.com/auth/sing-up", { // http://0.0.0.0:3001
                    Username: inf.email,
                    Password: inf.password,
                    Role: "USER"
                })
                cookie.set("token", data.data.Token)
                cookie.set("user_id", data.data.UserID)
                navigate("/")
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    setInf({...inf, password: "", conf: ""})
                    setError({visibility: true, message: "Пользователь уже зарегистрирован"})
                    setTimeout(() => setError({message: "", visibility: false}), 5000)
                }
            }
        } else {
            setInf({...inf, password: "", conf: ""})
            if (inf.password.length < 8) setError({visibility: true, message: "Длина пароля менее 8 символов"})
            else if (inf.password !== inf.conf) setError({visibility: true, message: "Пароли не совпадают"})
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
                    <h1>Регистрация</h1>
                    <form>
                        <label htmlFor="email">E-mail:</label>
                        <input type="email" name='email' required value={inf.email}
                            onChange={(event) => setInf({...inf, email: event.target.value})}/>
                        <label htmlFor="pass">Пароль:</label>
                        <input type="password" name='pass' required value={inf.password}
                            onChange={(event) => setInf({...inf, password: event.target.value})}/>
                        <label htmlFor="pass_confirm">Повторите пароль:</label>
                        <input value={inf.conf} type="password" name='pass_confirm' required
                            onChange={(event) => setInf({...inf, conf: event.target.value})}/>
                        <button onClick={action}>Зарегистрироваться</button>
                    </form>
                </div>
                <div className={classes.navs}>
                    <Link to='/'>&laquo; Вернуться на главную</Link>
                    <Link to='/login'>Уже есть аккаунт?&raquo;</Link>
                </div>
            </main>
        </>
    )
}

export default RegPage
