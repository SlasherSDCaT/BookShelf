import classes from './Footer.module.css'

const Footer = () => {

  return (
      <div className={classes.cont}>
        <div className={classes.policy}>
          &copy; 2024, MyBookshelf - виртуальная книжная полка made by Мухаметшин А.Р. from РТУ МИРЭА
        </div>
      </div>
  )
}

export default Footer