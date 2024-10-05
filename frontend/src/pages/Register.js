import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { register } from '../redux/actions/authActions';
import { TextField, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';

function Register() {
  const [form, setForm] = useState({ name: '', email: '', password: '' });
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    dispatch(register(form, navigate));
  };

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        name="name"
        label="Имя"
        value={form.name}
        onChange={handleChange}
        fullWidth
        required
        style={{ marginBottom: '1rem' }}
      />
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
        Регистрация
      </Button>
    </form>
  );
}

export default Register;
