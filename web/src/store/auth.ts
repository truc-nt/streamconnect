import { createSlice } from "@reduxjs/toolkit";

const authSlice = createSlice({
  name: "auth",
  initialState: {
    userId: null,
    username: null,
    isShowLoginModal: false,
  },
  reducers: {
    setLogin(state, action) {
      state.userId = action.payload.userId;
      state.username = action.payload.username;
      state.isShowLoginModal = false;
    },
    setLogout(state) {
      state.userId = null;
      state.username = null;
    },
    toggleLoginModal(state) {
      state.isShowLoginModal = !state.isShowLoginModal;
    },
  },
});

export const { setLogin, setLogout, toggleLoginModal } = authSlice.actions;
export default authSlice.reducer;
