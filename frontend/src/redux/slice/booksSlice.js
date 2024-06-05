// slices/booksSlice.js
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axiosInstance from '../axiosInstance';

export const fetchBooks = createAsyncThunk('books/fetchBooks', async () => {
    const response = await axiosInstance.get('/books');
    return response.data;
  });
  
  export const fetchComment = createAsyncThunk('books/fetchComment', async (book_id) => {
    const response = await axiosInstance.get(`/comments/${book_id}`);
    return response.data;
  });

  export const addBook = createAsyncThunk('books/addBook', async (bookData) => {
    const response = await axiosInstance.post('/books', bookData);
    return response.data;
  });

  export const deleteBook = createAsyncThunk('books/deleteBook', async (book_id) => {
    const response = await axiosInstance.delete(`/books/${book_id}`);
    return response.data;
  });
  export const addComment = createAsyncThunk('books/addComment', async (commentData) => {
    console.log(commentData)
    const response = await axiosInstance.post('/comments', commentData);
    return response.data;
});
  
  const booksSlice = createSlice({
    name: 'books',
    initialState: {
      books: [],
      comment: [],
      status: 'idle',
      error: null,
    },
    reducers: {},
    extraReducers: (builder) => {
      builder
        .addCase(fetchBooks.pending, (state) => {
          state.status = 'loading';
        })
        .addCase(fetchBooks.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.books = action.payload;
        })
        .addCase(fetchBooks.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
        .addCase(addBook.pending, (state) => {
          state.status = 'loading';
        })
        .addCase(addBook.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.books.push(action.payload);
        })
        .addCase(addBook.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
        .addCase(fetchComment.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.comment = action.payload;
        })
        .addCase(addComment.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.comment.push(action.payload);
        })
        .addCase(deleteBook.fulfilled, (state, action) => {
          state.status = 'succeeded';
        });
    },
  });  

export default booksSlice.reducer;
