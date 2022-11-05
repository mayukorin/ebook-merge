import React from "react";
import AuthService from "../auth";

export const AuthContext = React.createContext<{token: string | null;}>({
    token: null,
});


export const AuthContextProvider: React.FC<{children: React.ReactNode }> = ({children}) => {
    const [token, setToken] = React.useState<string | null>(null);

    React.useEffect(() => {
      AuthService.observeToken((token) => {
        setToken(token);
      });
    }, []);
    
    return (
        <AuthContext.Provider value={{token}}>
            {children}
        </AuthContext.Provider>
    );
};