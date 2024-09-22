"use client";
import { configureStore } from "@reduxjs/toolkit";
import authReducer from "./auth";
import livestreamCreateReducer from "./livestream_create";
import loadingReducer from "./loading";
import checkoutReducer from "./checkout";
import { useSelector, useDispatch } from "react-redux";
import type { TypedUseSelectorHook } from "react-redux";

export const store = configureStore({
  reducer: {
    authReducer: authReducer,
    livestreamCreateReducer: livestreamCreateReducer,
    loadingReducer: loadingReducer,
    checkoutReducer: checkoutReducer,
  },
});

export type AppStore = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch: () => AppDispatch = useDispatch;
export const useAppSelector: TypedUseSelectorHook<AppStore> = useSelector;
