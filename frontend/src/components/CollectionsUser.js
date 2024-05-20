import React, { useEffect, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { fetchCollectionById, fetchCollectionByUser, fetchCollections } from '../redux/slice/collectionsSlice';
import { Card, Row, Col, Modal, Button } from 'react-bootstrap';
import Cookies from 'universal-cookie';

const CollectionsUser = () => {
    const dispatch = useDispatch();
    const collections = useSelector((state) => state.collections.collectionsUser);
    const collectionStatus = useSelector((state) => state.collections.status);
    const selectedCollection = useSelector((state) => state.collections.selectedCollection);

    const [showCollectionModal, setShowCollectionModal] = useState(false);
    const [showBookModal, setShowBookModal] = useState(false);
    const [selectedBook, setSelectedBook] = useState(null);
    const cookie = new Cookies();
    const user_id = cookie.get('user_id');

    useEffect(() => {
        if (collectionStatus === 'idle') {
            dispatch(fetchCollectionByUser(user_id));
            dispatch(fetchCollections());
        }
    }, []);

    const handleCloseCollectionModal = () => setShowCollectionModal(false);
    const handleCloseBookModal = () => setShowBookModal(false);

    const handleShowCollectionModal = (collectionId) => {
        dispatch(fetchCollectionById(collectionId));
        setShowCollectionModal(true);
    };

    const handleShowBookModal = (book) => {
        setSelectedBook(book);
        setShowBookModal(true);
    };

    return (
        <>
            <Row>
                {collections.map((collection) => (
                    <Col key={collection.collection_id} md={4}>
                        <Card onClick={() => handleShowCollectionModal(collection.collection_id)} style={{ marginBottom: '10px', backgroundColor: '#bbd6ee'}}>
                            <Card.Body>
                                <Card.Title>{collection.name}</Card.Title>
                                <Card.Text>{collection.description}</Card.Text>
                            </Card.Body>
                        </Card>
                    </Col>
                ))}
            </Row>

            {selectedCollection && (
                <Modal show={showCollectionModal} onHide={handleCloseCollectionModal}>
                    <Modal.Header closeButton>
                        <Modal.Title>{selectedCollection.name}</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <p><strong>Description:</strong> {selectedCollection.description}</p>
                        <p><strong>Rating:</strong> {selectedCollection.rating}</p>
                        <h5>{selectedCollection.books !== null && "Books in this collection:"}</h5>
                        {selectedCollection.books !== null && selectedCollection.books.map((book) => (
                            <Card key={book.book_id} className="mb-3" onClick={() => handleShowBookModal(book)} style={{ width: '13rem', display: 'inline-block', margin: '10px' }}>
                                <Card.Img variant="top" src={book.image} style={{ height: '16rem', objectFit: 'cover' }}/>
                                <Card.Body>
                                    <Card.Title>{book.title}</Card.Title>
                                    <Card.Text>{book.description}</Card.Text>
                                </Card.Body>
                            </Card>
                        ))}
                    </Modal.Body>
                    <Modal.Footer>
                        <Button>
                            Обновить
                        </Button>
                    </Modal.Footer>
                </Modal>
            )}

            {selectedBook && (
                <Modal show={showBookModal} onHide={handleCloseBookModal}>
                    <Modal.Header closeButton>
                        <Modal.Title>{selectedBook.title}</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <Card.Img variant="top" src={selectedBook.image}  style={{ float:'right', width: '15rem', height: '20rem', margin: '10px', objectFit: 'cover'}}/>
                        <p><strong>Автор:</strong> {selectedBook.author}</p>
                        <p><strong>Описание:</strong> {selectedBook.description}</p>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant='danger'>
                            Удалить из коллекции
                        </Button>
                    </Modal.Footer>
                </Modal>
            )}
        </>
    );
};

export default CollectionsUser;
