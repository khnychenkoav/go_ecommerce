import axios from 'axios';

export const fetchProducts = () => async (dispatch) => {
  const response = await axios.get('api/products');
  dispatch({ type: 'FETCH_PRODUCTS', payload: response.data });
};

export const fetchProduct = (id) => async (dispatch) => {
    try {
        const response = await axios.get(`/api/products/${id}`);
        dispatch({ type: 'FETCH_PRODUCT', payload: response.data });
      } catch (error) {
        console.error('Ошибка при получении товара:', error);
        // Вы можете добавить действие для обработки ошибки в Redux
      }
};
