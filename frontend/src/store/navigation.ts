import { createSlice } from "@reduxjs/toolkit";

interface IPlaceholderState {
  active: string;
}

const initialState: IPlaceholderState = {
  active: "explore",
};

export const placeholderSlice = createSlice({
  name: "placeholder",
  initialState,
  reducers: {},
});

export default placeholderSlice.reducer;
