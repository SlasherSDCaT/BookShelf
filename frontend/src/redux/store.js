// store.js
import { configureStore } from '@reduxjs/toolkit';
import collectionsReducer from './slice/collectionsSlice';
import booksReducer from './slice/booksSlice';

const store = configureStore({
  reducer: {
    collections: collectionsReducer,
    books: booksReducer,
  },
});

export default store;
