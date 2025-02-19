import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import Register from './pages/Register';
import Login from './pages/Login';
import ItemList from './pages/ItemList';
import ItemDetails from './pages/ItemDetails';
import ProtectedRoute from './components/ProtectedRoute';
import { isAuthenticated } from './Auth';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <ProtectedRoute path="/items" element={<ItemList />} isAuthenticated={isAuthenticated()} />
        <ProtectedRoute path="/items/:id" element={<ItemDetails />} isAuthenticated={isAuthenticated()} />
      </Routes>
    </Router>
  );
}

export default App;