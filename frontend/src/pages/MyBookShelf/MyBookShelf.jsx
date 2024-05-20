// pages/MainPage.js
import React, { useEffect, useState } from 'react';
import { Container, Button, Modal, Form, Accordion } from 'react-bootstrap';
import { useDispatch } from 'react-redux';
import CollectionsUser from '../../components/CollectionsUser';
import Books from '../../components/Books';
import Footer from '../../components/Footer/Footer';
import Header from '../../components/Header/Header';
import { addCollection } from '../../redux/slice/collectionsSlice';
import { addBook } from '../../redux/slice/booksSlice';
import Cookies from 'universal-cookie';

const MyBookShelf = () => {
    const cookie = new Cookies()
    useEffect(() => {
        const token = cookie.get('token');
        if (!token) {
          // return <Navigate to="/login" />; // You can use this if using React Router v6
          window.location.href = '/login';
        }
      }, []);

  const dispatch = useDispatch();

  const [showCollectionModal, setShowCollectionModal] = useState(false);
  const [showBookModal, setShowBookModal] = useState(false);

  const [collectionForm, setCollectionForm] = useState({
    user_id: cookie.get('user_id'),
    name: '',
    description: '',
    rating: 1
  });

  const [bookForm, setBookForm] = useState({
    user_id: cookie.get('user_id'),
    title: '',
    author: '',
    genre: '',
    description: '',
    image: '',
    body: ''
  });


  const handleCollectionChange = (e) => {
    setCollectionForm({ ...collectionForm, [e.target.name]: e.target.value });
  };

  const handleBookChange = (e) => {
    setBookForm({ ...bookForm, [e.target.name]: e.target.value });
  };

  const handleCollectionSubmit = async (e) => {
    e.preventDefault();
    dispatch(addCollection(collectionForm));
    window.location.reload();
  };

  const handleBookSubmit = async (e) => {
    e.preventDefault();
    dispatch(addBook(bookForm));
    window.location.reload();
  };

  return (
    <>
      <Header />
      <Container>
      <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
          <Accordion.Header>Коллекции</Accordion.Header>
          <Accordion.Body>
          <h2 style={{ display: 'inline' }}>Коллекции</h2>
        <div style={{ display: 'inline', marginLeft: '53rem'}}>
          <Button onClick={() => setShowCollectionModal(true)} style={{ marginBottom: '10px' }}>Добавить</Button>
        <Button variant='danger'  style={{ marginBottom: '10px' }}>Удалить</Button>
        </div>
        <CollectionsUser />
          </Accordion.Body>
        </Accordion.Item>
      </Accordion>
      <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
          <Accordion.Header>Книги</Accordion.Header>
          <Accordion.Body>
          <h2  style={{ display: 'inline' }}>Книги</h2>
          <div style={{ display: 'inline', marginLeft: '58rem'}}>
          <Button onClick={() => setShowBookModal(true)} style={{ marginBottom: '10px' }}>Добавить</Button>
        <Button variant='danger'  style={{ marginBottom: '10px' }}>Удалить</Button>
        </div>
        <Books />
          </Accordion.Body>
        </Accordion.Item>
      </Accordion>
      
      </Container>
      <Footer />

      {/* Add Collection Modal */}
      <Modal show={showCollectionModal} onHide={() => setShowCollectionModal(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Добавить коллекцию</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleCollectionSubmit}>
            <Form.Group controlId="collectionName">
              <Form.Label>Название</Form.Label>
              <Form.Control
                type="text"
                name="name"
                value={collectionForm.name}
                onChange={handleCollectionChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="collectionDescription">
              <Form.Label>Описание</Form.Label>
              <Form.Control
                type="text"
                name="description"
                value={collectionForm.description}
                onChange={handleCollectionChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="collectionRating">
              <Form.Label>Рейтинг</Form.Label>
              <Form.Control
                type="number"
                name="rating"
                value={collectionForm.rating}
                onChange={handleCollectionChange}
                required
              />
            </Form.Group>
            <Button variant="primary" type="submit">
              Добавить
            </Button>
          </Form>
        </Modal.Body>
      </Modal>

      {/* Add Book Modal */}
      <Modal show={showBookModal} onHide={() => setShowBookModal(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Добавить книгу</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleBookSubmit}>
            <Form.Group controlId="bookTitle">
              <Form.Label>Название</Form.Label>
              <Form.Control
                type="text"
                name="title"
                value={bookForm.title}
                onChange={handleBookChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="bookAuthor">
              <Form.Label>Автор</Form.Label>
              <Form.Control
                type="text"
                name="author"
                value={bookForm.author}
                onChange={handleBookChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="bookGenre">
              <Form.Label>Жанр</Form.Label>
              <Form.Control
                type="text"
                name="genre"
                value={bookForm.genre}
                onChange={handleBookChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="bookDescription">
              <Form.Label>Описание</Form.Label>
              <Form.Control
                as="textarea"
                rows={3}
                name="description"
                value={bookForm.description}
                onChange={handleBookChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="bookImage">
              <Form.Label>Image URL</Form.Label>
              <Form.Control
                type="text"
                name="image"
                value={bookForm.image}
                onChange={handleBookChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="bookBody">
              <Form.Label>Ссылка на книгу</Form.Label>
              <Form.Control
                type="text"
                name="body"
                value={bookForm.body}
                onChange={handleBookChange}
                required
              />
            </Form.Group>
            <Button variant="primary" type="submit">
              Добавить
            </Button>
          </Form>
        </Modal.Body>
      </Modal>
    </>
  );
};

export default MyBookShelf;
