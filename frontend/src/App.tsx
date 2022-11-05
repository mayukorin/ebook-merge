import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Index } from "./pages/Index";

export const App: React.FC = () => {
  return (
    <div>
        <BrowserRouter>
            <Routes>
                <Route path={"/"} element={<Index />} />
            </Routes>
        </BrowserRouter>
    </div>
  );
};
