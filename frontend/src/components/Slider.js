// components/Slider.js
import React, { useState } from 'react';
import { useSelector } from 'react-redux';
import { Carousel, Modal, Button } from 'react-bootstrap';

const Slider = () => {
  const books = useSelector((state) => state.books.books);
  const firstThreeBooks = books.slice(0, 3);

  const [show, setShow] = useState(false);
  const [selectedBook, setSelectedBook] = useState(null);

  const handleClose = () => setShow(false);
  const handleShow = (book) => {
    setSelectedBook(book);
    setShow(true);
  };

  return (
    <>
      <Carousel className='slider'>
        {firstThreeBooks.map((book) => (
          <Carousel.Item key={book.book_id} onClick={() => handleShow(book)}>
            <img className="d-block w-100" src={book.image} alt={book.title}/>
            <Carousel.Caption>
              <h3>{book.title}</h3>
              <p>{book.description}</p>
            </Carousel.Caption>
          </Carousel.Item>
        ))}
      </Carousel>

      {selectedBook && (
        <Modal show={show} onHide={handleClose}>
          <Modal.Header closeButton>
            <Modal.Title>{selectedBook.title}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <img src={selectedBook.image} alt={selectedBook.title} className="img-fluid" style={{ float:'right', width: '15rem', height: '20rem', margin: '10px', objectFit: 'cover'}}/>
            <p><strong>Author:</strong> {selectedBook.author}</p>
            <p><strong>Genre:</strong> {selectedBook.genre}</p>
            <p>{selectedBook.body}</p>
          </Modal.Body>
        </Modal>
      )}
    </>
  );
};

export default Slider;
