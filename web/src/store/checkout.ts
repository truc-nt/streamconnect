import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const initialState = {
  currentStep: 0,
  cartItemIds: [],
  prices: [],
  address: "",
} as {
  currentStep: number;
  cartItemIds: number[];
  prices: {
    ecommerceId: number;
    subTotal: number;
    shippingFee: number;
    discountTotal: number;
  }[];
  address: string;
};

export const checkoutReducer = createSlice({
  name: "checkout",
  initialState,
  reducers: {
    reset: () => initialState,
    setPrevStep: (state) => {
      state.currentStep = Math.max(state.currentStep - 1, 0);
    },
    setNextStep: (state) => {
      state.currentStep = Math.min(state.currentStep + 1, 2);
    },
    setCartItemIds: (state, action: PayloadAction<any>) => {
      state.cartItemIds = action.payload;
    },
    setPrices: (state, action: PayloadAction<any>) => {
      state.prices = action.payload;
    },
    setAddress: (state, action: PayloadAction<any>) => {
      state.address = action.payload;
    },
  },
});

export const { reset, setPrevStep, setNextStep, setCartItemIds, setPrices } =
  checkoutReducer.actions;

export default checkoutReducer.reducer;
