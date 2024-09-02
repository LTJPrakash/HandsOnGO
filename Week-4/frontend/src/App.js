import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import UserRegister from './components/UserRegister';
import UserLogin from './components/UserLogin';
import ProductList from './components/ProductList';
import ProductForm from './components/ProductForm';
import Navigation from './Navigation';

const App = () => {
  return (
    <Router>
      <div>
        <h1>Product Management</h1>
        <Navigation />
        <Routes>
          <Route path="/register" element={<UserRegister />} />
          <Route path="/login" element={<UserLogin />} />
          <Route path="/products" element={<ProductList />} />
          <Route path="/products/new" element={<ProductForm />} />
          <Route path="/products/edit/:id" element={<ProductForm />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
