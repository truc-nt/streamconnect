import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const initialState = {
  currentStep: 0,
  title: "",
  description: "",
  livestreamExternalVariants: [],
  startTime: "",
  endTime: "",
} as {
  currentStep: number;
  title: string;
  description: string;
  livestreamExternalVariants: {
    idProduct: number;
    idVariant: number;
    name: string;
    option: Record<string, string>;
    externalVariants: {
      idVariant: number;
      idExternalVariant: number;
      idEcommerce: number;
      price: number;
      stock: number;
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
    setPrevStep: (state) => {
      state.currentStep = Math.max(state.currentStep - 1, 0);
    },
    setNextStep: (state) => {
      state.currentStep = Math.min(state.currentStep + 1, 2);
    },
    setLivestreamInformation: (state, action: PayloadAction<any>) => {
      state.title = action.payload?.title;
      state.description = action.payload?.description;
      state.startTime = action.payload?.startTime;
      //state.currentStep = Math.min(state.currentStep + 1, 2);
    },
    setChosenLivestreamVariants: (state, action: PayloadAction<any>) => {
      state.livestreamExternalVariants = action.payload;
      state.currentStep = Math.min(state.currentStep + 1, 2);
    },
  },
});

export const {
  reset,
  setPrevStep,
  setNextStep,
  setLivestreamInformation,
  setChosenLivestreamVariants,
} = createReducer.actions;

export default createReducer.reducer;
