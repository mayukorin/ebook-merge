import { createTheme } from "@mui/material/styles";

export const theme = createTheme({
  palette: {
    primary: {
      light: "#52c7b8",
      main: "#009688",
      dark: "#00675b",
      contrastText: "#000000",
    },
    secondary: {
      light: "#a98274",
      main: "#795548",
      dark: "#4b2c20",
      contrastText: "#ffffff",
    },
    background: {
      paper: "#ffffff",
      default: "#000000",
    },
    text: { primary: "#000000", secondary: "#009688" },
  },
  typography: {
    fontFamily: ["Inter"].join(","),
  },
});

export const DamionFont = createTheme({
  typography: {
    fontFamily: ["Damion"].join(","),
  },
});

export const CourgetteFont = createTheme({
  typography: {
    fontFamily: ["Courgette"].join(","),
  },
});
