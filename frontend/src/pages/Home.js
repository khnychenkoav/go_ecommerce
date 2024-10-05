import React from 'react';
import { Typography } from '@mui/material';

function Home() {
  return (
    <div>
      <Typography variant="h4" gutterBottom>
        Добро пожаловать в наш интернет-магазин!
      </Typography>
      <Typography>
        Здесь вы найдете множество товаров по отличным ценам.
      </Typography>
    </div>
  );
}

export default Home;
