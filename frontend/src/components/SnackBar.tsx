import React from "react";
import { FlashMessageContext } from "../contexts/FlashMessageContext";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

export const FlashMessage: React.FC = () => {
  const { flashMessage } = React.useContext(FlashMessageContext);
  return (
    <Snackbar open={flashMessage !== ""}>
      <Alert severity="success" sx={{ width: "100%" }}>
        {flashMessage}
      </Alert>
    </Snackbar>
  );
};
