import { Stack, TextField, Button } from "@mui/material";
import * as React from "react";
import { shareNote } from "../apiCalls";
import { UserContext } from "../App";

export const ShareItem = (props: any) => {
  const { userInfo } = React.useContext(UserContext);
  const { token } = userInfo;
  const [inputType, setInputType] = React.useState<
    "primary" | "error" | "secondary" | "info" | "success" | "warning"
  >("primary");
  const [email, setEmail] = React.useState("");

  return (
    <Stack direction={"row"}>
      <TextField
        id="outlined-basic"
        label="Enter Email"
        variant="outlined"
        color={inputType}
        onChange={(event) => {
          setEmail(event.target.value);
        }}
      >
        {email}
      </TextField>
      <Button
        variant="contained"
        onClick={async () => {
          const response = await shareNote(token, props.id, email);
          if (response.status !== 200) {
            setInputType("warning");
          }
        }}
      >
        Share
      </Button>
    </Stack>
  );
};
