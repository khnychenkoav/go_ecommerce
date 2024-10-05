const initialState = {
    items: [],
    currentItem: null,
  };
  
  const productReducer = (state = initialState, action) => {
    switch (action.type) {
      case 'FETCH_PRODUCTS':
        return { ...state, items: action.payload };
      case 'FETCH_PRODUCT':
        return { ...state, currentItem: action.payload };
      default:
        return state;
    }
  };
  
  export default productReducer;
  