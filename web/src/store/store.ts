"use client";
import { configureStore } from "@reduxjs/toolkit";
import livestreamCreateReducer from "./livestream_create";
import loadingReducer from "./loading";
import { useSelector, useDispatch } from "react-redux";
import type { TypedUseSelectorHook } from "react-redux";

export const store = configureStore({
  reducer: {
    livestreamCreateReducer: livestreamCreateReducer,
    loadingReducer: loadingReducer,
  },
});

export type AppStore = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch: () => AppDispatch = useDispatch;
export const useAppSelector: TypedUseSelectorHook<AppStore> = useSelector;
