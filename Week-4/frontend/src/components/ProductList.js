import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import { getProducts } from '../api';

const ProductList = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await axios.get(getProducts);
        setProducts(response.data.Data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchProducts();
  }, []);

  return (
    <div>
      <h2>Product List</h2>
      <Link to="/products/new"><button>Add New Product</button></Link>
      <ul>
        {products.map((product) => (
          <li key={product._id}>
            {product.name} - ${product.price}
            <Link to={`/products/edit/${product._id}`}><button>Edit</button></Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ProductList;
