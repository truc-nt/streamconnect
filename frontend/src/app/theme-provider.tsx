"use client";

import {
  createTheme,
  ThemeProvider as MuiThemeProvider,
} from "@mui/material/styles";

const theme = createTheme({
  palette: {
    mode: "dark",
    text: {
      primary: "#fff",
    },
    primary: {
      main: "#08D2ED",
      contrastText: "white",
    },
    secondary: {
      main: "#fff",
    },
    background: {
      default: "#282A39",
      paper: "#282A39",
    },
    divider: "#fff",
  },
  components: {
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundImage: "none",
        },
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: "none",
        },
      },
    },
  },
});

export default function ThemeProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  return <MuiThemeProvider theme={theme}>{children}</MuiThemeProvider>;
}
