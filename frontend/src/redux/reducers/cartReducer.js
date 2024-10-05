const initialState = {
    items: [],
  };
  
  const cartReducer = (state = initialState, action) => {
    switch (action.type) {
      case 'FETCH_CART_ITEMS':
        return { ...state, items: action.payload };
      default:
        return state;
    }
  };
  
  export default cartReducer;
  