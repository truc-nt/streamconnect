import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const initialState = {
  open: false,
  message: "",
  type: "" as "error" | "warning" | "info" | "success",
} as {
  open: boolean;
  message: string;
  type: "error" | "warning" | "info" | "success";
};

export const alertReducer = createSlice({
  name: "alert",
  initialState,
  reducers: {
    setOpen: (state, action: PayloadAction<any>) => {
      state.message = action.payload.message;
      state.type = action.payload.type;
      state.open = true;
    },
    setClose: (state) => {
      state.open = false;
    },
  },
});

export const { setOpen, setClose } = alertReducer.actions;

export default alertReducer.reducer;
