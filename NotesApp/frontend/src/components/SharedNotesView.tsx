import * as React from "react";
import Note, { NoteProps } from "./Note";
import { Divider, Fade, Stack } from "@mui/material";
import SharedNote from "./SharedNote";

export default function SharedNotesView(props: any) {
  const { notes } = props;
  return (
    <div style={{ overflowY: "scroll" }}>
      <Stack
        direction="row"
        divider={<Divider orientation="vertical" flexItem />}
        spacing={2}
      >
        {notes.map((note: any, noteIndex: number) => {
          return (
            <Fade
              in={true}
              style={{ transitionDelay: `${500 + noteIndex * 150}ms` }}
            >
              <div>
                <SharedNote {...note}></SharedNote>
              </div>
            </Fade>
          );
        })}
      </Stack>
    </div>
  );
}
