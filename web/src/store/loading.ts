import { createSlice } from "@reduxjs/toolkit";

type LoadingState = {
  open: boolean;
};

const initialState = {
  open: false,
} as LoadingState;

export const loading = createSlice({
  name: "loading",
  initialState,
  reducers: {
    setOpen: (state) => {
      state.open = true;
    },
    setClose: (state) => {
      state.open = false;
    },
  },
});

export const { setOpen, setClose } = loading.actions;
export default loading.reducer;
