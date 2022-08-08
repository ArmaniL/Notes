import * as React from "react";
import { useEffect, useState } from "react";
import NoteGrid from "./NoteGrid";

export default function HomePage(props: any) {

    const [notes, setNotes] = useState([])
    const baseUrl = "http://localhost:1000/notes";




    useEffect(() => {

        const getNotes = async function () {
            const requestOptions = {
                method: "GET",
                headers: new Headers({
                    'Authorization': 'Basic ' + props.token,
                    'Content-Type': 'application/x-www-form-urlencoded',
                }),
            }
            const response = await fetch(baseUrl, requestOptions)
            const { notes } = (await response.json());
            setNotes(notes);
            console.log(notes)
        }

        getNotes();
    }, [])

    return (
        <NoteGrid notes={notes}></NoteGrid>
    )

}