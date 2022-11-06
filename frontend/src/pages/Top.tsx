import React from "react";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { DamionFont } from "../materialui/theme";
import { ThemeProvider } from "@mui/material/styles";
import AuthService from "../auth";
import { useNavigate } from "react-router-dom";
import { FlashMessageDispatchContext } from "../contexts/FlashMessageContext";
import Grid from "@mui/material/Grid";

export const Top: React.FC = () => {
  const navigate = useNavigate();
  const { dispatch } = React.useContext(FlashMessageDispatchContext);

  return (
    <>
      <Card sx={{ minWidth: 275 }}>
        <CardContent>
          <ThemeProvider theme={DamionFont}>
            <Typography
              variant="h4"
              component="div"
              sx={{ flexGrow: 1 }}
              color="#009688"
              align="center"
            >
              EBookMerge
            </Typography>
          </ThemeProvider>
          <Typography component="div" align="center">
            Gmail と連携することで，複数サービスで購入した電子書籍を一元管理！
          </Typography>
        </CardContent>
        <CardActions>
          <Grid container justifyContent="center">
            <Button
              variant="contained"
              color="secondary"
              sx={{ borderRadius: 2 }}
              onClick={() =>
                AuthService.login().then(() => {
                  dispatch({
                    type: "change",
                    text: "ログインしました",
                  });
                  navigate("/ebooks");
                })
              }
            >
              使ってみる
              <br />
              Google アカウントでログイン
            </Button>
          </Grid>
        </CardActions>
        <CardContent>
          <Typography variant="body2">
            ※ 一元管理できる電子書籍は，Gmail
            でログインしているサービスで購入したものに限ります．それ以外のサービスで購入した電子書籍を一元管理したい場合は，その書籍の購入確認メールをGmail
            に転送してください．
          </Typography>
        </CardContent>
      </Card>
    </>
  );
};
