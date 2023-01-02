import * as React from "react";
import { IconButton, Stack, TextField } from "@mui/material";
import { Navigate, useLocation } from "react-router";
import SaveAltIcon from "@mui/icons-material/SaveAlt";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { saveNote } from "../apiCalls";
import { UserContext } from "../App";
export default function NotePage(props: any) {
  const { userInfo } = React.useContext(UserContext);
  const [token, setToken] = React.useState(userInfo.token);
  const [redirect, setRedirect] = React.useState(false);
  const [header, setHeader] = React.useState("");
  const [content, setContent] = React.useState("");

  if (redirect) {
    return <Navigate to={"/"}></Navigate>;
  }

  return (
    <Stack direction={"column"}>
      <Stack direction={"row"}>
        <IconButton
          onClick={async () => {
            await saveNote(token, { header, content });
            setRedirect(true);
          }}
        >
          <SaveAltIcon></SaveAltIcon>
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
