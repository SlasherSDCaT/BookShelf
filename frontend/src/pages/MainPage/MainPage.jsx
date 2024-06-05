// pages/MainPage.js
import React from 'react';
import { Button, Container, Accordion } from 'react-bootstrap';
import Slider from '../../components/Slider';
import Collections from '../../components/Collections';
import Books from '../../components/Books';
import Footer from '../../components/Footer/Footer';
import Header from '../../components/Header/Header';
import { useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';

function MainPage() {
  const cookie = new Cookies()
  const navigate = useNavigate();

  useEffect(() => {
    const token = cookie.get('token');
    if (!token) {
      // return <Navigate to="/login" />; // You can use this if using React Router v6
      navigate("/login");
    }
  }, []);

  return (<>
  <Header/>
  <Container>
    <h1 className='hello'>BookShelf - все книги интернета в одном месте!</h1>
    <Slider />
    <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
          <Accordion.Header>Коллекции</Accordion.Header>
          <Accordion.Body>
          <h2>Коллекции</h2>
          <Collections />
          </Accordion.Body>
        </Accordion.Item>
      </Accordion>
      <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
          <Accordion.Header>Книги</Accordion.Header>
          <Accordion.Body>
          <h2 style={{display: 'inline' }}>Книги</h2>
          {/* <Button style={{ marginLeft: '63rem', marginBottom: '10px' }}>
            Найти по жанру
          </Button> */}
          <Books />
          </Accordion.Body>
        </Accordion.Item>
      </Accordion>
  </Container>
  <Footer/>
  </>
  
)};

export default MainPage;