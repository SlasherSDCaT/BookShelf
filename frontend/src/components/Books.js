import React, { useEffect, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { fetchBooks, fetchComment, addComment } from '../redux/slice/booksSlice';
import { Card, Row, Col, Modal, Button, Accordion, ListGroup, Form } from 'react-bootstrap';
import { addBookToCollection } from '../redux/slice/collectionsSlice';
import Cookies from 'universal-cookie';

const Books = () => {
  const dispatch = useDispatch();
  const books = useSelector((state) => state.books.books);
  const bookStatus = useSelector((state) => state.books.status);
  const collections = useSelector((state) => state.collections.collectionsUser);
  const comments = useSelector((state) => state.books.comment);
  const cookie = new Cookies();
  const user_id = cookie.get('user_id');

  const [showBookModal, setShowBookModal] = useState(false);
  const [showCollectionModal, setShowCollectionModal] = useState(false);
  const [showCommentModal, setShowCommentModal] = useState(false);
  const [selectedBook, setSelectedBook] = useState(null);
  const [commentForm, setCommentForm] = useState({ user_id: user_id, book_id: '', rating: 2, text: '' });

  useEffect(() => {
      if (bookStatus === 'idle') {
          dispatch(fetchBooks());
      }
  }, [bookStatus, dispatch]);

  const handleCloseBookModal = () => setShowBookModal(false);
  const handleShowBookModal = (book) => {
      setSelectedBook(book);
      dispatch(fetchComment(book.book_id));
      setShowBookModal(true);
  };

  const handleShowCollectionModal = () => setShowCollectionModal(true);
  const handleCloseCollectionModal = () => setShowCollectionModal(false);

  const handleShowCommentModal = () => {
      setCommentForm({ ...commentForm, book_id: selectedBook.book_id });
      setShowCommentModal(true);
  };
  const handleCloseCommentModal = () => setShowCommentModal(false);

  const handleAddToCollection = (collectionId, bookId) => {
    console.log('load ' + bookId);
    dispatch(addBookToCollection({ collectionId, bookId }));
    setShowCollectionModal(false);
  };

  const handleCommentChange = (e) => {
    const { name, value } = e.target;
    setCommentForm({ ...commentForm, [name]: name === 'rating' ? Number(value) : value });
};


  const handleCommentSubmit = (e) => {
      e.preventDefault();
      dispatch(addComment(commentForm));
      setShowCommentModal(false);
      setShowBookModal(false);
  };

  return (
      <>
          <Row>
              {books.map((book) => (
                  <Col key={book.book_id} md={3}>
                      <Card onClick={() => handleShowBookModal(book)} style={{ width: '18rem', marginBottom: '10px' }}>
                          <Card.Img variant="top" src={book.image} style={{ height: '20rem', objectFit: 'cover' }} />
                          <Card.Body>
                              <Card.Title>{book.title}</Card.Title>
                          </Card.Body>
                      </Card>
                  </Col>
              ))}
          </Row>

          {selectedBook && (
              <Modal show={showBookModal} onHide={handleCloseBookModal}>
                  <Modal.Header closeButton>
                      <Modal.Title>{selectedBook.title}</Modal.Title>
                  </Modal.Header>
                  <Modal.Body>
                      <img src={selectedBook.image} alt={selectedBook.title} className="img-fluid" style={{
                          float: 'right',
                          width: '200px',
                          height: '300px',
                          objectFit: 'cover',
                          marginRight: '10px',
                          marginBottom: '10px'
                      }} />
                      <p><strong>Автор:</strong> {selectedBook.author}</p>
                      <p><strong>Жанр:</strong> {selectedBook.genre}</p>
                      <p><strong>Ссылка:</strong> {selectedBook.body}</p>
                      <Accordion>
                          <Accordion.Item eventKey="0">
                              <Accordion.Header>Посмотреть описание</Accordion.Header>
                              <Accordion.Body>
                                  <p>{selectedBook.description}</p>
                              </Accordion.Body>
                          </Accordion.Item>
                      </Accordion>
                      <Accordion>
                          <Accordion.Item eventKey="1">
                              <Accordion.Header>Посмотреть отзывы</Accordion.Header>
                              <Accordion.Body>
                                  <ListGroup>
                                      {comments.map((comment) => (
                                          <ListGroup.Item key={comment.id}>
                                              {comment.text}
                                              {/* <Button variant="secondary">
                                                  Удалить
                                              </Button> */}
                                          </ListGroup.Item>
                                      ))}
                                  </ListGroup>
                              </Accordion.Body>
                          </Accordion.Item>
                      </Accordion>
                  </Modal.Body>
                  <Modal.Footer>
                      <Button variant="primary" onClick={handleShowCommentModal}>
                          Добавить отзыв
                      </Button>
                      <Button variant="primary" onClick={handleShowCollectionModal}>
                          Добавить в коллекцию
                      </Button>
                  </Modal.Footer>
              </Modal>
          )}

          <Modal show={showCollectionModal} onHide={handleCloseCollectionModal}>
              <Modal.Header closeButton>
                  <Modal.Title>Добавить в коллекцию</Modal.Title>
              </Modal.Header>
              <Modal.Body>
                  <ListGroup>
                      {collections.map((collection) => (
                          <ListGroup.Item key={collection.id}>
                              {collection.name}
                              <Button variant="primary" onClick={() => handleAddToCollection(collection.collection_id, selectedBook.book_id)} style={{ float: 'right' }}>
                                  Добавить
                              </Button>
                          </ListGroup.Item>
                      ))}
                  </ListGroup>
              </Modal.Body>
          </Modal>

          <Modal show={showCommentModal} onHide={handleCloseCommentModal}>
              <Modal.Header closeButton>
                  <Modal.Title>Добавить отзыв</Modal.Title>
              </Modal.Header>
              <Modal.Body>
                  <Form onSubmit={handleCommentSubmit}>
                      <Form.Group controlId="commentText">
                          <Form.Label>Текст отзыва</Form.Label>
                          <Form.Control
                              as="textarea"
                              rows={3}
                              name="text"
                              value={commentForm.text}
                              onChange={handleCommentChange}
                              required
                          />
                      </Form.Group>
                      <Form.Group controlId="commentRating">
                          <Form.Label>Рейтинг</Form.Label>
                          <Form.Control
                              type="number"
                              name="rating"
                              value={commentForm.rating}
                              onChange={handleCommentChange}
                              min={1}
                              max={5}
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

export default Books;