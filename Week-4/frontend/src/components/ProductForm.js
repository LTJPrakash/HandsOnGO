import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import { addProduct, getProducts } from '../api';

const ProductForm = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [price, setPrice] = useState(0);
  const [category, setCategory] = useState('');
  const [stock, setStock] = useState(0);

  useEffect(() => {
    if (id) {
      const fetchProduct = async () => {
        try {
          const response = await axios.get(`${addProduct}/${id}`);
          const product = response.data;
          setName(product.name);
          setDescription(product.description);
          setPrice(product.price);
          setCategory(product.category);
          setStock(product.stock);
        } catch (error) {
          console.error(error);
        }
      };
      fetchProduct();
    }
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const numericPrice = parseFloat(price); // Ensure price is a number
    const numericStock = parseInt(stock, 10); // Ensure stock is a numbe
    try {
      if (id) {
        await axios.put(`${addProduct}/${id}`, { name, description, price: numericPrice, category, stock:numericStock });
        alert('Product updated');
      } else {
        await axios.post(addProduct, { name, description, price: numericPrice, category, stock:numericStock });
        alert('Product created');
      }
      navigate('/products');
    } catch (error) {
      console.error(error);
      alert('Operation failed');
    }
  };

  return (
    <div>
      <h2>{id ? 'Update Product' : 'Add Product'}</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <textarea
          placeholder="Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <input
          type="number"
          placeholder="Price"
          value={price}
          onChange={(e) => setPrice(e.target.value)}
        />
        <input
          type="text"
          placeholder="Category"
          value={category}
          onChange={(e) => setCategory(e.target.value)}
        />
        <input
          type="number"
          placeholder="Stock"
          value={stock}
          onChange={(e) => setStock(e.target.value)}
        />
        <button type="submit">{id ? 'Update' : 'Create'}</button>
      </form>
      <button onClick={() => navigate('/products')}>Back to Product List</button>
    </div>
  );
};

export default ProductForm;
