import * as React from 'react';
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import { Button } from '@mui/material';
import { Navigate } from 'react-router-dom'

enum mode {
    Sign_Up = 0,
    Login,
    LoggedIn
}
const baseUrl = "http://localhost:1000";
const loginUrl = `${baseUrl}/login`
const signUpUrl = `${baseUrl}/signup`

export default function LoginForm(props: any) {
    const [formMode, setFormMode] = React.useState<mode>(mode.Login);
    const [email, setEmail] = React.useState<string>("");
    const [password, setPassword] = React.useState<string>("");
    const [errorMessage, setErrorMessage] = React.useState<string>("")

    return formMode === mode.LoggedIn ? <Navigate to="/dashboard" /> :

        <form >
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <TextField
                        id="email-input"
                        name="email"
                        label="Email"
                        type="name"
                        onChange={(event) => {
                            const { value } = event.target;
                            setEmail(value);
                        }}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        id="password-input"
                        name="password"
                        label="Password"
                        type="password"
                        onChange={(event) => {
                            const { value } = event.target;
                            setPassword(value);
                        }}
                    />
                </Grid>

                <Grid item xs={12}>
                    <Button onClick={async () => {
                        const requestOptions = {
                            body: JSON.stringify({
                                email, password
                            }),
                            method: "POST"
                        }
                        const requestUrl = formMode === mode.Login ? loginUrl : signUpUrl
                        const response = await fetch(requestUrl, requestOptions);
                        const responseBody = await response.json();
                        const { message } = responseBody;
                        if (message == "Succesful Login") {
                            const { token } = responseBody
                            props.authControls(token)
                            setFormMode(mode.LoggedIn)
                        }
                        else if (message == "Succesful Sign Up") {
                            setFormMode(mode.Login)
                        }
                        else {
                            setErrorMessage(message)
                        }
                    }}>
                        Submit
                    </Button>
                </Grid>
                <Grid item xs={12}>
                    <Button onClick={
                        () => {
                            const newMode = formMode === mode.Login ? mode.Sign_Up : mode.Login;
                            setFormMode(newMode)
                        }
                    } >
                        {
                            formMode === mode.Login ? "I do not have an account" : "I do have an account"
                        }
                    </Button>
                </Grid>
                <Grid item xs={12}>
                    {errorMessage}
                </Grid>
            </Grid>
        </form>


}

