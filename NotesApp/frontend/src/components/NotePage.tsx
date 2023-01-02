import * as React from "react";
import { IconButton, Stack, TextField } from "@mui/material";
import { Navigate, useLocation } from "react-router";
import SaveAltIcon from "@mui/icons-material/SaveAlt";
import ClearIcon from "@mui/icons-material/Clear";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { updateNote, deleteNote } from "../apiCalls";
import { UserContext } from "../App";
export default function NotePage(props: any) {
  const { state } = useLocation();
  const { userInfo } = React.useContext(UserContext);
  const [token, setToken] = React.useState(userInfo.token);
  const { id } = state as any;
  const [header, setHeader] = React.useState((state as any).header);
  const [content, setContent] = React.useState((state as any).content);
  const [redirect, setRedirect] = React.useState(false);

  if (redirect) {
    return <Navigate to={"/"}></Navigate>;
  }

  return (
    <Stack direction={"column"}>
      <Stack direction={"row"}>
        <IconButton
          onClick={async () => {
            await updateNote(token, { noteId: id, header, content });
          }}
        >
          <SaveAltIcon></SaveAltIcon>
        </IconButton>
        <IconButton
          onClick={async () => {
            await deleteNote(token, id);
            setRedirect(true);
          }}
        >
          <ClearIcon></ClearIcon>
        </IconButton>
        <IconButton
          onClick={() => {
            setRedirect(true);
          }}
        >
          <ExitToAppIcon></ExitToAppIcon>
        </IconButton>
      </Stack>
      <div style={{ minWidth: "400px", margin: "auto", marginTop: "100px" }}>
        <Stack direction={"column"} alignContent="center" spacing={5}>
          <TextField
            variant="standard"
            defaultValue={header}
            label="Title"
            onChange={(event) => {
              const { value } = event.target;
              setHeader(value);
            }}
          />
          <TextField
            id="outlined-multiline-static"
            multiline
            defaultValue={content}
            onChange={(event) => {
              const { value } = event.target;
              setContent(value);
            }}
          />
        </Stack>
      </div>
    </Stack>
  );
}
