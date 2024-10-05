import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchCartItems, removeFromCart } from '../redux/actions/cartActions';
import { Typography, Button } from '@mui/material';

function Cart() {
  const dispatch = useDispatch();
  const cartItems = useSelector((state) => state.cart.items);

  useEffect(() => {
    dispatch(fetchCartItems());
  }, [dispatch]);

  useEffect(() => {
    console.log('Обновленные товары в корзине:', cartItems);
  }, [cartItems]);

  const handleRemove = (id) => {
    dispatch(removeFromCart(id));
  };

  return (
    <div>
      <Typography variant="h4">Ваша корзина</Typography>
      {cartItems.length === 0 ? (
        <Typography>Ваша корзина пуста</Typography>
      ) : (
        cartItems.map((item) => (
          <div key={item.id}>
            <Typography variant="h6">{item.product.name}</Typography>
            <Typography>Количество: {item.quantity}</Typography>
            <Button variant="contained" onClick={() => handleRemove(item.id)}>
              Удалить
            </Button>
          </div>
        ))
      )}
    </div>
  );
}

export default Cart;
