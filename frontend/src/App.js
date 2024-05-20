import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import './App.css';
import MainPage from './pages/MainPage/MainPage';
import RegPage from './pages/RegPage/RegPage';
import LogPage from './pages/LogPage/LogPage';
import { Provider } from 'react-redux';
import store from './redux/store';
import MyBookShelf from './pages/MyBookShelf/MyBookShelf';

function App() {
    return (
    <Provider store={store}>
      <BrowserRouter className="App">
        <Routes>
          <Route path='/' element={<MainPage/>}/>
          <Route path='/main' element={<MainPage/>}/>
          <Route path='/myshelf' element={<MyBookShelf/>}/>
          <Route path='/register' element={<RegPage/>}/>
          <Route path='/login' element={<LogPage/>}/>
        </Routes>
      </BrowserRouter>
    </Provider>
  );
}

export default App;
