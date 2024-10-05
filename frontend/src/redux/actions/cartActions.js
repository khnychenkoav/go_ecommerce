import axios from 'axios';

export const fetchCartItems = () => async (dispatch) => {
  const response = await axios.get('api/cart');
  dispatch({ type: 'FETCH_CART_ITEMS', payload: response.data });
};

export const addToCart = (productId) => async (dispatch) => {
    try {
      console.log('Добавление товара в корзину, ID:', productId);
      const response = await axios.post('/api/cart/add', {
        product_id: productId,
        quantity: 1,
      });
      console.log('Ответ от сервера при добавлении в корзину:', response.data);
      dispatch(fetchCartItems());
    } catch (error) {
      console.error('Ошибка при добавлении в корзину:', error.response || error);
      // Дополнительная обработка ошибки
    }
  };

export const removeFromCart = (id) => async (dispatch) => {
  await axios.delete(`api/cart/remove/${id}`);
  dispatch(fetchCartItems());
};
