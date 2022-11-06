import React from "react";
import AuthService from "../auth";

export const AuthContext = React.createContext<{
  token: string | null;
  initialized: boolean;
}>({
  token: null,
  initialized: false,
});

export const AuthContextProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [token, setToken] = React.useState<string | null>(null);
  const [initialized, setInitialized] = React.useState<boolean>(false);

  React.useEffect(() => {
    AuthService.observeToken((token) => {
      console.log(token);
      setToken(token);
      setInitialized(true);
    });
  }, []);

  return (
    <AuthContext.Provider value={{ token, initialized }}>
      {children}
    </AuthContext.Provider>
  );
};
