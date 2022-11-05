import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { theme } from "./materialui/theme";
import { ThemeProvider } from "@mui/material/styles";
import { Index } from "./pages/Index";
import { Header } from "./components/Header";
import { Global } from "@emotion/react";
import { globalStyle } from "./styles/global";
import { AuthContextProvider } from "./contexts/authContext";

export const App: React.FC = () => {
  return (
    <div>
      <AuthContextProvider>
        <ThemeProvider theme={theme}>
          <Global
            styles={globalStyle}
          />
          <BrowserRouter>
            <Header />
            <Routes>
                <Route path={"/"} element={<Index />} />
            </Routes>
          </BrowserRouter>
        </ThemeProvider>
      </AuthContextProvider>
    </div>
  );
};
