// slices/collectionsSlice.js
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axiosInstance from '../axiosInstance';

export const fetchCollections = createAsyncThunk('collections/fetchCollections', async () => {
  const response = await axiosInstance.get('/collections');
  return response.data;
});

export const fetchCollectionById = createAsyncThunk('collections/fetchCollectionById', async (collectionId) => {
  const response = await axiosInstance.get(`/collections/${collectionId}`);
  return response.data;
});


export const fetchCollectionByUser = createAsyncThunk('collections/fetchCollectionByUser', async (userId) => {
  const response = await axiosInstance.get(`/collections/user/${userId}`);
  return response.data;
});

export const addCollection = createAsyncThunk('collections/addCollection', async (bookData) => {
  const response = await axiosInstance.post(`/collections`, bookData);
  return response.data;
});

const collectionsSlice = createSlice({
  name: 'collections',
  initialState: {
    collections: [],
    collectionsUser: [],
    status: 'idle',
    error: null,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchCollections.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchCollections.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.collections = action.payload;
      })
      .addCase(fetchCollections.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message;
      })
      .addCase(fetchCollectionById.fulfilled, (state, action) => {
        state.selectedCollection = action.payload;
      })
      .addCase(fetchCollectionByUser.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchCollectionByUser.fulfilled, (state, action) => {
        state.collectionsUser = action.payload;
      })
      .addCase(addCollection.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.collections.push(action.payload);
      });
  },
});

export default collectionsSlice.reducer;