import * as React from "react";
import { useEffect, useState } from "react";
import NoteGrid from "./NoteGrid";
import IconButton from '@mui/material/IconButton';
import { Grid } from "@mui/material";
import { getNotes } from "../apiCalls";
import Note from "../Note";
import { Navigate } from 'react-router-dom'


export default function HomePage(props: any) {

    const [notes, setNotes] = useState([])
    const [goToNewNote, setGotoNewNote] = useState(false)
    const append = (note: Note) => {
        notes.unshift(note);
    }
    useEffect(() => {
        const makeCall = async () => {
            const notes = await getNotes(props.token);
            setNotes(notes);
            console.log(notes)
        }
        makeCall();
    }, [])


    // add navigaation to go to new note form

    return goToNewNote ? <Navigate to="/new-note" /> : (
        <Grid>
            <Grid item >
                <IconButton onClick={() => {
                    setGotoNewNote(true);
                }}>Add</IconButton>
            </Grid>
            <Grid item >
                <NoteGrid notes={notes}></NoteGrid>
            </Grid>
        </Grid>
    )

}