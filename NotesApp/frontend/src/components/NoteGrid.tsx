
import * as React from 'react';
import Grid from "@mui/material/Grid";
import { useEffect, useState } from "react";
import Note, { NoteProps } from "./Note";

export default function NoteGrid(props: { notes: Array<NoteProps> }) {
    const { notes } = props;

    return (

        <Grid container spacing={3}>

            {notes.map((note) => {
                return (
                    <Grid item xs={4}>
                        <Note {...note} ></Note>
                    </Grid>
                )
            })}

        </Grid>
    )
}