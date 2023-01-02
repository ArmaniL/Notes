import * as React from "react";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import {
  Button,
  CardHeader,
  IconButton,
  Input,
  Menu,
  MenuItem,
  Popover,
  Stack,
  TextField,
} from "@mui/material";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import { ShareItem } from "./ShareNotePopover";
export interface NoteProps {
  id: string;
  header: string;
  content: string;
  lastUpdated: string;
  authenticationToken: string;
  remove: any;
}

export default function Note(props: NoteProps) {
  const { id, header, content, lastUpdated, authenticationToken, remove } =
    props;
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const [popAnchor, setpopAnchor] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const handlePopoverClose = () => {
    setpopAnchor(null);
  };

  const handleShare = (event: React.MouseEvent<any>) => {
    setpopAnchor(event.currentTarget);
  };

  const popOpen = Boolean(popAnchor);
  const type = popOpen ? "simple-popover" : undefined;

  return (
    <Card
      style={{
        cursor: "pointer",
        minWidth: "250px",
        maxHeight: "400px",
        minHeight: "400px",
      }}
      sx={{ maxWidth: 345 }}
    >
      <CardHeader
        action={
          <>
            <IconButton onClick={handleClick}>{<MoreVertIcon />}</IconButton>
            <Menu
              id="basic-menu"
              anchorEl={anchorEl}
              open={open}
              onClose={handleClose}
            >
              <MenuItem
                onClick={() => {
                  remove(id);
                }}
              >
                Delete
              </MenuItem>
              <MenuItem onClick={(event) => handleShare(event)}>Share</MenuItem>
              <Popover
                id={type}
                open={popOpen}
                anchorEl={popAnchor}
                onClose={handlePopoverClose}
                anchorOrigin={{
                  vertical: "bottom",
                  horizontal: "left",
                }}
              >
                <ShareItem id={id}></ShareItem>
              </Popover>
            </Menu>
          </>
        }
      ></CardHeader>
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {header}
        </Typography>
        <Typography paragraph={true} variant="body2" color="text.secondary">
          {content}
        </Typography>
      </CardContent>
    </Card>
  );
}
