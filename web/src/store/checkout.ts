import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const initialState = {
  currentStep: 0,
  addressId: null,
  groupByShop: [],
} as {
  currentStep: number;
  addressId: number | null;
  groupByShop: {
    shopId: number | null;
    groupByEcommerce: {
      cartItemIds: number[];
      ecommerceId: number;
      subTotal: number;
      voucherIds: number[];
      internalDiscountTotal: number;
    }[];
  }[];
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
    setGroupByShop: (state, action: PayloadAction<any>) => {
      console.log(action.payload);
      state.groupByShop = action.payload;
    },
    setAddressId: (state, action: PayloadAction<any>) => {
      state.addressId = action.payload;
    },
  },
});

export const { reset, setPrevStep, setNextStep, setGroupByShop, setAddressId } =
  checkoutReducer.actions;

export default checkoutReducer.reducer;
