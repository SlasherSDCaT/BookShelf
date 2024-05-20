// components/Books.js
import React, { useEffect, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { fetchBooks, fetchComment } from '../redux/slice/booksSlice';
import { Card, Row, Col, Modal, Button, Accordion, ListGroup } from 'react-bootstrap';

const Books = () => {
    const dispatch = useDispatch();
    const books = useSelector((state) => state.books.books);
    const bookStatus = useSelector((state) => state.books.status);
    
    const [show, setShow] = useState(false);
    const [selectedBook, setSelectedBook] = useState(null);
    
    const comments = useSelector((state) => state.books.comment);
    useEffect(() => {
      if (bookStatus === 'idle') {
        dispatch(fetchBooks());
      }
    }, [bookStatus, dispatch]);
  
    const handleClose = () => setShow(false);
    const handleShow = (book) => {
      setSelectedBook(book);
      dispatch(fetchComment(book.book_id));
      setShow(true);
    };
  
    return (
      <>
        <Row>
          {books.map((book) => (
            <Col key={book.book_id} md={3}>
              <Card onClick={() => handleShow(book)} style={{ width: '18rem', marginBottom: '10px'}}>
                <Card.Img variant="top" src={book.image} style={{ height: '20rem', objectFit: 'cover'}}/>
                <Card.Body>
                  <Card.Title>{book.title}</Card.Title>
                </Card.Body>
              </Card>
            </Col>
          ))}
        </Row>
  
        {selectedBook && (
          <Modal show={show} onHide={handleClose}>
            <Modal.Header closeButton>
              <Modal.Title>{selectedBook.title}</Modal.Title>
            </Modal.Header>
            <Modal.Body><img src={selectedBook.image} alt={selectedBook.title} className="img-fluid" style={{ 
                                                      float: 'right', 
                                                      width: '200px', 
                                                      height: '300px',
                                                      objectFit: 'cover',
                                                      marginRight: '10px', 
                                                      marginBottom: '10px' 
                                                    }}/>
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
              <Button variant="primary">
                Добавить отзыв
              </Button>
              <Button variant="primary">
                Добавить в коллекцию
              </Button>
              <Button variant="primary">
                Обновить
              </Button>
            </Modal.Footer>
          </Modal>
        )}
      </>
    );
  };
  
  export default Books;