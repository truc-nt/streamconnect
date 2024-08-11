import { Alert as MuiAlert } from "@mui/material";
import { useTimeout } from "usehooks-ts";
import { useAppDispatch } from "@/store/store";
import { setClose } from "@/store/alert";

const Alert = ({
  message,
  type,
}: {
  message: string;
  type: "error" | "warning" | "info" | "success";
}) => {
  const dispatch = useAppDispatch();
  useTimeout(() => dispatch(setClose()), 3000);

  return (
    <div
      style={{
        position: "fixed",
        top: 20,
        left: "50%",
        transform: "translateX(-50%)",
        zIndex: 1300,
      }}
    >
      <MuiAlert variant="filled" severity={type}>
        {message}
      </MuiAlert>
      ;
    </div>
  );
};

export default Alert;
