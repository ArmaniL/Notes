import * as React from 'react';
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import IconButton from '@mui/material/IconButton';
import { saveNote } from '../apiCalls';
import { Navigate } from 'react-router-dom'

export default function NewNoteForm(props: any) {
    const [header, setHeader] = React.useState<string>("");
    const [content, setContent] = React.useState<string>("");
    const [submitted, setSubmitted] = React.useState<boolean>(false);

    return submitted ? <Navigate to='/dashboard'></Navigate> : <form >
        <Grid container spacing={2}>
            <Grid item xs={12}>
                <TextField
                    variant='standard'
                    id="header-input"
                    name="Header"
                    label="Header"
                    type="string"
                    onChange={(event) => {
                        const { value } = event.target;
                        setHeader(value);
                    }}
                />
            </Grid>

            <Grid item xs={12}>
                <TextField
                    id="content-input"
                    name="Content"
                    label="Content"
                    type="string"
                    size="medium"
                    multiline={true}
                    onChange={(event) => {
                        const { value } = event.target;
                        setContent(value);
                    }}
                />
            </Grid>

            <Grid item xs={12}>
                <IconButton onClick={
                    async () => {
                        console.log(props.token, header, content)
                        const user = localStorage.getItem('user')
                        await saveNote(props.token, { header, content, user })
                        setSubmitted(true)
                    }
                } >
                    Finish
                </IconButton>
            </Grid>

        </Grid>
    </form>


}
