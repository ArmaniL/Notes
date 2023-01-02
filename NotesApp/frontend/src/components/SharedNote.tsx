import * as React from "react";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import { CardHeader, IconButton } from "@mui/material";
import ClearRoundedIcon from "@mui/icons-material/ClearRounded";
import IosShareIcon from "@mui/icons-material/IosShare";

export default function SharedNote(props: any) {
  const { id, header, content, user } = props;
  const [isHovered, SetHovered] = React.useState(false);

  const handleMouseIn = () => {
    SetHovered(true);
    setTimeout(() => SetHovered(false), 8000);
  };

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
      <CardHeader>{user}</CardHeader>
      <CardContent onMouseOver={handleMouseIn}>
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
