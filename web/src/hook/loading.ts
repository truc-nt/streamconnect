import { useAppDispatch } from "@/store/store";
import { setOpen, setClose } from "@/store/loading";
import { useCallback } from "react";
import { App } from "antd";

const useLoading = (
  loadingFetch: (...params: any[]) => Promise<any>,
  successMessage: string,
  errorMessage: string,
) => {
  const dispatch = useAppDispatch();
  const { message } = App.useApp();

  const execute = useCallback(
    async (...params: any[]) => {
      dispatch(setOpen());
      try {
        await loadingFetch(...params);
        message.success(successMessage);
      } catch (error) {
        message.error(errorMessage);
        throw error;
      } finally {
        dispatch(setClose());
      }
    },
    [loadingFetch, successMessage, errorMessage, dispatch, message],
  );

  return execute;
};

export default useLoading;
