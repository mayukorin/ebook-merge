import React from "react";
import { EbookApi, OAuth2TokenApi } from "../openapi/apis";
import { Configuration } from "../openapi/runtime";
import { API_HOST } from "../../env";
import { AuthContext } from "./AuthContext";

const BASE_URL = API_HOST + "/v1";

export const ApiClientContext = React.createContext<{ ebook: EbookApi; oauth2Token: OAuth2TokenApi; }>({
  ebook: new EbookApi(
    new Configuration({
      basePath: BASE_URL,
    })
  ),
  oauth2Token: new OAuth2TokenApi(
    new Configuration({
      basePath: BASE_URL,
    })
  )
});

export const ApiClientContextProvider: React.FC<{
  children: React.ReactNode;
}> = ({ children }) => {
  const authContext = React.useContext(AuthContext);

  const ebook = React.useMemo(() => {
    if (authContext.token === null) {
      return new EbookApi(
        new Configuration({
          basePath: BASE_URL,
        })
      );
    } else {
      return new EbookApi(
        new Configuration({
          basePath: BASE_URL,
          apiKey: "Bearer " + authContext.token,
        })
      );
    }
  }, [authContext.token]);

  const oauth2Token = React.useMemo(() => {
    if (authContext.token === null) {
      return new OAuth2TokenApi(
        new Configuration({
          basePath: BASE_URL,
        })
      );
    } else {
      return new OAuth2TokenApi(
        new Configuration({
          basePath: BASE_URL,
          apiKey: "Bearer " + authContext.token,
        })
      );
    }
  }, [authContext.token]);

  return (
    <ApiClientContext.Provider value={{ ebook, oauth2Token }}>
      {children}
    </ApiClientContext.Provider>
  );
};
