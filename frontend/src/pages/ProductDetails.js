import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchProduct } from '../redux/actions/productActions';
import { addToCart } from '../redux/actions/cartActions';
import { Typography, Button } from '@mui/material';
import { useParams } from 'react-router-dom';

function ProductDetails() {
  const dispatch = useDispatch();
  const { id } = useParams();
  const product = useSelector((state) => state.products.currentItem);

  useEffect(() => {
    if (id) {
      dispatch(fetchProduct(id));
    }
  }, [dispatch, id]);

  const handleAddToCart = () => {
    if (product) {
      dispatch(addToCart(product.id));
    }
  };

  if (!product) return <div>Загрузка...</div>;

  return (
    <div>
      <Typography variant="h4">{product.name}</Typography>
      <Typography>{product.description}</Typography>
      <Typography>{product.price} руб.</Typography>
      <Button variant="contained" onClick={handleAddToCart}>
        Добавить в корзину
      </Button>
    </div>
  );
}

export default ProductDetails;
