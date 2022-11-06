import React from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { ApiClientContext } from "../contexts/ApiClientContext";
import { FlashMessageDispatchContext } from "../contexts/FlashMessageContext";
import Box from "@mui/material/Box";
import ReactLoading from "react-loading";
import Typography from "@mui/material/Typography";
import { css } from "@emotion/react";

export const GmailOAuth2Callback: React.FC = () => {
  const ebookClient = React.useContext(ApiClientContext).ebook;
  const oAuth2TokenClient = React.useContext(ApiClientContext).oauth2Token;
  const { dispatch } = React.useContext(FlashMessageDispatchContext);
  let isCreated = false;

  const query = new URLSearchParams(useLocation().search);
  const navigate = useNavigate();

  const code: string | null = query.get("code");
  const userId: number = Number(query.get("state"));

  const mxAutoStyle = css`
    margin: 30px auto;
  `;

  React.useEffect(() => {
    if (code !== null && query.get("state") !== null && !isCreated) {
      isCreated = true;
      return () => {
        oAuth2TokenClient
        .createGmailApiOauth2Token({ body: { code, userId } })
        .then(() => {
          ebookClient
          .scanEbooks()
          .catch((e) => {
            console.error(e);
          })
          .then(() => {
            dispatch({
              type: "change",
              text: "Gmail連携が完了しました",
            });
            navigate("/ebooks");
          });
        })
        .catch((e) => {
          console.error(e);
        });
      };
    } 
  }, []);

  return (
    <>
      <ReactLoading
        type="spin"
        color="#795548"
        height="100px"
        width="100px"
        css={mxAutoStyle}
      />
      <Typography component="div">
        <Box textAlign="center">
          Gmailアドレスを連携中です．しばらくお待ちください．．
        </Box>
      </Typography>
    </>
  );
};