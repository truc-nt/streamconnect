"use client";
import { configureStore } from "@reduxjs/toolkit";
import placeholderReducer from "./navigation";

export const store = configureStore({
  reducer: {
    placeholder: placeholderReducer,
  },
});

export type RootStore = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
