
import * as React from 'react';
import Grid from "@mui/material/Grid";
import { useEffect, useState } from "react";
import Note, { NoteProps } from "./Note";

export default function NoteGrid(props: { notes: Array<NoteProps> }) {
    const { notes } = props;

    return (

        <Grid container spacing={3}>

            {notes.map((note, noteIndex) => {
                return (
                    <Grid item xs={4} key={noteIndex}>
                        <Note {...note} ></Note>
                    </Grid>
                )
            })}

        </Grid>
    )
}