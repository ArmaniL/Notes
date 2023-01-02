import * as React from "react";
import {
  Button,
  Fade,
  FormGroup,
  FormLabel,
  Grow,
  Stack,
  Typography,
} from "@mui/material";
import { Navigate } from "react-router-dom";
import { UserContext } from "../App";
import { Input } from "@mui/material";

enum mode {
  Sign_Up = 0,
  Login,
  LoggedIn,
}
const baseUrl = "http://localhost:1000";
const loginUrl = `${baseUrl}/login`;
const signUpUrl = `${baseUrl}/signup`;

export default function LoginForm(props: any) {
  const { setUserInfo } = React.useContext(UserContext);
  const [formMode, setFormMode] = React.useState<mode>(mode.Login);
  const [email, setEmail] = React.useState<string>("");
  const [password, setPassword] = React.useState<string>("");
  const [errorMessage, setErrorMessage] = React.useState<string>("");

  return formMode === mode.LoggedIn ? (
    <Navigate to="/" />
  ) : (
    <Stack
      direction={"column"}
      spacing={10}
      justifyContent="center"
      alignItems="center"
    >
      <Stack direction={"column"} spacing={5} alignItems="center">
        <Fade in={true} style={{ transitionDelay: "200ms" }}>
          <Typography variant="h2">Cloud Notes</Typography>
        </Fade>
      </Stack>
      <div style={{ maxWidth: "300px", margin: "auto", marginTop: "100px" }}>
        <Grow in={true} style={{ transitionDelay: "600ms" }}>
          <FormGroup>
            <FormLabel>Email</FormLabel>
            <Input
              name="email"
              onChange={(event) => {
                const { value } = event.target;
                setEmail(value);
              }}
            />
            <FormLabel>Password</FormLabel>
            <Input
              type="password"
              onChange={(event) => {
                const { value } = event.target;
                setPassword(value);
              }}
            />
            <Button
              onClick={async () => {
                const requestOptions = {
                  body: JSON.stringify({
                    email,
                    password,
                  }),
                  method: "POST",
                };
                const requestUrl =
                  formMode === mode.Login ? loginUrl : signUpUrl;
                const response = await fetch(requestUrl, requestOptions);
                const responseBody = await response.json();
                const { message } = responseBody;
                if (message == "Succesful Login") {
                  const { shared_notes, token } = responseBody;
                  setUserInfo({ token, shared_notes });
                  setFormMode(mode.LoggedIn);
                } else if (message == "Succesful Sign Up") {
                  setFormMode(mode.Login);
                } else {
                  setErrorMessage(message);
                }
              }}
            >
              Submit
            </Button>

            <Button
              onClick={() => {
                const newMode =
                  formMode === mode.Login ? mode.Sign_Up : mode.Login;
                setFormMode(newMode);
              }}
            >
              {formMode === mode.Login
                ? "I do not have an account"
                : "I do have an account"}
            </Button>
            {errorMessage}
          </FormGroup>
        </Grow>
      </div>
    </Stack>
  );
}
