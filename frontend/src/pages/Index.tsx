import React from "react";
import { API_HOST, FIREBASE_API_KEY } from "../../env";

export const Index: React.FC = () => {
  console.log(API_HOST);
  console.log(FIREBASE_API_KEY);
  return (
    <div>
        ホーム画面
    </div>
  );
};