import { useAppDispatch } from "@/store/store";
import { setOpen, setClose } from "@/store/loading";
import { useCallback } from "react";
import { App } from "antd";
import { AxiosError } from "axios";

const useLoading = (
  loadingFetch: (...params: any[]) => Promise<any>,
  successMessage: string = "",
  errorMessage: string = "",
  onSuccess?: (res: any) => void,
  onError?: (error: any) => void,
) => {
  const dispatch = useAppDispatch();
  const { message } = App.useApp();

  const execute = useCallback(async (...params: any[]) => {
    dispatch(setOpen());
    try {
      const res = await loadingFetch(...params);
      if (successMessage) message.success(successMessage);
      onSuccess?.(res);
      return res;
    } catch (error: AxiosError | any) {
      const errMsg = errorMessage || error?.message || "An error occurred";
      message.error(errMsg);
      onError?.(error);
    } finally {
      dispatch(setClose());
    }
  }, []);

  return execute;
};

export default useLoading;
