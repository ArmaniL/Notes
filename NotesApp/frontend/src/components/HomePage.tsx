import * as React from "react";
import { useEffect, useState } from "react";
import IconButton from "@mui/material/IconButton";
import { getNotes, deleteNote, getNote } from "../apiCalls";
import Note from "../Note";
import { Navigate } from "react-router-dom";
import AddIcon from "@mui/icons-material/Add";
import { UserContext } from "../App";
import NotesView from "./NotesView";
import { Fade, Menu, MenuItem, Stack, Typography } from "@mui/material";
import SharedNotesView from "./SharedNotesView";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";

export const NotesContext = React.createContext({ notes: [], callbacks: null });

export default function HomePage(props: any) {
  const { userInfo } = React.useContext(UserContext);
  const { shared_notes } = userInfo;
  const [token, setToken] = React.useState(userInfo.token);
  if (!token || token.length < 2) {
    return <Navigate to="/login" />;
  }

  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = () => {
    window.location.reload();
  };
  const [notes, setNotes] = useState([]);
  const [external_notes, setExternalNotes] = useState([]);
  const [goToNewNote, setGotoNewNote] = useState(false);
  const append = (note: Note) => {
    notes.unshift(note);
  };
  useEffect(() => {
    const getAllNotes = async () => {
      const notes = await getNotes(token);
      setNotes(notes);
      const promises = shared_notes.map((noteID: string) => {
        return getNote(token, noteID);
      });
      const responses = await Promise.all(promises);
      setExternalNotes(responses);
      console.log(external_notes);
    };
    getAllNotes();
  }, []);

  const removeNote = async (id: string) => {
    await deleteNote(token, id);
    const newNotes = notes.filter((note) => note.id != id);
    setNotes(newNotes);
  };

  const callbacks = {
    delete: removeNote,
  };

  return goToNewNote ? (
    <Navigate to="/new-note" />
  ) : (
    <NotesContext.Provider value={{ notes, callbacks }}>
      <Stack spacing={2}>
        <Fade in={true} style={{ transitionDelay: "400ms" }}>
          <Stack direction={"row"} spacing={4}>
            <Typography variant="h2">My Notes</Typography>

            <IconButton
              size="large"
              onClick={() => {
                setGotoNewNote(true);
              }}
            >
              <AddIcon />
            </IconButton>
            <IconButton size="large" onClick={handleClick}>
              <AccountCircleIcon />
            </IconButton>
            <Menu
              id="basic-menu"
              anchorEl={anchorEl}
              open={open}
              onClose={handleClose}
            >
              <MenuItem onClick={handleClose}>My account</MenuItem>
              <MenuItem onClick={handleLogout}>Logout</MenuItem>
            </Menu>
          </Stack>
        </Fade>
        <NotesView notes={notes} callbacks={callbacks}></NotesView>
        <Fade in={true} style={{ transitionDelay: "400ms" }}>
          <Typography variant="h2">Shared Notes</Typography>
        </Fade>
        <SharedNotesView
          notes={external_notes}
          callbacks={callbacks}
        ></SharedNotesView>
      </Stack>
    </NotesContext.Provider>
  );
}
