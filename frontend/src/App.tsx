import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { theme } from "./materialui/theme";
import { ThemeProvider } from "@mui/material/styles";
import { Top } from "./pages/Top";
import { Header } from "./components/Header";
import { Global } from "@emotion/react";
import { globalStyle } from "./styles/global";
import { AuthContextProvider, AuthContext } from "./contexts/AuthContext";
import { ApiClientContextProvider } from "./contexts/ApiClientContext";
import { EBookList } from "./pages/EBookList";
import { FlashMessageProvider } from "./contexts/FlashMessageContext";
import { FlashMessage } from "./components/SnackBar";
import Container from "@mui/material/Container";
import { GmailOAuth2Callback } from "./pages/GmailOAuth2Callback";

export const App: React.FC = () => {
  return (
    <div>
      <AuthContextProvider>
        <ApiClientContextProvider>
          <FlashMessageProvider>
            <ThemeProvider theme={theme}>
              <Global styles={globalStyle} />
              <Initializing>
                <BrowserRouter>
                  <FlashMessage />
                  <Header />
                  <Container
                    sx={{ mt: 2, mx: "auto", width: { xs: "80%", md: "60%" } }}
                  >
                    <Routes>
                      <Route path={"/"} element={<Top />} />
                      <Route path={"/ebooks"} element={<EBookList />} />
                      <Route
                        path={"/oauth2callback"}
                        element={<GmailOAuth2Callback />}
                      />
                    </Routes>
                  </Container>
                </BrowserRouter>
              </Initializing>
            </ThemeProvider>
          </FlashMessageProvider>
        </ApiClientContextProvider>
      </AuthContextProvider>
    </div>
  );
};

const Initializing: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const initialized = React.useContext(AuthContext).initialized;
  if (!initialized) {
    return <div></div>;
  }
  return <>{children}</>;
};
