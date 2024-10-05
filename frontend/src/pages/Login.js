import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { login } from '../redux/actions/authActions';
import { TextField, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [form, setForm] = useState({ email: '', password: '' });
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    dispatch(login(form, navigate));
  };

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        name="email"
        label="Email"
        value={form.email}
        onChange={handleChange}
        fullWidth
        required
        style={{ marginBottom: '1rem' }}
      />
      <TextField
        name="password"
        label="Пароль"
        type="password"
        value={form.password}
        onChange={handleChange}
        fullWidth
        required
        style={{ marginBottom: '1rem' }}
      />
      <Button variant="contained" color="primary" type="submit">
        Войти
      </Button>
    </form>
  );
}

export default Login;
