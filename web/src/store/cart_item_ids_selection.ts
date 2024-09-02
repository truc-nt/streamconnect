import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { Key } from "react";

const initialState = {
  cartItems: [],
} as {
  cartItems: {
    cartItemId: Key;
    price: 0;
    quantity: 0;
  }[];
};

export const cartItemIdsSelection = createSlice({
  name: "cart_item_ids_selection",
  initialState,
  reducers: {
    reset: () => initialState,
    setSelectedCartItems: (state, action: PayloadAction<any>) => {
      state.cartItems = action.payload;
    },
  },
});

export const { reset, setSelectedCartItems } = cartItemIdsSelection.actions;

export default cartItemIdsSelection.reducer;
