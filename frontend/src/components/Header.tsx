import React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { DamionFont } from "../materialui/theme";
import { CourgetteFont } from "../materialui/theme";
import { ThemeProvider } from "@mui/material/styles";
// import { AuthContext } from "../contexts/AuthContext";
import AuthService from "../auth";
import { useNavigate } from "react-router-dom";
// import { FlashMessageDispatchContext } from "../contexts/FlashMessageContext";

export const Header: React.FC = () => {
//   const authContext = React.useContext(AuthContext);
  const navigate = useNavigate();
//   const { dispatch } = React.useContext(FlashMessageDispatchContext);

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <ThemeProvider theme={DamionFont}>
            <Typography variant="h4" component="div" sx={{ flexGrow: 1 }}>
              EBookMerge
            </Typography>
          </ThemeProvider>
          <ThemeProvider theme={CourgetteFont}>
              <Button color="inherit" onClick={() => {
                AuthService.login();
                // navigate("/ebooks");
              }}>LogIn</Button>
            </ThemeProvider>
        </Toolbar>
      </AppBar>
    </Box>
  );
};
