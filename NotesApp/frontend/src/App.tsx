import * as React from "react";
import LoginForm from "./components/LoginForm";
import {
    BrowserRouter as Router,
    Route,
    Routes,

} from "react-router-dom";
import NoteGrid from "./components/NoteGrid";
import HomePage from "./components/HomePage";

export default function App() {
    const [token, setToken] = React.useState<string>("")
    return <Router>
        <Routes>
            <Route path="/" element={<LoginForm authControls={setToken}></LoginForm>}></Route>
            <Route path="/dashboard" element={<HomePage token={token} />}></Route>
        </Routes>
    </Router>
}