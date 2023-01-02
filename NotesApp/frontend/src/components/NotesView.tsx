import * as React from "react";
import Note, { NoteProps } from "./Note";
import { NotesContext } from "./HomePage";
import { Divider, Fade, Stack } from "@mui/material";

export default function NotesView(props: {
  notes: Array<NoteProps>;
  callbacks: any;
}) {
  const initialAnimationDelay = 500;
  const { notes, callbacks } = React.useContext(NotesContext);
  return (
    <div style={{ overflowY: "scroll", padding: "50px" }}>
      <Stack
        direction="row"
        divider={<Divider orientation="vertical" flexItem />}
        spacing={2}
      >
        {notes.map((note, noteIndex) => {
          return (
            <Fade
              in={true}
              style={{ transitionDelay: `${500 + noteIndex * 150}ms` }}
            >
              <div>
                <Note {...note} remove={callbacks.delete}></Note>
              </div>
            </Fade>
          );
        })}
      </Stack>
    </div>
  );
}
