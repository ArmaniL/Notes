import * as React from "react";
import LoginForm from "./components/LoginForm";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./components/HomePage";
import NewNoteForm from "./components/NewNotePage";
import NotePage from "./components/NotePage";

export const UserContext = React.createContext({
  userInfo: null,
  setUserInfo: null,
});
export default function App() {
  const [userInfo, setUserInfo] = React.useState({
    token: "",
    shared_notes: [],
  });
  const [token, setToken] = React.useState("");
  return (
    <UserContext.Provider value={{ userInfo, setUserInfo }}>
      <Router>
        <Routes>
          <Route path="/" element={<HomePage />}></Route>
          <Route path="/login" element={<LoginForm></LoginForm>}></Route>
          <Route path="/new-note" element={<NewNoteForm />}></Route>
          <Route path="/note" element={<NotePage></NotePage>}></Route>
        </Routes>
      </Router>
    </UserContext.Provider>
  );
}
