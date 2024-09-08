import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const initialState = {
  title: "",
  description: "",
  livestreamExternalVariants: [],
  startTime: "",
  endTime: "",
} as {
  title: string;
  description: string;
  livestreamExternalVariants: {
    productId: number;
    variantId: number;
    name: string;
    imageUrl: string;
    option: { [key: string]: string };
    externalVariants: {
      externalVariantId: number;
      ecommerceId: number;
      price: number;
      quantity: number;
    }[];
  }[];
  startTime: string;
  endTime: string;
};

export const createReducer = createSlice({
  name: "livestream_create",
  initialState,
  reducers: {
    reset: () => initialState,
    setLivestreamInformation: (state, action: PayloadAction<any>) => {
      state.title = action.payload?.title;
      state.description = action.payload?.description;
      state.startTime = action.payload?.startTime;
      //state.currentStep = Math.min(state.currentStep + 1, 2);
    },
    setChosenLivestreamVariants: (state, action: PayloadAction<any>) => {
      state.livestreamExternalVariants = action.payload;
    },
  },
});

export const { reset, setLivestreamInformation, setChosenLivestreamVariants } =
  createReducer.actions;

export default createReducer.reducer;
