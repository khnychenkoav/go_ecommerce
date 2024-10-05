import React from 'react';
import {
  Card,
  CardContent,
  Typography,
  Button,
  CardActions,
} from '@mui/material';
import { Link } from 'react-router-dom';

function ProductCard({ product }) {
  return (
    <Card style={{ margin: '1rem', width: '300px' }}>
      <CardContent>
        <Typography variant="h5">{product.name}</Typography>
        <Typography>{product.description}</Typography>
        <Typography>{product.price} руб.</Typography>
      </CardContent>
      <CardActions>
        <Button
          variant="contained"
          size="small"
          component={Link}
          to={`/products/${product.id}`}
        >
          Подробнее
        </Button>
      </CardActions>
    </Card>
  );
}

export default ProductCard;
