"use client";
import {useRouter} from "next/navigation";
import {Avatar, Box, Button, Container, CssBaseline, FormControl, Grid, TextField, Typography} from "@mui/material";
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Link from "next/link";
import {useEffect} from "react";
import {login} from "@/api/auth";

export default function Auth() {
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      router.push("/");
      router.refresh();
    }
  }, []);

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    try {
      const response = await login(formData.get("account") as string, formData.get("password") as string);
      localStorage.setItem("token", response.data.token);
      router.push("/");
      router.refresh();
    } catch (error) {
      console.log(error);
    }
  }

  return (
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box sx={{marginTop: (theme) => theme.spacing(8),
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center'}}>
          <Avatar sx={{margin: (theme) => theme.spacing(1),
            backgroundColor: "secondary.main"}}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <form onSubmit={handleSubmit} style={{width: '100%'}}>
            <FormControl fullWidth={true}>
              <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="account"
                  label="Email Address or Username"
                  name="account"
                  autoFocus
              />
              <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="current-password"
              />
            </FormControl>

            <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                sx={{margin: (theme) => theme.spacing(3, 0, 2)}}
            >
              Sign In
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="#">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="#">
                  {"Don't have an account? Sign Up"}
                </Link>
              </Grid>
            </Grid>
          </form>
        </Box>
      </Container>
  );
}