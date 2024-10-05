import axios from 'axios';

export const register = (formData, navigate) => async (dispatch) => {
  try {
    await axios.post('api/register', formData);
    dispatch({ type: 'REGISTER_SUCCESS' });
    navigate('/login');
  } catch (error) {
    console.error('Ошибка регистрации', error);
  }
};

export const login = (formData, navigate) => async (dispatch) => {
  try {
    const response = await axios.post('api/login', formData);
    const token = response.data.token;
    localStorage.setItem('token', token);
    axios.defaults.headers.common['Authorization'] = token;
    dispatch({ type: 'LOGIN_SUCCESS', payload: { token } });
    navigate('/');
  } catch (error) {
    console.error('Ошибка входа', error);
  }
};

export const logout = () => (dispatch) => {
  localStorage.removeItem('token');
  delete axios.defaults.headers.common['Authorization'];
  dispatch({ type: 'LOGOUT' });
};
