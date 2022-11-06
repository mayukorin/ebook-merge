import { ApiClientContext } from "../contexts/ApiClientContext";
import React from "react";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableRow from "@mui/material/TableRow";
import { Ebook } from "../openapi/models/Ebook";
import UpdateIcon from "@mui/icons-material/Update";
import ReactLoading from "react-loading";
import { css } from "@emotion/react";
import CableIcon from "@mui/icons-material/Cable";

export const EBookList: React.FC = () => {
  const ebookClient = React.useContext(ApiClientContext).ebook;
  const oAuth2TokenClient = React.useContext(ApiClientContext).oauth2Token;
  const [ebooks, setEbooks] = React.useState<Array<Ebook>>([]);
  const [loading, setLoading] = React.useState<boolean>(false);

  const mxAutoStyle = css`
    margin: 30px auto;
  `;

  React.useEffect(() => {
    setLoading(true);
    ebookClient
      .listEbooks()
      .then((res) => {
        console.log(res);
        if (res.ebooks === undefined) {
          return;
        }
        setEbooks(res.ebooks);
      })
      .catch((e) => {
        console.error(e);
      })
      .finally(() => setLoading(false));
  }, []);

    const updateOnClick = React.useCallback(() => {
      setLoading(true);
      ebookClient
        .scanEbooks()
        .catch((e) => {
          console.error(e);
        })
        .finally(() => setLoading(false));
    }, [ebookClient]);

    const connectOnClick = React.useCallback(() => {
      oAuth2TokenClient
        .confirmGmailApi()
        .then((res) => {
          window.location.replace(res.googleConcentPageUrl);
        })
        .catch((e) => {
          console.error(e);
        })
    }, [oAuth2TokenClient]);

  return (
    <div>
      <Button
        variant="contained"
        color="secondary"
        sx={{ borderRadius: 2 }}
        startIcon={<UpdateIcon />}
        onClick={() => updateOnClick()}
      >
        Gmailから更新
      </Button>
      <br />
      <br />
      <Typography
        variant="h4"
        component="span"
        sx={{ borderBottom: 5, borderColor: "primary.main" }}
      >
        電子書籍
      </Typography>
      {!loading ? (
        <>
          <Grid container sx={{ mt: 2 }}>
            <Grid item xs={12}>
              <TableContainer component={Paper}>
                <Table
                  aria-label="simple table"
                  sx={{ wordWrap: "break-word" }}
                >
                  <TableBody>
                    {ebooks.map((e) => (
                      <EbookRow ebook={e} key={e.id} />
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </Grid>
          </Grid>
          <br />
          <Grid container justifyContent="flex-end">
            <Button
              variant="contained"
              color="secondary"
              sx={{ borderRadius: 2 }}
              startIcon={<CableIcon />}
              onClick={() => connectOnClick()}
            >
              新しいGmailアドレスを連携
            </Button>
          </Grid>
        </>
      ) : (
        <>
          <ReactLoading
            type="spin"
            color="#795548"
            height="100px"
            width="100px"
            css={mxAutoStyle}
          />
        </>
      )}
    </div>
  );
};

const EbookRow: React.FC<{ ebook: Ebook }> = ({ ebook }) => {
  return (
    <TableRow key={ebook.title}>
      <TableCell>
        <Grid container>
          <Grid item xs={6} md={4}>
            {ebook.ebookService?.name === undefined ? (
              <EbookServiceBadge ebookServiceName={""} />
            ) : (
              <EbookServiceBadge ebookServiceName={ebook.ebookService.name} />
            )}
          </Grid>
          <Grid item xs={6} md={8}>
            {ebook.title}
          </Grid>
        </Grid>
      </TableCell>
    </TableRow>
  );
};

const EbookServiceBadge: React.FC<{ ebookServiceName: string }> = ({
  ebookServiceName,
}) => {
  let bgc = "";
  let tc = "";

  switch (ebookServiceName) {
    case "Kindle":
      bgc = "#DDA0DD37";
      tc = "#800080";
      break;
    case "BookLive":
      bgc = "#FFDAB93D";
      tc = "#FF7F50";
      break;
  }

  return (
    <Typography
      component="span"
      sx={{ m: 1, p: 1, backgroundColor: bgc, color: tc, borderRadius: 2 }}
    >
      {ebookServiceName}
    </Typography>
  );
};
