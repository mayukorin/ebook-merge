import React from "react";

interface Action {
  type: string;
  text: string;
}

export const FlashMessageDispatchContext = React.createContext(
  {} as {
    dispatch: React.Dispatch<Action>;
  }
);

export const FlashMessageContext = React.createContext(
  {} as {
    flashMessage: string;
  }
);

export const FlashMessageProvider: React.FC<{
  children: React.ReactNode;
}> = ({ children }) => {
  const flashMessageReducer = (flashMessage: string, action: Action) => {
    switch (action.type) {
      case "change": {
        return action.text;
      }
      case "reset": {
        return "";
      }
      default: {
        throw Error("Unknown action: " + action.type);
      }
    }
  };

  const initialMessage = "";

  const [flashMessage, dispatch] = React.useReducer(
    flashMessageReducer,
    initialMessage
  );

  React.useEffect(() => {
    if (flashMessage !== "") {
      setTimeout(() => {
        dispatch({
          type: "reset",
          text: "",
        });
      }, 5 * 1000);
    }
  }, [flashMessage]);

  return (
    <FlashMessageDispatchContext.Provider value={{ dispatch }}>
      <FlashMessageContext.Provider value={{ flashMessage }}>
        {children}
      </FlashMessageContext.Provider>
    </FlashMessageDispatchContext.Provider>
  );
};
